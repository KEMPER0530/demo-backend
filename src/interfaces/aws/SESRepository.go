package aws

import ()

type SESRepository struct {
	SES SES
}

func (ses *SESRepository) GetRegion() string {
	return ses.SES.GetRegion()
}

func (ses *SESRepository) GetKeyid() string {
	return ses.SES.GetKeyid()
}

func (ses *SESRepository) GetSecretkey() string {
	return ses.SES.GetSecretkey()
}
