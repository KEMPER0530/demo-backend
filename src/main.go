package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"

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

var i *int
var from, to, sub, body, user, input, output, createdat *string

func main() {
	if os.Getenv("GO_ENV") == "production" {
		fmt.Println("Starting production mode...")
		lambda.Start(HandleRequests)
	} else {
		fmt.Println("Starting development mode...")
		loadEnvVars()
		initArgs()

		// ローカルでテストするためのダミーのAppSyncEventを作成
		var testEvent AppSyncEvent

		switch *i {
		case 0:
			testEvent = AppSyncEvent{
				OperationName: "createNuxtMail",
				Arguments: map[string]interface{}{
					"from":    *from,
					"to":      *to,
					"subject": *sub,
					"body":    *body,
				},
			}
		case 1:
			testEvent = AppSyncEvent{
				OperationName: "createChatGpt",
				Arguments: map[string]interface{}{
					"user":      *user,
					"input":     *input,
					"output":    *output,
					"createdat": *createdat,
				},
			}
		default:
			log.Fatal("Invalid mode flag")
		}

		result, err := HandleRequests(context.Background(), testEvent)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(result)
	}
}

func HandleRequests(ctx context.Context, request AppSyncEvent) (interface{}, error) {
	fmt.Printf("Request received: %+v\n", request) // デバッグログ
	switch request.OperationName {
	case "createChatGpt":
		var args domain.ChatGptResult
		if err := mapstructure.Decode(request.Arguments, &args); err != nil {
			return nil, err
		}
		return handleChatGptResult(domain.ChatGptResult{
			User:      args.User,
			Input:     args.Input,
			Output:    args.Output,
			Createdat: args.Createdat,
		})
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
	fmt.Printf("Handling NuxtMail: %+v\n", dnm) // デバッグログ
	NuxtMailController := controllers.NewNuxtMailController(infrastructure.NewSES())
	return NuxtMailController.SendSESEmail(dnm)
}

func handleChatGptResult(dnm domain.ChatGptResult) (domain.Res, error) {
	fmt.Printf("Handling ChatGptResult: %+v\n", dnm) // デバッグログ
	PutChatGptController := controllers.NewPutChatGptController()
	return PutChatGptController.PutChatGptResult(dnm, infrastructure.NewDynamoDB())
}

func initArgs() {
	i = flag.Int("i", 0, "mode flag(0:SESメール,1:ChatGptResult)")
	from = flag.String("f", os.Getenv("AWS_SES_TEST_FROM"), "SES から送信するメッセージの MAIL FROM ドメイン")
	to = flag.String("t", os.Getenv("AWS_SES_TEST_SUCCESS_TO"), "SES から送信するメッセージの MAIL TO ドメイン")
	sub = flag.String("s", "テスト件名", "メールの件名")
	body = flag.String("b", "テスト本文", "メールの本文")
	user = flag.String("us", "ユーザ名", "ユーザ名")
	input = flag.String("in", "インプット", "インプット")
	output = flag.String("ou", "アウトプット", "アウトプット")
	createdat = flag.String("cr", "作成日", "作成日")
	flag.Parse()
}

func loadEnvVars() {
	err := godotenv.Load(fmt.Sprintf("src/infrastructure/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
}
