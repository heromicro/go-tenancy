package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestCaptcha(t *testing.T) {
	e := base.BaseTester(t)
	obj := e.GET("v1/public/captcha").
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("captchaId", "picPath")

	data.Value("captchaId").String().NotEmpty()
	data.Value("picPath").String().NotEmpty()
}
