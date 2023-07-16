package infrastructure

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kemper0530/demo-backend/src/domain"
	"log"
	"os"
)

var i *int
var from, to, sub, body, user, input, output, createdat *string

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

func LocalTest() (interface{}, error) {
	fmt.Println("Starting development mode...")
	loadEnvVars()
	initArgs()

	// ローカルでテストするためのダミーのAppSyncEventを作成
	var testEvent domain.AppSyncEvent

	switch *i {
	case 0:
		testEvent = domain.AppSyncEvent{
			OperationName: "createNuxtMail",
			Arguments: map[string]interface{}{
				"from":    *from,
				"to":      *to,
				"subject": *sub,
				"body":    *body,
			},
		}
	case 1:
		testEvent = domain.AppSyncEvent{
			OperationName: "createChatGptResult",
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

	result, err := RouteRequest(context.Background(), testEvent)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
