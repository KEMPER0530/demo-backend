package controllers

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/usecase"
	"testing"
)

type stubSES struct{}

func (s *stubSES) GetRegion() string    { return "region" }
func (s *stubSES) GetKeyid() string     { return "key" }
func (s *stubSES) GetSecretkey() string { return "secret" }

type fakeSESRepo struct{}

func (f *fakeSESRepo) GetRegion() string    { return "region" }
func (f *fakeSESRepo) GetKeyid() string     { return "key" }
func (f *fakeSESRepo) GetSecretkey() string { return "secret" }

type fakeMailRepo struct {
	sendErr error
}

func (f *fakeMailRepo) Send(arg domain.NuxtMail, region string, id string, secret string) (*string, error) {
	tmp := "id"
	if f.sendErr != nil {
		return &tmp, f.sendErr
	}
	return &tmp, nil
}

type fakeChatRepo struct {
	createErr error
	putErr    error
}

func (f *fakeChatRepo) CreateTableIfNotExists(d *dynamo.DB, tableName string) error {
	return f.createErr
}
func (f *fakeChatRepo) PutResult(table *dynamo.Table, arg domain.ChatGptResult) error {
	return f.putErr
}

func TestNewNuxtMailControllerAndSend(t *testing.T) {
	controller := NewNuxtMailController(&stubSES{})
	if controller == nil {
		t.Fatalf("expected controller")
	}
	if controller.Interactor.SES == nil || controller.Interactor.NM == nil {
		t.Fatalf("interactor dependencies not initialized")
	}

	controller.Interactor = usecase.NuxtMailInteractor{
		SES: &fakeSESRepo{},
		NM:  &fakeMailRepo{},
	}
	res, err := controller.SendSESEmail(domain.NuxtMail{From: "f", To: "t", Subject: "s", Body: "b"})
	if err != nil || res.Response != 200 {
		t.Fatalf("expected successful send")
	}

	controller.Interactor = usecase.NuxtMailInteractor{
		SES: &fakeSESRepo{},
		NM:  &fakeMailRepo{sendErr: errors.New("send failed")},
	}
	res, err = controller.SendSESEmail(domain.NuxtMail{From: "f", To: "t", Subject: "s", Body: "b"})
	if err == nil || res.Response != 500 {
		t.Fatalf("expected send failure")
	}
}

func TestNewPutChatGptControllerAndPut(t *testing.T) {
	controller := NewPutChatGptController()
	if controller == nil {
		t.Fatalf("expected controller")
	}
	if controller.Interactor.CGR == nil {
		t.Fatalf("interactor repository not initialized")
	}

	controller.Interactor.CGR = &fakeChatRepo{}
	db := dynamo.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("ap-northeast-1"))
	res, err := controller.PutChatGptResult(domain.ChatGptResult{User: "u"}, db)
	if err != nil || res.Response != 200 {
		t.Fatalf("expected success")
	}

	controller.Interactor.CGR = &fakeChatRepo{createErr: errors.New("create failed")}
	res, err = controller.PutChatGptResult(domain.ChatGptResult{User: "u"}, db)
	if err == nil || res.Response != 500 {
		t.Fatalf("expected create failure")
	}
}
