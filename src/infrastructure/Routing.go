package infrastructure

import (
	"context"
	"fmt"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/aws"
	"github.com/kemper0530/demo-backend/src/interfaces/controllers"
	"github.com/mitchellh/mapstructure"
)

type nuxtMailController interface {
	SendSESEmail(arg domain.NuxtMail) (domain.Res, error)
}

type chatGptController interface {
	PutChatGptResult(arg domain.ChatGptResult, d *dynamo.DB) (domain.Res, error)
}

var decodeRouting = mapstructure.Decode
var newSESRouting = NewSES
var newNuxtMailController = func(ses aws.SES) nuxtMailController {
	return controllers.NewNuxtMailController(ses)
}
var newPutChatGptController = func() chatGptController {
	return controllers.NewPutChatGptController()
}
var newDynamoDBRouting = NewDynamoDB

func RouteRequest(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
	fmt.Printf("Request received: %+v\n", request) // デバッグログ
	switch request.OperationName {
	case "createChatGptResult":
		var args domain.ChatGptResult
		if err := decodeRouting(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleChatGptResult(args)
	case "createNuxtMail":
		var args domain.NuxtMail
		if err := decodeRouting(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleNuxtMail(args)
	default:
		return nil, fmt.Errorf("unknown operation: %s", request.OperationName)
	}
}

func handleNuxtMail(dnm domain.NuxtMail) (domain.Res, error) {
	fmt.Printf("Handling NuxtMail: %+v\n", dnm) // デバッグログ
	NuxtMailController := newNuxtMailController(newSESRouting())
	return NuxtMailController.SendSESEmail(dnm)
}

func handleChatGptResult(dcgr domain.ChatGptResult) (domain.Res, error) {
	fmt.Printf("Handling ChatGptResult: %+v\n", dcgr) // デバッグログ
	PutChatGptController := newPutChatGptController()
	return PutChatGptController.PutChatGptResult(dcgr, newDynamoDBRouting())
}
