package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"sync"
	// controller
	controllers "mailform-demo-backend/src/interfaces/controllers"
	// domain
	domain "mailform-demo-backend/src/domain"
	// infrastructure
	infrastructure "mailform-demo-backend/src/infrastructure"
)

func main() {
	// lambda上で実行するとき
	if os.Getenv("GO_ENV") == "production" {
		fmt.Println("Production mode...")
		// lambda関数の実行
		fmt.Println("lambda Start...")
		lambda.Start(handler)
		// local環境で実行するとき
	} else {
		fmt.Println("Development mode...")
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal(err)
		}
		// 引数が足りてない場合、処理終了
		if (len(os.Args) - 1) < 4 {
			defer func() {
				fmt.Println("引数が足りてません")
			}()
			panic("処理を終了します")
		}
		dnm := domain.NuxtMail{}
		// 引数の中身を一件ずつ出力します
		for i, v := range os.Args {
			fmt.Printf("args[%d] -> %s\n", i, v)
		}

		// 送信元
		dnm.From = os.Args[1]
		// 送信先
		dnm.To = os.Args[2]
		// 件名
		dnm.Subject = os.Args[3]
		// メール本文
		dnm.Body = os.Args[4]

		//チャネルを作成して変数に代入
		r := make(chan domain.Res, 1)
		e := make(chan error, 1)
		// goroutineより、main関数が終了するより時間がかかるため
		// 処理まちをさせるために、syncパッケージのWaitGroupを利用します
		var wg sync.WaitGroup
		wg.Add(1)

		go SendMail(dnm, &wg, r, e)

		// メイン関数の処理がsay関数より早く終了するために
		// WaitGroupでgoroutineの待ち合わせをします
		wg.Wait()

		res := <-r
		err = <-e

		if (err == nil) || reflect.ValueOf(err).IsNil() {
			log.Println(res)
		} else {
			log.Println(err)
		}
		return
	}
}

func handler(ctx context.Context, arg domain.NuxtMail) (domain.Res, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		log.Printf("aws_request_id: %v", lc.AwsRequestID)
	}
	// goroutineより、main関数が終了するより時間がかかるため
	// 処理まちをさせるために、syncパッケージのWaitGroupを利用します
	var wg sync.WaitGroup
	wg.Add(1)

	//チャネルを作成して変数に代入
	r := make(chan domain.Res, 1)
	e := make(chan error, 1)

	go SendMail(arg, &wg, r, e)

	// メイン関数の処理がsay関数より早く終了するために
	// WaitGroupでgoroutineの待ち合わせをします
	wg.Wait()

	res := <-r
	err := <-e

	return res, err
}

func SendMail(arg domain.NuxtMail, wg *sync.WaitGroup, r chan domain.Res, e chan error) {
	// WaitGroupを終了させるコード
	defer wg.Done()
	// AWSの設定値読込
	ses := infrastructure.NewSES()
	// メール送信処理準備
	NuxtMailController := controllers.NewNuxtMailController(ses)
	// メール送信処理
	res, err := NuxtMailController.SendSESEmail(arg)
	r <- res
	e <- err
}
