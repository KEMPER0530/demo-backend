package controllers

import (
	"github.com/kemper0530/demo-backend/src/domain"
	"github.com/kemper0530/demo-backend/src/interfaces/aws"
	"github.com/kemper0530/demo-backend/src/usecase"
)

type NuxtMailController struct {
	Interactor usecase.NuxtMailInteractor
}

func NewNuxtMailController(ses aws.SES) *NuxtMailController {
	return &NuxtMailController{
		Interactor: usecase.NuxtMailInteractor{
			SES: &aws.SESRepository{SES: ses},
			NM:  &aws.NuxtMailRepository{},
		},
	}
}

func (controller *NuxtMailController) SendSESEmail(arg domain.NuxtMail) (domain.Res, error) {
	res, error := controller.Interactor.SendSESEmail(arg)
	return res, error
}
