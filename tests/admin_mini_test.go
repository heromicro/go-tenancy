package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestMiniList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/mini/getMiniList"
	base.PostList(auth, url, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
}

func TestMiniProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":      "中德澳线上点餐商城",
		"appId":     "YJ3s1abt7MAfT6gWVKoDresdfsdf",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfsdfsdfssa",
		"remark":    "中德澳线上点餐商城",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	miniId := CreateMini(auth, data, http.StatusOK, "创建成功")
	if miniId == 0 {
		return
	}
	defer DeleteMini(auth, miniId, http.StatusOK, "删除成功")

	{
		url := "v1/admin/mini/getMiniList"
		pageKeys := base.ResponseKeys{
			{Type: "int", Key: "pageSize", Value: 10},
			{Type: "int", Key: "page", Value: 1},
			{Type: "array", Key: "list", Value: []base.ResponseKeys{
				{
					{Type: "uint", Key: "id", Value: miniId},
					{Type: "string", Key: "name", Value: data["name"]},
					{Type: "string", Key: "appId", Value: data["appId"]},
					{Type: "string", Key: "appSecret", Value: data["appSecret"]},
					{Type: "string", Key: "remark", Value: data["remark"]},
				},
			},
			},
			{Type: "int", Key: "total", Value: 1},
		}
		base.PostList(auth, url, base.PageRes, pageKeys, http.StatusOK, "获取成功")
	}

	update := map[string]interface{}{
		"id":        miniId,
		"name":      "中德澳线上点餐商城1",
		"appId":     "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfbnfgh",
		"remark":    "中德澳线上点餐商城1",
	}

	{
		url := fmt.Sprintf("v1/admin/mini/updateMini/%d", miniId)
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}

	{
		url := fmt.Sprintf("v1/admin/mini/getMiniById/%d", miniId)
		keys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: miniId},
			{Type: "string", Key: "name", Value: update["name"]},
			{Type: "string", Key: "appId", Value: update["appId"]},
			{Type: "string", Key: "appSecret", Value: update["appSecret"]},
			{Type: "string", Key: "remark", Value: update["remark"]},
		}
		base.GetById(auth, url, miniId, keys, http.StatusOK, "操作成功")
	}
}

func TestMiniRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":      "中德澳上线护理商城",
		"appId":     "YJ3s1abt7MAfT6gWVKoD",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AM",
		"remark":    "中德澳上线护理商城",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	miniId := CreateMini(auth, data, http.StatusOK, "创建成功")
	if miniId == 0 {
		return
	}
	defer DeleteMini(auth, miniId, http.StatusOK, "删除成功")

	miniId = CreateMini(auth, data, response.BAD_REQUEST_ERROR, "添加失败:商户名称已被注冊")
	if miniId == 0 {
		return
	}
	defer DeleteMini(auth, miniId, http.StatusOK, "删除成功")
}

func CreateMini(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/admin/mini/createMini"
	keys := base.IdKeys
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}

func DeleteMini(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/admin/mini/deleteMini/%d", id)
	base.Delete(auth, url, status, message)
}
