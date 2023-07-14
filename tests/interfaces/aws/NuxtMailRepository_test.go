package interfaces

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kemper0530/demo-backend/src/domain"
	infra "github.com/kemper0530/demo-backend/src/infrastructure"
	cont "github.com/kemper0530/demo-backend/src/interfaces/controllers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	fmt.Println("start test TestSend")

	if os.Getenv("GO_ENV") == "development" {
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("../../../src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			t.Errorf("環境変数読込エラー")
		}
	}

	// パラメータ値の設定
	arg := domain.NuxtMail{}

	// オブジェクトの生成
	ses := infra.NewSES()

	t.Run("【正常系】Sendメソッドの検証(正常配信)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = e

		// 検証
		assert.NotEqual(t, *m, nil)
		t.Logf("msgID: %s", *m)

	})

	t.Run("【正常系】Sendメソッドの検証(バウンス用)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_BOUNCE_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = e

		// 検証
		assert.NotEqual(t, *m, nil)
		t.Logf("msgID: %s", *m)

	})

	t.Run("【正常系】Sendメソッドの検証(応答)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_OOTO_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = e

		// 検証
		assert.NotEqual(t, *m, nil)
		t.Logf("msgID: %s", *m)

	})

	t.Run("【正常系】Sendメソッドの検証(拒否)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_COMPLAINT_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = e

		// 検証
		assert.NotEqual(t, *m, nil)
		t.Logf("msgID: %s", *m)

	})

	t.Run("【異常系】Sendメソッドの検証(Fromの指定なし)", func(t *testing.T) {

		// AWS設定値
		arg.From = ""
		arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = m

		// 検証
		assert.NotEqual(t, e, nil)
		t.Logf("Error: %s", e)

	})

	t.Run("【異常系】Sendメソッドの検証(Toの指定なし)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = ""
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = m

		// 検証
		assert.NotEqual(t, e, nil)
		t.Logf("Error: %s", e)

	})

	t.Run("【異常系】Sendメソッドの検証(Regionの不正値)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = ""
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = m

		// 検証
		assert.NotEqual(t, e, nil)
		t.Logf("Error: %s", e)

	})

	t.Run("【異常系】Sendメソッドの検証(Keyidの不正値)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = ""
		var s string = NuxtMailController.Interactor.SES.GetSecretkey()

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = m

		// 検証
		assert.NotEqual(t, e, nil)
		t.Logf("Error: %s", e)

	})

	t.Run("【異常系】Sendメソッドの検証(Secretkeyの不正値)", func(t *testing.T) {

		// AWS設定値
		arg.From = os.Getenv("AWS_SES_TEST_FROM")
		arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
		arg.Subject = "件名:テスト"
		arg.Body = "本文:テスト"

		NuxtMailController := cont.NewNuxtMailController(ses)

		var r string = NuxtMailController.Interactor.SES.GetRegion()
		var k string = NuxtMailController.Interactor.SES.GetKeyid()
		var s string = ""

		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		_ = m

		// 検証
		assert.NotEqual(t, e, nil)
		t.Logf("msgID: %s", e)

	})
	fmt.Println("end test TestSend")
}

func BenchmarkMailSend(b *testing.B) {
	fmt.Println("start test BenchmarkMailSend")

	if os.Getenv("GO_ENV") == "development" {
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("../../../src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			b.Errorf("環境変数読込エラー")
		}
	}

	// パラメータ値の設定
	arg := domain.NuxtMail{}

	// オブジェクトの生成
	ses := infra.NewSES()

	// AWS設定値
	arg.From = os.Getenv("AWS_SES_TEST_FROM")
	arg.To = os.Getenv("AWS_SES_TEST_SUCCESS_TO")
	arg.Subject = "件名:テスト"
	arg.Body = "本文:テスト"

	NuxtMailController := cont.NewNuxtMailController(ses)

	var r string = NuxtMailController.Interactor.SES.GetRegion()
	var k string = NuxtMailController.Interactor.SES.GetKeyid()
	var s string = NuxtMailController.Interactor.SES.GetSecretkey()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// sendメソッドの実行
		m, e := NuxtMailController.Interactor.NM.Send(arg, r, k, s)
		// 検証
		assert.NotEqual(b, *m, nil)
		b.Logf("msgID: %s", *m)

		if e != nil {
			b.Errorf("failed Send()")
		}
	}
	b.StopTimer()
	fmt.Println("end test BenchmarkMailSend")
}
