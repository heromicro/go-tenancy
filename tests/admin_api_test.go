package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestApiList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/api/getApiList"
	pageKeys := base.ResponseKeys{
		{Type: "int", Key: "pageSize", Value: 10},
		{Type: "int", Key: "page", Value: 1},
		{Type: "array", Key: "list", Value: nil},
		{Type: "int", Key: "total", Value: 289},
	}
	base.PostList(auth, url, base.PageRes, pageKeys, http.StatusOK, "获取成功")
}
func TestAllApi(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/api/getAllApis"
	base.GetList(auth, url, 0, nil, http.StatusOK, "获取成功")
}

func TestApiProcess(t *testing.T) {
	create := map[string]interface{}{
		"apiGroup":    "test_api_process",
		"description": "test_api_process",
		"method":      "POST",
		"path":        "test_api_process",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	apiId, apiPath, apiMethod := CreateApi(auth, create, http.StatusOK, "创建成功")
	if apiId == 0 || apiPath == "" || apiMethod == "" {
		t.Errorf("添加 api 失败")
		return
	}
	defer DeleteApi(auth, apiId, apiPath, apiMethod)
	update := map[string]interface{}{
		"id":          apiId,
		"apiGroup":    "update_test_api_process",
		"description": "update_test_api_process",
		"method":      "POST",
		"path":        "update_test_api_process",
	}
	{
		url := fmt.Sprintf("v1/admin/api/updateApi/%d", apiId)
		base.Update(auth, url, update, http.StatusOK, "修改成功")
	}
	{
		url := fmt.Sprintf("v1/admin/api/getApiById/%d", apiId)
		keys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: apiId},
			{Type: "string", Key: "path", Value: update["path"]},
			{Type: "string", Key: "method", Value: update["method"]},
			{Type: "string", Key: "description", Value: update["description"]},
			{Type: "string", Key: "apiGroup", Value: update["apiGroup"]},
		}
		base.GetById(auth, url, apiId, keys, http.StatusOK, "操作成功")
	}

}

func TestApiRegisterError(t *testing.T) {
	create := map[string]interface{}{
		"apiGroup":    "auth",
		"description": "用户注册",
		"method":      "GET",
		"path":        "/v1/auth/logout",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	apiId, apiPath, apiMethod := CreateApi(auth, create, response.BAD_REQUEST_ERROR, "添加失败:存在相同api")
	if apiId == 0 || apiPath == "" || apiMethod == "" {
		return
	}
	defer DeleteApi(auth, apiId, apiPath, apiMethod)
}

func CreateApi(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) (uint, string, string) {
	url := "v1/admin/api/createApi"
	keys := base.ResponseKeys{
		{Type: "uint", Key: "id", Value: uint(0)},
		{Type: "string", Key: "path", Value: ""},
		{Type: "string", Key: "method", Value: ""},
	}
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId(), keys.GetStringValue("path"), keys.GetStringValue("method")
}

func DeleteApi(auth *httpexpect.Expect, id uint, apiPath, apiMethod string) {
	obj := auth.DELETE("v1/admin/api/deleteApi").
		WithJSON(map[string]interface{}{
			"id":     id,
			"path":   apiPath,
			"method": apiMethod,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusOK)
	obj.Value("message").String().Equal("删除成功")
}
