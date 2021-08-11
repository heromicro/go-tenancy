package base

import (
	"net/http"

	"github.com/gavv/httpexpect"
)

func apiList(auth *httpexpect.Expect) {
	obj := auth.POST("v1/admin/api/getApiList").
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
	first.Keys().ContainsOnly("id", "path", "description", "apiGroup", "method", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func allApi(auth *httpexpect.Expect) {
	obj := auth.POST("v1/admin/api/getAllApis").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object().Value("apis").Array()
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"path",
		"description",
		"apiGroup",
		"method",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)
}

func createApi(auth *httpexpect.Expect, create map[string]interface{}) (uint, string, string) {
	obj := auth.POST("v1/admin/api/createApi").
		WithJSON(create).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
	api := obj.Value("data").Object()
	apiId := api.Value("id").Number().Raw()
	apiPath := api.Value("path").String().Raw()
	apiMethod := api.Value("method").String().Raw()
	return uint(apiId), apiPath, apiMethod
}

func updateApi(auth *httpexpect.Expect, apiId uint, update map[string]interface{}) {
	obj := auth.POST("v1/admin/api/updateApi").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")
}

func getApi(auth *httpexpect.Expect, apiId uint, update map[string]interface{}) {
	obj := auth.POST("v1/admin/api/getApiById").
		WithJSON(map[string]interface{}{"id": apiId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	api := obj.Value("data").Object().Value("api").Object()

	api.Value("id").Number().Ge(0)
	api.Value("path").String().Equal(update["path"].(string))
	api.Value("description").String().Equal(update["description"].(string))
	api.Value("apiGroup").String().Equal(update["apiGroup"].(string))
	api.Value("method").String().Equal(update["method"].(string))
}

func deleteApi(auth *httpexpect.Expect, apiId uint, apiPath, apiMethod string) {
	// setUserAuthority
	obj := auth.DELETE("v1/admin/api/deleteApi").
		WithJSON(map[string]interface{}{
			"id":     apiId,
			"path":   apiPath,
			"method": apiMethod,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")
}
