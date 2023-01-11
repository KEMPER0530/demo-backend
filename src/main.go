package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/joho/godotenv"

    "mailform-demo-backend/src/interfaces/controllers"
    "mailform-demo-backend/src/domain"
    "mailform-demo-backend/src/infrastructure"
)

func main() {
    if os.Getenv("GO_ENV") == "production" {
        fmt.Println("Starting production mode...")
        lambda.Start(handler)
    } else {
        fmt.Println("Starting development mode...")
        loadEnvVars()
        checkArgs()
        dnm := domain.NuxtMail{
            From:    os.Args[1],
            To:      os.Args[2],
            Subject: os.Args[3],
            Body:    os.Args[4],
        }
        result, err := SendMail(dnm)
        if err != nil {
            log.Println(err)
            return
        }
        log.Println(result)
    }
}

func handler(ctx context.Context, arg domain.NuxtMail) (domain.Res, error) {
    result, err := SendMail(arg)
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

func checkArgs() {
    if len(os.Args) < 5 {
        log.Fatal("Missing arguments.")
    }
}

func SendMail(dnm domain.NuxtMail) (domain.Res, error) {
    return controllers.SendMail(dnm)
}
