package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestConfigList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/config/getConfigList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
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
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "configName", "configKey", "configType", "required", "value", "configRule", "info", "sort", "userType", "status", "sysConfigCategoryId", "typeName", "createdAt", "updatedAt")
	first.Value("id").Number().Equal(1)
	first.Value("configName").String().Equal("网站域名")
	first.Value("configKey").String().Equal("site_url")
	first.Value("configType").String().Equal("input")
	first.Value("configRule").String().Equal("")
	first.Value("info").String().Equal("")
	first.Value("typeName").String().Equal("文本框")
	first.Value("required").Number().Equal(2)
	first.Value("sort").Number().Equal(0)
	first.Value("userType").Number().Equal(2)
	first.Value("status").Number().Equal(1)
	first.Value("sysConfigCategoryId").Number().Equal(2)
}

func TestGetConfigMap(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/admin/config/getConfigMap/alipay").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestConfigProcess(t *testing.T) {
	data := map[string]interface{}{
		"configKey":           "sdfsdfsdf",
		"configName":          "sdfgdfgdsg",
		"configRule":          "",
		"configType":          "number",
		"info":                "sdafgasdfdsf",
		"required":            1,
		"sort":                1,
		"status":              1,
		"sysConfigCategoryId": 2,
		"userType":            2,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/config/createConfig").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	config := obj.Value("data").Object()
	config.Value("id").Number().Ge(0)

	config.Value("configKey").String().Equal(data["configKey"].(string))
	config.Value("configName").String().Equal(data["configName"].(string))
	config.Value("configRule").String().Equal(data["configRule"].(string))
	config.Value("configType").String().Equal(data["configType"].(string))
	config.Value("info").String().Equal(data["info"].(string))
	config.Value("required").Number().Equal(data["required"].(int))
	config.Value("sort").Number().Equal(data["sort"].(int))
	config.Value("status").Number().Equal(data["status"].(int))
	config.Value("sysConfigCategoryId").Number().Equal(data["sysConfigCategoryId"].(int))
	config.Value("userType").Number().Equal(data["userType"].(int))

	configId := config.Value("id").Number().Raw()
	configKey := config.Value("configKey").String().Raw()
	if configId > 0 {
		// getCreateConfigMap
		obj = auth.GET("v1/admin/config/getCreateConfigMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// getUpdateConfigMap
		obj = auth.GET(fmt.Sprintf("v1/admin/config/getUpdateConfigMap/%d", int(configId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		update := map[string]interface{}{
			"configKey":           "sdfsdfsdf",
			"configName":          "sdfgdfgdsg",
			"configRule":          "",
			"configType":          "number",
			"info":                "sdafgasdfdsf",
			"required":            1,
			"sort":                1,
			"status":              1,
			"sysConfigCategoryId": 2,
			"userType":            2,
		}

		obj = auth.PUT(fmt.Sprintf("v1/admin/config/updateConfig/%d", int(configId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		config = obj.Value("data").Object()

		config.Value("configKey").String().Equal(update["configKey"].(string))
		config.Value("configName").String().Equal(update["configName"].(string))
		config.Value("configRule").String().Equal(update["configRule"].(string))
		config.Value("configType").String().Equal(update["configType"].(string))
		config.Value("info").String().Equal(update["info"].(string))
		config.Value("required").Number().Equal(update["required"].(int))
		config.Value("sort").Number().Equal(update["sort"].(int))
		config.Value("status").Number().Equal(update["status"].(int))
		config.Value("sysConfigCategoryId").Number().Equal(update["sysConfigCategoryId"].(int))
		config.Value("userType").Number().Equal(update["userType"].(int))

		obj = auth.GET("v1/admin/config/getConfigByKey/" + configKey).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		config = obj.Value("data").Object()

		config.Value("id").Number().Ge(0)
		config.Value("configKey").String().Equal(update["configKey"].(string))
		config.Value("configName").String().Equal(update["configName"].(string))
		config.Value("configRule").String().Equal(update["configRule"].(string))
		config.Value("configType").String().Equal(update["configType"].(string))
		config.Value("info").String().Equal(update["info"].(string))
		config.Value("required").Number().Equal(update["required"].(int))
		config.Value("sort").Number().Equal(update["sort"].(int))
		config.Value("status").Number().Equal(update["status"].(int))
		config.Value("sysConfigCategoryId").Number().Equal(update["sysConfigCategoryId"].(int))
		config.Value("userType").Number().Equal(update["userType"].(int))

		obj = auth.GET(fmt.Sprintf("v1/admin/config/getConfigByID/%d", int(configId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		config = obj.Value("data").Object()

		config.Value("id").Number().Ge(0)
		config.Value("configKey").String().Equal(update["configKey"].(string))
		config.Value("configName").String().Equal(update["configName"].(string))
		config.Value("configRule").String().Equal(update["configRule"].(string))
		config.Value("configType").String().Equal(update["configType"].(string))
		config.Value("info").String().Equal(update["info"].(string))
		config.Value("required").Number().Equal(update["required"].(int))
		config.Value("sort").Number().Equal(update["sort"].(int))
		config.Value("status").Number().Equal(update["status"].(int))
		config.Value("sysConfigCategoryId").Number().Equal(update["sysConfigCategoryId"].(int))
		config.Value("userType").Number().Equal(update["userType"].(int))

		// changeConfigStatus
		obj = auth.POST("v1/admin/config/changeConfigStatus").
			WithJSON(map[string]interface{}{
				"id":     configId,
				"status": 2,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// setUserAuthority
		obj = auth.DELETE(fmt.Sprintf("v1/admin/config/deleteConfig/%d", int(configId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}

func TestConfigRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"configKey":           "site_url",
		"configName":          "sdfgdfgdsg",
		"configRule":          "",
		"configType":          "number",
		"info":                "sdafgasdfdsf",
		"required":            1,
		"sort":                1,
		"status":              1,
		"sysConfigCategoryId": 2,
		"userType":            2,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/config/createConfig").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusBadRequest)
	obj.Value("message").String().Equal("添加失败:设置key:site_url已经使用")

}
