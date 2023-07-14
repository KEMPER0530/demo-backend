package usecase

import (
	"github.com/kemper0530/demo-backend/src/domain"
)

type ChatGptInteractor struct {
	GPT GPTRepository
	CGR ChatGptRepository
}

func (i *ChatGptInteractor) SendChatGptPrompt(arg domain.ChatGpt) (domain.Res, error) {
	keyid := i.GPT.GetKeyid()
	result, err := i.CGR.SendPrompt(arg, keyid)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}
	return domain.Res{Response: 200, Result: result}, nil
}
