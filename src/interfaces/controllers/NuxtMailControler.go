package controllers

import (
	"mailform-demo-backend/src/domain"
	"mailform-demo-backend/src/interfaces/aws"
	"mailform-demo-backend/src/usecase"
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
