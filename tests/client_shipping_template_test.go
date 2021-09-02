package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestShippingTemplateList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	{

		shipTempId, _ := CreateShippingTemplate(auth, "物流邮费模板_templist", http.StatusOK, "创建成功")
		if shipTempId == 0 {
			return
		}
		defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")
	}
	{
		shipTempId, _ := CreateShippingTemplate(auth, "陕西物流邮费模板", http.StatusOK, "创建成功")
		if shipTempId == 0 {
			return
		}
		defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")
	}

	params := []base.Param{
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": ""},
			ResponseKeys: base.ResponseKeys{
				{Key: "pageSize", Value: 10},
				{Key: "page", Value: 1},
				{Key: "total", Value: 2},
			},
		},
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "物流"},
			ResponseKeys: base.ResponseKeys{
				{Key: "pageSize", Value: 10},
				{Key: "page", Value: 1},
				{Key: "total", Value: 1},
			},
		},
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "陕西"},
			ResponseKeys: base.ResponseKeys{
				{Key: "pageSize", Value: 10},
				{Key: "page", Value: 1},
				{Key: "total", Value: 1},
			},
		},
	}
	for _, param := range params {
		url := "v1/merchant/shippingTemplate/getShippingTemplateList"
		base.PostList(auth, url, param.Args, param.ResponseKeys, http.StatusOK, "获取成功")
	}
}

func TestShippingTemplateSelect(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/shippingTemplate/getShippingTemplateSelect"
	base.Get(auth, url, nil, http.StatusOK, "获取成功")
}

func TestShippingTemplateProcess(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	shipTempId, _ := CreateShippingTemplate(auth, "物流模板名称", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Errorf("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")

	update := map[string]interface{}{
		"name":       "物流模板名称_update",
		"type":       1,
		"appoint":    1,
		"undelivery": 1,
		"isDefault":  1,
		"sort":       1,
	}
	{
		url := fmt.Sprintf("v1/merchant/shippingTemplate/updateShippingTemplate/%d", shipTempId)
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}
	{
		url := fmt.Sprintf("v1/merchant/shippingTemplate/getShippingTemplateById/%d", shipTempId)
		keys := base.ResponseKeys{
			{Key: "id", Value: shipTempId},
			{Key: "name", Value: update["name"]},
			{Key: "type", Value: update["type"]},
			{Key: "appoint", Value: update["appoint"]},
			{Key: "undelivery", Value: update["undelivery"]},
			{Key: "isDefault", Value: update["isDefault"]},
			{Key: "sort", Value: update["sort"]},
		}
		base.Get(auth, url, nil, http.StatusOK, "操作成功", keys)
	}
}

func CreateShippingTemplate(auth *httpexpect.Expect, name string, status int, message string) (uint, map[string]interface{}) {
	create := map[string]interface{}{
		"name":       name,
		"type":       2,
		"appoint":    2,
		"undelivery": 2,
		"isDefault":  1,
		"sort":       2,
	}
	url := "v1/merchant/shippingTemplate/createShippingTemplate"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId(), create
}

func DeleteShippingTemplate(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/shippingTemplate/deleteShippingTemplate/%d", id)
	base.Delete(auth, url, status, message)
}
