package infrastructure

import (
	"context"
	"errors"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/aws"
	"testing"
)

type fakeSES struct{}

func (f *fakeSES) GetRegion() string    { return "region" }
func (f *fakeSES) GetKeyid() string     { return "keyid" }
func (f *fakeSES) GetSecretkey() string { return "secret" }

type fakeNuxtController struct {
	res domain.Res
	err error
	arg domain.NuxtMail
}

func (f *fakeNuxtController) SendSESEmail(arg domain.NuxtMail) (domain.Res, error) {
	f.arg = arg
	return f.res, f.err
}

type fakeChatController struct {
	res domain.Res
	err error
	arg domain.ChatGptResult
}

func (f *fakeChatController) PutChatGptResult(arg domain.ChatGptResult, d *dynamo.DB) (domain.Res, error) {
	f.arg = arg
	return f.res, f.err
}

func TestRouteRequestUnknownOperation(t *testing.T) {
	_, err := RouteRequest(context.Background(), domain.AppSyncEvent{
		OperationName: "unknown",
		Arguments:     map[string]interface{}{},
	})
	if err == nil {
		t.Fatalf("expected error for unknown operation")
	}
}

func TestRouteRequestDecodeErrors(t *testing.T) {
	origDecode := decodeRouting
	defer func() { decodeRouting = origDecode }()

	decodeRouting = func(input interface{}, output interface{}) error {
		return errors.New("decode error")
	}

	_, err := RouteRequest(context.Background(), domain.AppSyncEvent{
		OperationName: "createNuxtMail",
		Arguments:     map[string]interface{}{"from": "a"},
	})
	if err == nil {
		t.Fatalf("expected decode error for nuxt mail")
	}

	_, err = RouteRequest(context.Background(), domain.AppSyncEvent{
		OperationName: "createChatGptResult",
		Arguments:     map[string]interface{}{"user": "a"},
	})
	if err == nil {
		t.Fatalf("expected decode error for chatgpt")
	}
}

func TestRouteRequestNuxtMailSuccess(t *testing.T) {
	origSES := newSESRouting
	origController := newNuxtMailController
	defer func() {
		newSESRouting = origSES
		newNuxtMailController = origController
	}()

	fc := &fakeNuxtController{res: domain.Res{Response: 200, Result: "ok"}}
	newSESRouting = func() *SES {
		return &SES{}
	}
	newNuxtMailController = func(s aws.SES) nuxtMailController {
		return fc
	}

	got, err := RouteRequest(context.Background(), domain.AppSyncEvent{
		OperationName: "createNuxtMail",
		Arguments: map[string]interface{}{
			"from":    "from@example.com",
			"to":      "to@example.com",
			"subject": "sub",
			"body":    "body",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	res := got.(domain.Res)
	if res.Response != 200 || fc.arg.To != "to@example.com" {
		t.Fatalf("unexpected routing result")
	}
}

func TestRouteRequestChatGptSuccess(t *testing.T) {
	origController := newPutChatGptController
	origDynamo := newDynamoDBRouting
	defer func() {
		newPutChatGptController = origController
		newDynamoDBRouting = origDynamo
	}()

	fc := &fakeChatController{res: domain.Res{Response: 200, Result: "ok"}}
	newPutChatGptController = func() chatGptController {
		return fc
	}
	newDynamoDBRouting = func() *dynamo.DB { return nil }

	got, err := RouteRequest(context.Background(), domain.AppSyncEvent{
		OperationName: "createChatGptResult",
		Arguments: map[string]interface{}{
			"user":      "user",
			"input":     "in",
			"output":    "out",
			"createdat": "2023-01-01",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	res := got.(domain.Res)
	if res.Response != 200 || fc.arg.User != "user" {
		t.Fatalf("unexpected routing result")
	}
}

func TestHandleFunctionsErrorPropagation(t *testing.T) {
	origSES := newSESRouting
	origNuxt := newNuxtMailController
	origChat := newPutChatGptController
	origDynamo := newDynamoDBRouting
	defer func() {
		newSESRouting = origSES
		newNuxtMailController = origNuxt
		newPutChatGptController = origChat
		newDynamoDBRouting = origDynamo
	}()

	newSESRouting = func() *SES { return &SES{} }
	newNuxtMailController = func(s aws.SES) nuxtMailController {
		return &fakeNuxtController{res: domain.Res{Response: 500, Result: "failed"}, err: errors.New("send failed")}
	}
	_, err := handleNuxtMail(domain.NuxtMail{To: "x"})
	if err == nil {
		t.Fatalf("expected error from nuxt controller")
	}

	newPutChatGptController = func() chatGptController {
		return &fakeChatController{res: domain.Res{Response: 500, Result: "failed"}, err: errors.New("put failed")}
	}
	newDynamoDBRouting = func() *dynamo.DB { return nil }
	_, err = handleChatGptResult(domain.ChatGptResult{User: "u"})
	if err == nil {
		t.Fatalf("expected error from chat controller")
	}
}
