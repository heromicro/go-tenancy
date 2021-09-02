package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestGetAttrTemplateList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/attrTemplate/getAttrTemplateList"
	base.PostList(auth, url, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
}

func TestAttrTemplateProcess(t *testing.T) {
	data := map[string]interface{}{
		"templateName": "fsdaf_data",
		"templateValue": []map[string]interface{}{
			{"value": "value_data", "detail": []string{"S"}},
		},
	}
	auth, tenancyId := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	attrTemplateId := CreateAttrTemplate(auth, data, http.StatusOK, "创建成功")
	if attrTemplateId == 0 {
		t.Errorf("添加商品规格模板失败")
		return
	}
	defer DeleteAttrTemplate(auth, attrTemplateId, http.StatusOK, "删除成功")

	update := map[string]interface{}{
		"templateName": "fsdaf_update",
		"templateValue": []map[string]interface{}{
			{
				"value":  "value",
				"detail": []string{"L"},
			},
		},
	}

	{
		url := fmt.Sprintf("v1/merchant/attrTemplate/updateAttrTemplate/%d", attrTemplateId)
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}

	{
		url := fmt.Sprintf("v1/merchant/attrTemplate/getAttrTemplateById/%d", attrTemplateId)
		keys := base.ResponseKeys{
			{Key: "id", Value: attrTemplateId},
			{Key: "sysTenancyId", Value: tenancyId},
			{Key: "templateName", Value: update["templateName"]},
			{Key: "createdAt", Value: update["createdAt"]},
			{Key: "updatedAt", Value: update["updatedAt"]},
			{Key: "templateValue", Value: []base.ResponseKeys{
				{
					{Key: "value", Value: update["templateValue"].([]map[string]interface{})[0]["value"]},
					{Key: "detail", Value: update["templateValue"].([]map[string]interface{})[0]["detail"].([]string)},
				},
			}},
		}
		base.Get(auth, url, nil, http.StatusOK, "操作成功", keys)
	}

}

func DeleteAttrTemplate(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/attrTemplate/deleteAttrTemplate/%d", id)
	base.Delete(auth, url, status, message)
}

func CreateAttrTemplate(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/merchant/attrTemplate/createAttrTemplate"
	keys := base.ResponseKeys{
		{Key: "id", Value: 0},
	}
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}
