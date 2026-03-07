package infrastructure

import (
	"os"
	"testing"
)

func TestNewSESAndGetters(t *testing.T) {
	t.Setenv("AWS_SES_REGION", "ap-northeast-1")
	t.Setenv("AWS_SES_ACCESS_KEY_ID", "key")
	t.Setenv("AWS_SES_SECRET_KEY", "secret")

	ses := NewSES()
	if ses == nil {
		t.Fatalf("expected non nil SES")
	}
	if ses.GetRegion() != "ap-northeast-1" {
		t.Fatalf("unexpected region")
	}
	if ses.GetKeyid() != "key" {
		t.Fatalf("unexpected key id")
	}
	if ses.GetSecretkey() != "secret" {
		t.Fatalf("unexpected secret key")
	}

	_ = os.Getenv("AWS_SES_REGION")
}
