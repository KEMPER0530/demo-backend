package infrastructure

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"testing"
	"github.com/stretchr/testify/assert"
	infra "mailform-demo-backend/src/infrastructure"
)

func TestSES(t *testing.T) {
  fmt.Println("start test TestSES")

	if os.Getenv("GO_ENV") == "development" {
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("../../src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			t.Errorf("環境変数読込エラー")
		}
	}

  t.Run("【正常系】AWSの設定値が正常に取得できているかどうかを検証", func(t *testing.T){
    var r string = os.Getenv("AWS_SES_REGION")
    var k string = os.Getenv("AWS_SES_ACCESS_KEY_ID")
    var s string = os.Getenv("AWS_SES_SECRET_KEY")

    ses := infra.NewSES()

    if ses == nil {
      t.Errorf("failed NewSES()")
    }

    assert.Equal(t, r, ses.GetRegion(),"not equal")
    assert.Equal(t, k, ses.GetKeyid(),"not equal")
    assert.Equal(t, s, ses.GetSecretkey(),"not equal")

    t.Logf("ses: %p", ses )
    t.Logf("ses.Region: %s", r )
    t.Logf("ses.Keyid: %s", k )
    t.Logf("ses.Secretkey: %s", s )
  })

  t.Run("【異常系】AWSの設定値が正常に取得できているかどうかを検証", func(t *testing.T){
    var r string = "region"
    var k string = "keyid"
    var s string = "secretid"

    ses := infra.NewSES()

    if ses == nil {
      t.Errorf("failed NewSES()")
    }

    assert.NotEqual(t, r, ses.GetRegion(),"not equal")
    assert.NotEqual(t, k, ses.GetKeyid(),"not equal")
    assert.NotEqual(t, s, ses.GetSecretkey(),"not equal")

    t.Logf("ses: %p", ses)
    t.Logf("ses.Region: %s", r)
    t.Logf("ses.Keyid: %s", k)
    t.Logf("ses.Secretkey: %s", s)
  })
  fmt.Println("end test TestSES")
}

func BenchmarkSES(b *testing.B) {
    fmt.Println("start test BenchmarkSES")
    b.ResetTimer()
    for i:= 0; i < b.N; i++ {
      ses := infra.NewSES()
      if ses == nil {
       b.Errorf("failed NewSES()")
      }
    }
    b.StopTimer()
    fmt.Println("end test BenchmarkSES")
}
