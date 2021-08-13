package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientAutoList(t *testing.T) {
	t.SkipNow()
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/userLabel/auto/getLabelList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(2)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("userNum", "sysTenancyId", "id", "createdAt", "min", "max", "userLabelId", "labelName", "updatedAt", "type")
}

func TestClientAutoProcess(t *testing.T) {
	t.SkipNow()
	data := map[string]interface{}{
		"labelName": "sdfsdfs34234",
	}
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/userLabel/auto/createUserLabel").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	userLabel := obj.Value("data").Object()
	userLabel.Value("id").Number().Ge(0)
	userLabel.Value("labelName").String().Equal(data["labelName"].(string))
	userLabelId := userLabel.Value("id").Number().Raw()
	if userLabelId > 0 {
		update := map[string]interface{}{
			"labelName": "sdfsdfs213213",
		}
		obj = auth.PUT(fmt.Sprintf("v1/merchant/userLabel/auto/updateUserLabel/%d", int(userLabelId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")

		// deleteUserLabel
		obj = auth.DELETE(fmt.Sprintf("v1/merchant/userLabel/auto/deleteUserLabel/%d", int(userLabelId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
