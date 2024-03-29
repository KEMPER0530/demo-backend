package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

func NewDynamoDB() *dynamo.DB {
	//dynamoDBへ接続
	var db *dynamo.DB

	sess := session.Must(session.NewSession())

	if os.Getenv("GO_ENV") == "production" {
		db = dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DYNAMODB_REGION"))})
	} else {
		creds := credentials.NewStaticCredentials(os.Getenv("AWS_DYNAMODB_ACCESS_KEY_ID"), os.Getenv("AWS_DYNAMODB_SECRET_ACCESS_KEY"), "")
		db = dynamo.New(sess, aws.NewConfig().WithRegion(os.Getenv("AWS_DYNAMODB_REGION")).WithCredentials(creds))
	}

	return db
}
