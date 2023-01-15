package usecase

import (
	"mailform-demo-backend/src/domain"
)

type ChatGptRepository interface {
	SendPrompt(arg domain.ChatGpt, keyid string) (string, error)
}
