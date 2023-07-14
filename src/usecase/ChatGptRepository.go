package usecase

import (
	"github.com/kemper0530/demo-backend/src/domain"
)

type ChatGptRepository interface {
	SendPrompt(arg domain.ChatGpt, keyid string) (string, error)
}
