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

var getenvMain = os.Getenv
var printMain = fmt.Println
var lambdaStartMain = lambda.Start
var localTestMain = infrastructure.LocalTest
var logMain = log.Println
var routeRequestMain = infrastructure.RouteRequest

func main() {
	if getenvMain("GO_ENV") == "production" {
		printMain("Starting production mode...")
		lambdaStartMain(HandleRequests)
	} else {
		printMain("Starting development mode...")
		result, err := localTestMain()
		if err != nil {
			logMain(err)
			return
		}
		logMain(result)
	}
}

func HandleRequests(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
	return routeRequestMain(ctx, request)
}
