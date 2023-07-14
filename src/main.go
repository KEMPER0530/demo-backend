package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"

	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/infrastructure"
	"github.com/kemper0530/demo-backend/src/interfaces/controllers"
)

type AppSyncEvent struct {
	OperationName string                 `mapstructure:"operationName"`
	Arguments     map[string]interface{} `mapstructure:"arguments"`
}

type NuxtMailArgs struct {
	Body      string `mapstructure:"body"`
	CreatedAt string `mapstructure:"createdat"`
	From      string `mapstructure:"from"`
	Subject   string `mapstructure:"subject"`
	To        string `mapstructure:"to"`
	UpdatedAt string `mapstructure:"updatedat"`
}

func main() {
	lambda.Start(HandleRequests)
}

func HandleRequests(ctx context.Context, request AppSyncEvent) (interface{}, error) {
	switch request.OperationName {
	case "createChatGpt":
		var args domain.ChatGpt
		if err := mapstructure.Decode(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleChatGpt(args)
	case "createNuxtMail":
		var args NuxtMailArgs
		if err := mapstructure.Decode(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleNuxtMail(domain.NuxtMail{
			From:    args.From,
			To:      args.To,
			Subject: args.Subject,
			Body:    args.Body,
		})
	default:
		return nil, fmt.Errorf("unknown operation: %s", request.OperationName)
	}
}

func handleNuxtMail(dnm domain.NuxtMail) (domain.Res, error) {
	NuxtMailController := controllers.NewNuxtMailController(infrastructure.NewSES())
	return NuxtMailController.SendSESEmail(dnm)
}

func handleChatGpt(dnm domain.ChatGpt) (domain.Res, error) {
	ChatGptController := controllers.NewChatGptController(infrastructure.NewGPT())
	return ChatGptController.SendChatGptPrompt(dnm)
}
