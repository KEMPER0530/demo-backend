package usecase

import ()

type SESRepository interface {
	GetRegion() string
	GetKeyid() string
	GetSecretkey() string
}
