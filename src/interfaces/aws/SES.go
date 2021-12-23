package aws

type SES interface {
	GetRegion() string
	GetKeyid() string
	GetSecretkey() string
}
