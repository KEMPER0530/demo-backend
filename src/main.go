package main

import (
  "log"
  "fmt"
  "os"
  "context"
  "github.com/joho/godotenv"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/lambdacontext"
  // controller
  controllers "mailform-demo-backend/src/interfaces/controllers"
  // domain
  domain "mailform-demo-backend/src/domain"
)

func init() {
    fmt.Println("init...")

    if os.Getenv("GO_ENV") == "production" {
      fmt.Println("Production mode...")
    }else{
      fmt.Println("Development mode...")
      // 環境変数ファイルの読込
      err := godotenv.Load(fmt.Sprintf("src/infrastructure/%s.env", os.Getenv("GO_ENV")))
      if err != nil {
        log.Fatal(err)
      }
		}
}

func main() {
  fmt.Println("lambda Start...")
  lambda.Start(handler)
}

func handler(ctx context.Context, arg domain.NuxtMail) (domain.Res, error) {
  lc, ok := lambdacontext.FromContext(ctx)
  if ok {
    log.Printf("aws_request_id: %v", lc.AwsRequestID)
  }

  NuxtMailController := controllers.NewNuxtMailController()
  res, err := NuxtMailController.SendSESEmail(arg)
  return res, err
}
