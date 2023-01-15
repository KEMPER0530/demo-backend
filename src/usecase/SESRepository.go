package usecase

type SESRepository interface {
	GetRegion() string
	GetKeyid() string
	GetSecretkey() string
}
