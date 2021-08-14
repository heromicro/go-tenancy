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
		create := map[string]interface{}{
			"name":       "物流邮费模板",
			"type":       2,
			"appoint":    2,
			"undelivery": 2,
			"isDefault":  1,
			"sort":       2,
		}
		shipTempId := CreateShippingTemplate(auth, create, http.StatusOK, "创建成功")
		if shipTempId == 0 {
			return
		}
		defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")
	}
	{
		create := map[string]interface{}{
			"name":       "陕西物流邮费模板",
			"type":       2,
			"appoint":    2,
			"undelivery": 2,
			"isDefault":  1,
			"sort":       2,
		}
		shipTempId := CreateShippingTemplate(auth, create, http.StatusOK, "创建成功")
		if shipTempId == 0 {
			return
		}
		defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")
	}

	params := []base.Param{
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": ""},
			ResponseKeys: base.ResponseKeys{
				{Type: "int", Key: "pageSize", Value: 10},
				{Type: "int", Key: "page", Value: 1},
				{Type: "int", Key: "total", Value: 2},
			},
		},
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "物流"},
			ResponseKeys: base.ResponseKeys{
				{Type: "int", Key: "pageSize", Value: 10},
				{Type: "int", Key: "page", Value: 1},
				{Type: "int", Key: "total", Value: 1},
			},
		},
		{
			Args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "陕西"},
			ResponseKeys: base.ResponseKeys{
				{Type: "int", Key: "pageSize", Value: 10},
				{Type: "int", Key: "page", Value: 1},
				{Type: "int", Key: "total", Value: 1},
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
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func TestShippingTemplateProcess(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	create := map[string]interface{}{
		"name":       "物流模板名称",
		"type":       2,
		"appoint":    2,
		"undelivery": 2,
		"isDefault":  2,
		"sort":       2,
	}
	shipTempId := CreateShippingTemplate(auth, create, http.StatusOK, "创建成功")
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
			{Type: "uint", Key: "id", Value: shipTempId},
			{Type: "string", Key: "name", Value: update["name"]},
			{Type: "int", Key: "type", Value: update["type"]},
			{Type: "int", Key: "appoint", Value: update["appoint"]},
			{Type: "int", Key: "undelivery", Value: update["undelivery"]},
			{Type: "int", Key: "isDefault", Value: update["isDefault"]},
			{Type: "int", Key: "sort", Value: update["sort"]},
		}
		base.GetById(auth, url, shipTempId, keys, http.StatusOK, "操作成功")
	}
}

func CreateShippingTemplate(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/merchant/shippingTemplate/createShippingTemplate"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}

func DeleteShippingTemplate(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/shippingTemplate/deleteShippingTemplate/%d", id)
	base.Delete(auth, url, status, message)
}
