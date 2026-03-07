package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kemper0530/demo-backend/src/domain"
	"testing"
)

type fakeSESClient struct {
	input  *ses.SendEmailInput
	output *ses.SendEmailOutput
	err    error
}

func (f *fakeSESClient) SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	f.input = input
	return f.output, f.err
}

func TestNuxtMailRepositorySendSessionError(t *testing.T) {
	origSession := newSessionSES
	origClient := newSESClient
	defer func() {
		newSessionSES = origSession
		newSESClient = origClient
	}()

	newSessionSES = func(cfgs ...*aws.Config) (*session.Session, error) {
		return nil, errors.New("session failed")
	}
	newSESClient = func(newSession *session.Session) sesEmailAPI {
		return &fakeSESClient{}
	}

	repo := &NuxtMailRepository{}
	msgID, err := repo.Send(domain.NuxtMail{From: "a", To: "b", Subject: "s", Body: "body"}, "region", "id", "secret")
	if err == nil {
		t.Fatalf("expected error")
	}
	if msgID != nil {
		t.Fatalf("message id must be nil on error")
	}
}

func TestNuxtMailRepositorySendEmailError(t *testing.T) {
	origSession := newSessionSES
	origClient := newSESClient
	defer func() {
		newSessionSES = origSession
		newSESClient = origClient
	}()

	var capturedCfg *aws.Config
	newSessionSES = func(cfgs ...*aws.Config) (*session.Session, error) {
		capturedCfg = cfgs[0]
		return &session.Session{}, nil
	}
	fc := &fakeSESClient{err: errors.New("send failed")}
	newSESClient = func(newSession *session.Session) sesEmailAPI {
		return fc
	}

	repo := &NuxtMailRepository{}
	msgID, err := repo.Send(domain.NuxtMail{From: "from", To: "to", Subject: "subject", Body: "body"}, "ap-northeast-1", "id", "secret")
	if err == nil {
		t.Fatalf("expected send error")
	}
	if msgID != nil {
		t.Fatalf("expected nil message id")
	}
	if capturedCfg == nil || aws.StringValue(capturedCfg.Region) != "ap-northeast-1" {
		t.Fatalf("region not passed to session config")
	}
	if fc.input == nil {
		t.Fatalf("send input should be set")
	}
}

func TestNuxtMailRepositorySendSuccess(t *testing.T) {
	origSession := newSessionSES
	origClient := newSESClient
	defer func() {
		newSessionSES = origSession
		newSESClient = origClient
	}()

	newSessionSES = func(cfgs ...*aws.Config) (*session.Session, error) {
		return &session.Session{}, nil
	}
	expectedID := "message-id"
	fc := &fakeSESClient{
		output: &ses.SendEmailOutput{
			MessageId: aws.String(expectedID),
		},
	}
	newSESClient = func(newSession *session.Session) sesEmailAPI {
		return fc
	}

	repo := &NuxtMailRepository{}
	msgID, err := repo.Send(domain.NuxtMail{From: "from@ex", To: "to@ex", Subject: "subject", Body: "body"}, "region", "id", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msgID == nil || *msgID != expectedID {
		t.Fatalf("unexpected message id")
	}
	if aws.StringValue(fc.input.Source) != "from@ex" {
		t.Fatalf("unexpected source")
	}
	if len(fc.input.Destination.ToAddresses) != 1 || aws.StringValue(fc.input.Destination.ToAddresses[0]) != "to@ex" {
		t.Fatalf("unexpected destination")
	}
	if aws.StringValue(fc.input.Message.Subject.Data) != "subject" {
		t.Fatalf("unexpected subject")
	}
	if aws.StringValue(fc.input.Message.Body.Text.Data) != "body" {
		t.Fatalf("unexpected body")
	}
}
