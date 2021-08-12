package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestApiList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	base.ApiList(auth)

}
func TestAllApi(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	base.AllApi(auth)
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
	apiId, apiPath, apiMethod := base.CreateApi(auth, create, http.StatusOK, "创建成功")
	if apiId > 0 {

		update := map[string]interface{}{
			"id":          apiId,
			"apiGroup":    "update_test_api_process",
			"description": "update_test_api_process",
			"method":      "POST",
			"path":        "update_test_api_process",
		}
		base.UpdateApi(auth, apiId, update)
		base.GetApi(auth, apiId, update)
		base.DeleteApi(auth, apiId, apiPath, apiMethod)
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
	base.CreateApi(auth, create, response.BAD_REQUEST_ERROR, "添加失败:存在相同api")
}
