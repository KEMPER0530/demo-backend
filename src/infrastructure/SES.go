package infrastructure

import (
	"os"
)

type SES struct {
	Region    string
	Keyid     string
	Secretkey string
}

func NewSES() *SES {
	return &SES{
		Region:    os.Getenv("AWS_SES_REGION"),
		Keyid:     os.Getenv("AWS_SES_ACCESS_KEY_ID"),
		Secretkey: os.Getenv("AWS_SES_SECRET_KEY"),
	}
}

func (ses *SES) GetRegion() string {
	return ses.Region
}

func (ses *SES) GetKeyid() string {
	return ses.Keyid
}

func (ses *SES) GetSecretkey() string {
	return ses.Secretkey
}
