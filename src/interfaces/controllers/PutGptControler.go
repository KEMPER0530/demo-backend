package controllers

import (
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/database"
	"github.com/kemper0530/demo-backend/src/usecase"
)

type PutChatGptController struct {
	Interactor usecase.ChatGptInteractor
}

func NewPutChatGptController() *PutChatGptController {
	return &PutChatGptController{
		Interactor: usecase.ChatGptInteractor{
			CGR: &database.ChatGptRepository{},
		},
	}
}

func (controller *PutChatGptController) PutChatGptResult(arg domain.ChatGptResult, d *dynamo.DB) (domain.Res, error) {
	res, err := controller.Interactor.PutChatGptResult(arg, d)
	return res, err
}
