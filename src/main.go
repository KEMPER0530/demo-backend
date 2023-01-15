package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"

	"mailform-demo-backend/src/domain"
	"mailform-demo-backend/src/infrastructure"
	"mailform-demo-backend/src/interfaces/controllers"
)

var i *int
var from, to, sub, body, prompt *string

func main() {
	if os.Getenv("GO_ENV") == "production" {
		fmt.Println("Starting production mode...")
		lambda.Start(handlerNuxtMail)
		lambda.Start(handlerChatGpt)
	} else {
		fmt.Println("Starting development mode...")
		loadEnvVars()
		initArgs()
		var result domain.Res
		var err error
		switch *i {
		case 0:
			result, err = SendMail(domain.NuxtMail{
				From:    *from,
				To:      *to,
				Subject: *sub,
				Body:    *body,
			})
		case 1:
			result, err = SendPrompt(domain.ChatGpt{
				Prompt: *prompt,
			})
		default:
			log.Fatal("Invalid mode flag")
		}
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(result)
	}
}

func handlerNuxtMail(ctx context.Context, arg domain.NuxtMail) (domain.Res, error) {
	result, err := SendMail(arg)
	if err != nil {
		return domain.Res{}, err
	}
	return result, nil
}

func handlerChatGpt(ctx context.Context, arg domain.ChatGpt) (domain.Res, error) {
	result, err := SendPrompt(arg)
	if err != nil {
		return domain.Res{}, err
	}
	return result, nil
}

func loadEnvVars() {
	err := godotenv.Load(fmt.Sprintf("src/infrastructure/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
}

func initArgs() {
	i = flag.Int("i", 0, "mode flag(0:SESメール,1:ChatGpt)")
	from = flag.String("f", os.Getenv("AWS_SES_TEST_FROM"), "SES から送信するメッセージの MAIL FROM ドメイン")
	to = flag.String("t", os.Getenv("AWS_SES_TEST_SUCCESS_TO"), "SES から送信するメッセージの MAIL TO ドメイン")
	sub = flag.String("s", "テスト件名", "メールの件名")
	body = flag.String("b", "テスト本文", "メールの本文")
	prompt = flag.String("p", "Go言語の特徴について教えて下さい.", "ChatGptのprompt")
	flag.Parse()
}

func SendMail(dnm domain.NuxtMail) (domain.Res, error) {
	NuxtMailController := controllers.NewNuxtMailController(infrastructure.NewSES())
	return NuxtMailController.SendSESEmail(dnm)
}

func SendPrompt(dnm domain.ChatGpt) (domain.Res, error) {
	ChatGptController := controllers.NewChatGptController(infrastructure.NewGPT())
	return ChatGptController.SendChatGptPrompt(dnm)
}
