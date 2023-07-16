package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/infrastructure"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") == "production" {
		fmt.Println("Starting production mode...")
		lambda.Start(HandleRequests)
	} else {
		fmt.Println("Starting development mode...")
		result, err := infrastructure.LocalTest()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(result)
	}
}

func HandleRequests(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
	return infrastructure.RouteRequest(ctx, request)
}
