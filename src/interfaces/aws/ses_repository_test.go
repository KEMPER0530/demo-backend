package aws

import "testing"

type fakeSES struct {
	region string
	keyID  string
	secret string
}

func (f *fakeSES) GetRegion() string    { return f.region }
func (f *fakeSES) GetKeyid() string     { return f.keyID }
func (f *fakeSES) GetSecretkey() string { return f.secret }

func TestSESRepositoryDelegatesToSES(t *testing.T) {
	repo := &SESRepository{
		SES: &fakeSES{
			region: "ap-northeast-1",
			keyID:  "key",
			secret: "secret",
		},
	}

	if repo.GetRegion() != "ap-northeast-1" {
		t.Fatalf("unexpected region")
	}
	if repo.GetKeyid() != "key" {
		t.Fatalf("unexpected key id")
	}
	if repo.GetSecretkey() != "secret" {
		t.Fatalf("unexpected secret")
	}
}
