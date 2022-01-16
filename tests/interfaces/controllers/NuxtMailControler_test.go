package interfaces

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	mock "mailform-demo-backend/tests/mock_usecase"
	cont "mailform-demo-backend/src/interfaces/controllers"
	infra "mailform-demo-backend/src/infrastructure"
	"mailform-demo-backend/src/domain"
)

func TestNewNuxtMailController(t *testing.T) {
  fmt.Println("start test TestNewNuxtMailController")

	// 環境変数ファイルの読込
	err := godotenv.Load(fmt.Sprintf("../../../src/infrastructure/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		t.Errorf("環境変数読込エラー")
	}

	// パラメータ値の設定
	arg := domain.NuxtMail{}

	res := domain.Res{}

  // 値の設定
  ses := infra.NewSES()

	ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  t.Run("【正常系】メールが正常に配信された場合", func(t *testing.T){

		mockNuxtMailRepository := mock.NewMockNuxtMailRepository(ctrl)
		mockSESRepository := mock.NewMockSESRepository(ctrl)

    // AWS設定値
    arg.From = os.Getenv("AWS_SES_TEST_FROM")
    arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
    arg.Subject = "件名:テスト"
    arg.Body = "本文:テスト"

    NuxtMailController := cont.NewNuxtMailController(ses)

    if NuxtMailController == nil {
      t.Errorf("failed NewNuxtMailController()")
    }

		// mockの設定
    tmp := "5dba2abe-5ea1-4edb-863f-6408a0456567"
		tmpr := &tmp

		mockSESRepository.EXPECT().GetRegion().Return("region").AnyTimes()
		mockSESRepository.EXPECT().GetKeyid().Return("keyid").AnyTimes()
		mockSESRepository.EXPECT().GetSecretkey().Return("secretkey").AnyTimes()

		var r string = mockSESRepository.GetRegion()
		var k string = mockSESRepository.GetKeyid()
		var s string = mockSESRepository.GetSecretkey()

		mockNuxtMailRepository.EXPECT().Send(arg,r,k,s).Return(tmpr,nil)

		NuxtMailController.Interactor.SES = mockSESRepository
		NuxtMailController.Interactor.NM = mockNuxtMailRepository

    res,err = NuxtMailController.SendSESEmail(arg)

    assert.Equal(t, res.Responce, 200,"not equal")
    assert.Equal(t, res.Result, "success","not equal")

    t.Logf("res.Responce: %d", res.Responce)
    t.Logf("res.Result: %s", res.Result)
    t.Logf("err.Error: %s", err)
  })

  t.Run("【異常系】メール配信時にエラー", func(t *testing.T){

		mockNuxtMailRepository := mock.NewMockNuxtMailRepository(ctrl)
		mockSESRepository := mock.NewMockSESRepository(ctrl)

    // AWS設定値
    arg.From = os.Getenv("AWS_SES_TEST_FROM")
    arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
    arg.Subject = "件名:テスト"
    arg.Body = "本文:テスト"

    NuxtMailController := cont.NewNuxtMailController(ses)

    if NuxtMailController == nil {
      t.Errorf("failed NewNuxtMailController()")
    }

		// mockの設定
    tmp := "5dba2abe-5ea1-4edb-863f-6408a0456567"
		tmpr := &tmp

		mockSESRepository.EXPECT().GetRegion().Return("region").AnyTimes()
		mockSESRepository.EXPECT().GetKeyid().Return("keyid").AnyTimes()
		mockSESRepository.EXPECT().GetSecretkey().Return("secretkey").AnyTimes()

		var r string = mockSESRepository.GetRegion()
		var k string = mockSESRepository.GetKeyid()
		var s string = mockSESRepository.GetSecretkey()

		mockNuxtMailRepository.EXPECT().Send(arg,r,k,s).Return(tmpr, errors.New("Mock Error"))

		NuxtMailController.Interactor.SES = mockSESRepository
		NuxtMailController.Interactor.NM = mockNuxtMailRepository

    res,err = NuxtMailController.SendSESEmail(arg)

    assert.Equal(t, res.Responce, 500,"not equal")
    assert.Equal(t, res.Result, "failed","not equal")

    t.Logf("res.Responce: %d", res.Responce)
    t.Logf("res.Result: %s", res.Result)
    t.Logf("err.Error: %s", err)
  })

  fmt.Println("end test TestNewNuxtMailController")
}