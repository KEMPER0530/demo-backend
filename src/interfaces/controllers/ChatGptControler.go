package controllers

import (
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/chat"
	"github.com/kemper0530/demo-backend/src/usecase"
)

type ChatGptController struct {
	Interactor usecase.ChatGptInteractor
}

func NewChatGptController(gpt chat.GPT) *ChatGptController {
	return &ChatGptController{
		Interactor: usecase.ChatGptInteractor{
			GPT: &chat.GPTRepository{GPT: gpt},
			CGR: &chat.ChatGptRepository{},
		},
	}
}

func (controller *ChatGptController) SendChatGptPrompt(arg domain.ChatGpt) (domain.Res, error) {
	res, error := controller.Interactor.SendChatGptPrompt(arg)
	return res, error
}
