package usecase

import (
    "mailform-demo-backend/src/domain"
)

type NuxtMailInteractor struct {
    ses SESRepository
    nm  NuxtMailRepository
}

func NewNuxtMailInteractor(ses SESRepository, nm NuxtMailRepository) *NuxtMailInteractor {
    return &NuxtMailInteractor{ses: ses, nm: nm}
}

func (i *NuxtMailInteractor) SendSESEmail(arg domain.NuxtMail) (domain.Res, error) {
    region, keyid, secret := i.ses.GetRegion(), i.ses.GetKeyid(), i.ses.GetSecretkey()
    msgID, err := i.nm.Send(arg, region, keyid, secret)
    if err != nil {
        return domain.Res{Responce: 500, Result: "failed"}, err
    }
    return domain.Res{Responce: 200, Result: "success"}, nil
}
