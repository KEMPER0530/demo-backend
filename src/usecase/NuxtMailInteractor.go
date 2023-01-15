package usecase

import (
	"mailform-demo-backend/src/domain"
)

type NuxtMailInteractor struct {
	SES SESRepository
	NM  NuxtMailRepository
}

func (i *NuxtMailInteractor) SendSESEmail(arg domain.NuxtMail) (domain.Res, error) {
	region, keyid, secret := i.SES.GetRegion(), i.SES.GetKeyid(), i.SES.GetSecretkey()
	_, err := i.NM.Send(arg, region, keyid, secret)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}
	return domain.Res{Response: 200, Result: "success"}, nil
}
