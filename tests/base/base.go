package base

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/initialize/cache"
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
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewCompactPrinter(t),
		},
	})
}

func BaseWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := BaseTester(t)
	obj := e.POST("v1/public/admin/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusOK)
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

func TenancyWithLoginTester(t *testing.T) (*httpexpect.Expect, uint) {
	username, _ := cache.GetCacheString(g.TENANCY_CONFIG.Mysql.Dbname + ":username")
	id, _ := cache.GetCacheUint(g.TENANCY_CONFIG.Mysql.Dbname + ":id")
	if username == "" || id == 0 {
		t.Fatal("创建商户失败")
	}
	e := BaseTester(t)
	obj := e.POST("v1/public/merchant/login").
		WithJSON(map[string]interface{}{"username": username, "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusOK)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(2)
	user.Value("userName").String().Equal(username)
	user.Value("authorityName").String().Equal("商户管理员")
	user.Value("authorityType").Number().Equal(multi.TenancyAuthority)
	user.Value("authorityId").String().Equal("998")
	user.Value("defaultRouter").String().Equal("dashboard")
	user.Value("tenancyName").String().Equal("多商户平台直营医院")
	user.Value("tenancyId").Number().Equal(1)
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	}), uint(id)
}

func DeviceWithLoginTester(t *testing.T) *httpexpect.Expect {

	uuid, _ := cache.GetCacheString(g.TENANCY_CONFIG.Mysql.Dbname + ":uuid")
	username, _ := cache.GetCacheString(g.TENANCY_CONFIG.Mysql.Dbname + ":username")
	if uuid == "" {
		auth := BaseWithLoginTester(t)
		defer BaseLogOut(auth)

		_, username, uuid = CreateTenancy(auth, "tenancy_hospital", http.StatusOK, "创建成功")
		cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":uuid", uuid, 0)
		cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":username", username, 0)
	}

	if uuid == "" {
		t.Fatal("创建商户失败")
	}

	e := BaseTester(t)
	obj := e.POST("v1/public/device/login").
		WithJSON(map[string]interface{}{
			"uuid":       uuid,
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
	obj.Value("status").Number().Equal(http.StatusOK)
	obj.Value("message").String().Equal("登录成功")
	obj.Value("data").Object().Value("user").NotNull()
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func BaseLogOut(auth *httpexpect.Expect) {
	obj := auth.GET("v1/auth/logout").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusOK)
	obj.Value("message").String().Equal("退出登录")
}

func CreateTenancy(auth *httpexpect.Expect, username string, status int, message string) (uint, string, string) {
	data := map[string]interface{}{
		"username":      username,
		"name":          "宝安妇女儿童医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"status":        g.StatusTrue,
		"sysRegionCode": 1,
	}
	url := "v1/admin/tenancy/createTenancy"
	res := ResponseKeys{
		{Key: "id", Value: uint(0)},
		{Key: "uuid", Value: ""},
		{Key: "username", Value: ""},
	}
	Create(auth, url, data, res, status, message)
	return res.GetId(), res.GetStringValue("username"), res.GetStringValue("uuid")
}

func DeleteTenancy(auth *httpexpect.Expect, id uint) {
	url := fmt.Sprintf("v1/admin/tenancy/deleteTenancy/%d", id)
	defer Delete(auth, url, http.StatusOK, "删除成功")
}
