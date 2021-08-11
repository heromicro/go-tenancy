package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

type param struct {
	name   string
	args   map[string]interface{}
	length int
}

func TestTenancyList(t *testing.T) {
	ml := 3
	// 当天时间大于 15 号，当月数量为 4
	if time.Now().Day() > 15 {
		ml = 4
	}

	year, month, _ := time.Now().Date()
	date := fmt.Sprintf("%d/%02d/01-%d/%02d/30", year, month, year, month)
	params := []param{
		{name: "no_date", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": ""}, length: 4},
		{name: "today", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "today"}, length: 1},
		{name: "yesterday", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "yesterday"}, length: 1},
		{name: "lately7", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "lately7"}, length: 3},
		{name: "lately30", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "lately30"}, length: 4},
		{name: "month", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "month"}, length: ml},
		{name: "year", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": "year"}, length: 4},
		{name: "status_2", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "2", "keyword": "", "date": ""}, length: 1},
		{name: "status_1", args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1", "keyword": "", "date": date}, length: ml},
	}

	for _, param := range params {
		fmt.Println(param.name)
		list(t, param.args, param.length)
	}
}

func list(t *testing.T, params map[string]interface{}, length int) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/tenancy/getTenancyList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Equal(length)
	first := list.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"sales",
		"serviceScore",
		"Avatar",
		"sort",
		"isAudit",
		"servicePhone",
		"careCount",
		"isBest",
		"tele",
		"address",
		"Keyword",
		"postageScore",
		"mark",
		"name",
		"status",
		"Banner",
		"productScore",
		"State",
		"Info",
		"sysRegionCode",
		"businessTime",
		"isTrader",
		"copyProductNum",
		"createdAt",
		"updatedAt",
		"uuid",
	)
	first.Value("id").Number().Ge(0)
}

func TestTenancyByRegion(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/admin/tenancy/getTenancies/1").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Array().Length().Ge(1)
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
	obj := auth.GET("v1/admin/tenancy/getTenancyCount").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Object().Value("invalid").Equal(1)
	obj.Value("data").Object().Value("valid").Equal(4)
}

func TestTenancyProcess(t *testing.T) {
	data := map[string]interface{}{
		"username":      "bafvetyy",
		"name":          "宝安妇女儿童医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"status":        g.StatusTrue,
		"sysRegionCode": 1,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	tenancy := obj.Value("data").Object()
	tenancy.Value("id").Number().Ge(0)
	tenancyId := tenancy.Value("id").Number().Raw()
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

		obj = auth.PUT(fmt.Sprintf("v1/admin/tenancy/updateTenancy/%d", int(tenancyId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		tenancy = obj.Value("data").Object()

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

		// setUserAuthority
		obj = auth.DELETE(fmt.Sprintf("v1/admin/tenancy/deleteTenancy/%d", int(tenancyId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")

	}
}
func TestTenancyRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"username":      "bafvetyy",
		"name":          "宝安中心人民医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"status":        g.StatusTrue,
		"sysRegionCode": 1,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:商户名称已被注冊")

}

func TestTenancySelect(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/admin/tenancy/getTenancySelect").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

}
