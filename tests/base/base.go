package base

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/multi"
)

func BaseTester(t *testing.T) *httpexpect.Expect {
	handler := initialize.App()
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://127.0.0.1:8089/",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func BaseWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := BaseTester(t)
	obj := e.POST("v1/public/admin/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(1)
	user.Value("userName").String().Equal("admin")
	user.Value("email").String().Equal("admin@admin.com")
	user.Value("nickName").String().Equal("超级管理员")
	user.Value("authorityName").String().Equal("超级管理员")
	user.Value("authorityType").Number().Equal(multi.AdminAuthority)
	user.Value("authorityId").String().Equal("999")
	user.Value("defaultRouter").String().Equal("dashboard")
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func TenancyWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := BaseTester(t)
	obj := e.POST("v1/public/merchant/login").
		WithJSON(map[string]interface{}{"username": "a303176530", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(2)
	user.Value("userName").String().Equal("a303176530")
	user.Value("email").String().Equal("a303176530@admin.com")
	user.Value("nickName").String().Equal("商户管理员")
	user.Value("authorityName").String().Equal("商户管理员")
	user.Value("authorityType").Number().Equal(multi.TenancyAuthority)
	user.Value("authorityId").String().Equal("998")
	user.Value("defaultRouter").String().Equal("dashboard")
	user.Value("tenancyName").String().Equal("宝安中心人民医院")
	user.Value("tenancyId").Number().Equal(1)
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func DeviceWithLoginTester(t *testing.T) *httpexpect.Expect {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).First(&tenancy).Error
	if err != nil {
		t.Fatal(err)
	}
	e := BaseTester(t)
	obj := e.POST("v1/public/device/login").
		WithJSON(map[string]interface{}{
			"uuid":       tenancy.UUID,
			"name":       "八两金",
			"phone":      "13845687419",
			"sex":        2,
			"age":        32,
			"locName":    "泌尿科一区",
			"bedNum":     "15",
			"hospitalNo": "88956655",
			"disease":    "不孕不育"}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	data.Value("user").NotNull()
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func BaseLogOut(auth *httpexpect.Expect) {
	obj := auth.GET("v1/auth/logout").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("退出登录")
}
