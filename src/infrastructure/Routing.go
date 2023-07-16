package infrastructure

import (
	"context"
	"fmt"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/controllers"
	"github.com/mitchellh/mapstructure"
)

func RouteRequest(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
	fmt.Printf("Request received: %+v\n", request) // デバッグログ
	switch request.OperationName {
	case "createChatGptResult":
		var args domain.ChatGptResult
		if err := mapstructure.Decode(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleChatGptResult(args)
	case "createNuxtMail":
		var args domain.NuxtMail
		if err := mapstructure.Decode(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleNuxtMail(args)
	default:
		return nil, fmt.Errorf("unknown operation: %s", request.OperationName)
	}
}

func handleNuxtMail(dnm domain.NuxtMail) (domain.Res, error) {
	fmt.Printf("Handling NuxtMail: %+v\n", dnm) // デバッグログ
	NuxtMailController := controllers.NewNuxtMailController(NewSES())
	return NuxtMailController.SendSESEmail(dnm)
}

func handleChatGptResult(dcgr domain.ChatGptResult) (domain.Res, error) {
	fmt.Printf("Handling ChatGptResult: %+v\n", dcgr) // デバッグログ
	PutChatGptController := controllers.NewPutChatGptController()
	return PutChatGptController.PutChatGptResult(dcgr, NewDynamoDB())
}
