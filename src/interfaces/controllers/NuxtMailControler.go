package controllers

import (
    "mailform-demo-backend/src/interfaces/ses"
    "mailform-demo-backend/src/usecase"
    "mailform-demo-backend/src/domain"
)

type NuxtMailController struct {
    Interactor usecase.NuxtMailInteractor
}

func NewNuxtMailController() *NuxtMailController {
    return &NuxtMailController{
        Interactor: usecase.NuxtMailInteractor{
            NM: &ses.NuxtMailRepository{},
        },
    }
}

func (controller *NuxtMailController) SendSESEmail(arg domain.NuxtMail) (domain.Res,error){
    res,error := controller.Interactor.SendSESEmail(arg)
    return res,error
}
