package infrastructure

import "testing"

func TestNewDynamoDBProductionAndDevelopment(t *testing.T) {
	t.Setenv("AWS_DYNAMODB_REGION", "ap-northeast-1")

	t.Setenv("GO_ENV", "production")
	prod := NewDynamoDB()
	if prod == nil {
		t.Fatalf("expected production db")
	}

	t.Setenv("GO_ENV", "development")
	t.Setenv("AWS_DYNAMODB_ACCESS_KEY_ID", "dummy")
	t.Setenv("AWS_DYNAMODB_SECRET_ACCESS_KEY", "dummy")
	dev := NewDynamoDB()
	if dev == nil {
		t.Fatalf("expected development db")
	}
}
