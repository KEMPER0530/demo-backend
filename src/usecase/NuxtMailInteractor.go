package usecase

import (
	"log"
	"mailform-demo-backend/src/domain"
)

type NuxtMailInteractor struct {
	SES SESRepository
	NM  NuxtMailRepository
}

func (interactor *NuxtMailInteractor) SendSESEmail(arg domain.NuxtMail) (res domain.Res, err error) {
	res = domain.Res{}
	// AWS設定値取得
	region := interactor.SES.GetRegion()
	id := interactor.SES.GetKeyid()
	secret := interactor.SES.GetSecretkey()

	msgID, err := interactor.NM.Send(arg, region, id, secret)
	if err != nil {
		res.Responce = 500
		res.Result = "failed"
		return res, err
	}
	log.Println(msgID)
	res.Responce = 200
	res.Result = "success"

	return res, nil
}
