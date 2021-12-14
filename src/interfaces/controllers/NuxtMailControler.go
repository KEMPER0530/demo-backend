package controllers

import (
    "mailform-demo-backend/src/interfaces/gateway"
    "mailform-demo-backend/src/usecase"
    "mailform-demo-backend/src/domain"
)

type NuxtMailController struct {
    Interactor usecase.NuxtMailInteractor
}

func NewNuxtMailController() *NuxtMailController {
    return &NuxtMailController{
        Interactor: usecase.NuxtMailInteractor{
            NM: &gateway.NuxtMailRepository{},
        },
    }
}

func (controller *NuxtMailController) SendSESEmail(arg domain.NuxtMail) (domain.Res,error){
    res,error := controller.Interactor.SendSESEmail(arg)
    return res,error
}
