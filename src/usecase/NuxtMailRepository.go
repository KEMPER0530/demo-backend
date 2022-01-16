package usecase

import (
	"mailform-demo-backend/src/domain"
)

type NuxtMailRepository interface {
	Send(arg domain.NuxtMail, region string, id string, secret string) (*string, error)
}
