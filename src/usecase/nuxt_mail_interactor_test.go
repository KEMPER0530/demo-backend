package usecase

import (
	"errors"
	"github.com/kemper0530/demo-backend/src/domain"
	"testing"
)

type fakeSESRepo struct {
	region string
	keyID  string
	secret string
}

func (f *fakeSESRepo) GetRegion() string    { return f.region }
func (f *fakeSESRepo) GetKeyid() string     { return f.keyID }
func (f *fakeSESRepo) GetSecretkey() string { return f.secret }

type fakeMailRepo struct {
	gotArg    domain.NuxtMail
	gotRegion string
	gotID     string
	gotSecret string
	sendErr   error
}

func (f *fakeMailRepo) Send(arg domain.NuxtMail, region string, id string, secret string) (*string, error) {
	f.gotArg = arg
	f.gotRegion = region
	f.gotID = id
	f.gotSecret = secret
	tmp := "id"
	return &tmp, f.sendErr
}

func TestNuxtMailInteractorSendSESEmail(t *testing.T) {
	ses := &fakeSESRepo{region: "r", keyID: "k", secret: "s"}
	mail := &fakeMailRepo{}
	interactor := &NuxtMailInteractor{SES: ses, NM: mail}
	arg := domain.NuxtMail{From: "from", To: "to", Subject: "sub", Body: "body"}

	res, err := interactor.SendSESEmail(arg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Response != 200 || res.Result != "success" {
		t.Fatalf("unexpected response: %#v", res)
	}
	if mail.gotRegion != "r" || mail.gotID != "k" || mail.gotSecret != "s" {
		t.Fatalf("credentials were not forwarded")
	}
	if mail.gotArg.To != "to" {
		t.Fatalf("argument was not forwarded")
	}

	mail.sendErr = errors.New("send failed")
	res, err = interactor.SendSESEmail(arg)
	if err == nil {
		t.Fatalf("expected error")
	}
	if res.Response != 500 || res.Result != "failed" {
		t.Fatalf("unexpected failure response: %#v", res)
	}
}
