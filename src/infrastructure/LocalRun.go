package infrastructure

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kemper0530/demo-backend/src/domain"
	"io"
	"os"
)

type localArgs struct {
	mode      int
	from      string
	to        string
	subject   string
	body      string
	user      string
	input     string
	output    string
	createdAt string
}

var getenvLocal = os.Getenv
var loadEnvFile = func(path string) error { return godotenv.Load(path) }
var parseLocalArgs = defaultParseLocalArgs
var routeRequestLocal = RouteRequest

func defaultParseLocalArgs() (localArgs, error) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	mode := fs.Int("i", 0, "mode flag(0:SESメール,1:ChatGptResult)")
	from := fs.String("f", getenvLocal("AWS_SES_TEST_FROM"), "SES から送信するメッセージの MAIL FROM ドメイン")
	to := fs.String("t", getenvLocal("AWS_SES_TEST_SUCCESS_TO"), "SES から送信するメッセージの MAIL TO ドメイン")
	sub := fs.String("s", "テスト件名", "メールの件名")
	body := fs.String("b", "テスト本文", "メールの本文")
	user := fs.String("us", "ユーザ名", "ユーザ名")
	input := fs.String("in", "インプット", "インプット")
	output := fs.String("ou", "アウトプット", "アウトプット")
	createdat := fs.String("cr", "作成日", "作成日")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return localArgs{}, err
	}

	return localArgs{
		mode:      *mode,
		from:      *from,
		to:        *to,
		subject:   *sub,
		body:      *body,
		user:      *user,
		input:     *input,
		output:    *output,
		createdAt: *createdat,
	}, nil
}

func loadEnvVars() error {
	return loadEnvFile(fmt.Sprintf("src/infrastructure/%s.env", getenvLocal("GO_ENV")))
}

func LocalTest() (interface{}, error) {
	fmt.Println("Starting development mode...")
	if err := loadEnvVars(); err != nil {
		return nil, err
	}
	args, err := parseLocalArgs()
	if err != nil {
		return nil, err
	}

	// ローカルでテストするためのダミーのAppSyncEventを作成
	var testEvent domain.AppSyncEvent

	switch args.mode {
	case 0:
		testEvent = domain.AppSyncEvent{
			OperationName: "createNuxtMail",
			Arguments: map[string]interface{}{
				"from":    args.from,
				"to":      args.to,
				"subject": args.subject,
				"body":    args.body,
			},
		}
	case 1:
		testEvent = domain.AppSyncEvent{
			OperationName: "createChatGptResult",
			Arguments: map[string]interface{}{
				"user":      args.user,
				"input":     args.input,
				"output":    args.output,
				"createdat": args.createdAt,
			},
		}
	default:
		return nil, errors.New("invalid mode flag")
	}

	result, err := routeRequestLocal(context.Background(), testEvent)
	if err != nil {
		return nil, err
	}
	return result, nil
}
