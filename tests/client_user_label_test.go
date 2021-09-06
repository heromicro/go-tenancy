package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientUserLabelList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/userLabel/getLabelList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}

func TestClientUserLabelProcess(t *testing.T) {
	data := map[string]interface{}{
		"labelName": "sdfsdfs34234",
	}
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/userLabel/createUserLabel").
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
		obj = auth.PUT(fmt.Sprintf("v1/merchant/userLabel/updateUserLabel/%d", int(userLabelId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		userLabel = obj.Value("data").Object()

		userLabel.Value("id").Number().Ge(0)
		userLabel.Value("labelName").String().Equal(update["labelName"].(string))

		obj = auth.GET("v1/merchant/userLabel/getCreateUserLabelMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/merchant/userLabel/getUpdateUserLabelMap/%d", int(userLabelId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUserLabel
		obj = auth.DELETE(fmt.Sprintf("v1/merchant/userLabel/deleteUserLabel/%d", int(userLabelId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
