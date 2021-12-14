package usecase

import (
    "log"
    "mailform-demo-backend/src/domain"
)

type NuxtMailInteractor struct {
    NM NuxtMailRepository
}

func (interactor *NuxtMailInteractor) SendSESEmail(arg domain.NuxtMail) (res domain.Res, err error) {
    res = domain.Res{}

    from := arg.From
    to := arg.To
    subject := arg.Subject
    body := arg.Body

    msgID, err := interactor.NM.Send(from,to,subject,body)
    if err != nil {
        res.Responce = 500
        res.Result = "failed"
        return res,err
    }
    log.Println(msgID)
    res.Responce = 200
    res.Result = "success"

    return res,nil
}
