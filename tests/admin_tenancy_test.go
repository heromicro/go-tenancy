package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestTenancyList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/tenancy/getTenancyList"
	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "total", Value: 0},
	}
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}

func TestTenancyByRegion(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	tenancyId, _, _ := base.CreateTenancy(auth, "bafvetyy", http.StatusOK, "创建成功")
	if tenancyId == 0 {
		t.Fatal("创建失败")
	}
	defer base.DeleteTenancy(auth, tenancyId)
	base.Get(auth, "v1/admin/tenancy/getTenancies/1", nil, http.StatusOK, "获取成功")
}

func TestLoginTenancy(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/tenancy/loginTenancy/1").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	data.Value("token").String().NotEmpty()
	data.Value("url").String().NotEmpty()
}

func TestGetTenancyCount(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	tenancyId, _, _ := base.CreateTenancy(auth, "bafvetyy", http.StatusOK, "创建成功")
	if tenancyId == 0 {
		t.Fatal("创建失败")
	}
	defer base.DeleteTenancy(auth, tenancyId)

	obj := auth.GET("v1/admin/tenancy/getTenancyCount").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Object().Value("invalid").Equal(0)
	obj.Value("data").Object().Value("valid").Equal(2)
}

func TestTenancyProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	tenancyId, _, _ := base.CreateTenancy(auth, "bafvetyy", http.StatusOK, "创建成功")
	if tenancyId == 0 {
		t.Fatal("创建失败")
	}
	defer base.DeleteTenancy(auth, tenancyId)

	if tenancyId > 0 {
		update := map[string]interface{}{
			"username":      "bafvetyy",
			"name":          "宝安妇女儿童附属医院",
			"tele":          "0755-235689111",
			"address":       "xxx街道667号",
			"businessTime":  "08:30-17:40",
			"status":        g.StatusTrue,
			"sysRegionCode": 3,
		}

		obj := auth.PUT(fmt.Sprintf("v1/admin/tenancy/updateTenancy/%d", int(tenancyId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		tenancy := obj.Value("data").Object()

		tenancy.Value("name").String().Equal(update["name"].(string))
		tenancy.Value("tele").String().Equal(update["tele"].(string))
		tenancy.Value("address").String().Equal(update["address"].(string))
		tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
		tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))
		tenancy.Value("status").Number().Equal(update["status"].(int))

		obj = auth.GET(fmt.Sprintf("v1/admin/tenancy/getTenancyById/%d", int(tenancyId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		tenancy = obj.Value("data").Object()

		tenancy.Value("id").Number().Ge(0)
		tenancy.Value("uuid").String().NotEmpty()
		tenancy.Value("name").String().Equal(update["name"].(string))
		tenancy.Value("tele").String().Equal(update["tele"].(string))
		tenancy.Value("address").String().Equal(update["address"].(string))
		tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
		tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))
		tenancy.Value("status").Number().Equal(update["status"].(int))

		// changePasswordMap
		obj = auth.GET(fmt.Sprintf("v1/admin/tenancy/changePasswordMap/%d", int(tenancyId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// setTenancyRegion
		obj = auth.POST("v1/admin/tenancy/setTenancyRegion").
			WithJSON(map[string]interface{}{
				"id":            tenancyId,
				"sysRegionCode": 2,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// changeCopyMap
		obj = auth.GET(fmt.Sprintf("v1/admin/tenancy/changeCopyMap/%d", int(tenancyId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// setCopyProductNum
		obj = auth.POST(fmt.Sprintf("v1/admin/tenancy/setCopyProductNum/%d", int(tenancyId))).
			WithJSON(map[string]interface{}{
				"copyNum": 0,
				"num":     2,
				"type":    1,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// changeTenancyStatus
		obj = auth.POST("v1/admin/tenancy/changeTenancyStatus").
			WithJSON(map[string]interface{}{
				"id":     tenancyId,
				"status": 2,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")
	}
}
func TestTenancyRegisterError(t *testing.T) {

	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	{
		tenancyId, _, _ := base.CreateTenancy(auth, "bafvetyy", http.StatusOK, "创建成功")
		if tenancyId == 0 {
			t.Fatal("创建失败")
		}
		defer base.DeleteTenancy(auth, tenancyId)
	}
	{
		tenancyId, _, _ := base.CreateTenancy(auth, "bafvetyy", http.StatusBadRequest, "添加失败:商户名称已被注冊")
		if tenancyId > 0 {
			defer base.DeleteTenancy(auth, tenancyId)
		}
	}
}

func TestTenancySelect(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/tenancy/getTenancySelect"
	base.Get(auth, url, nil, http.StatusOK, "获取成功")
}
