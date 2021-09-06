package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestUserLabelList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/userLabel/getUserLabelList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}

func TestUserLabelProcess(t *testing.T) {
	data := map[string]interface{}{
		"labelName": "sdfsdfs34234",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/userLabel/createUserLabel").
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
		obj = auth.PUT(fmt.Sprintf("v1/admin/userLabel/updateUserLabel/%d", int(userLabelId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		userLabel = obj.Value("data").Object()

		userLabel.Value("id").Number().Ge(0)
		userLabel.Value("labelName").String().Equal(update["labelName"].(string))

		obj = auth.GET("v1/admin/userLabel/getCreateUserLabelMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/userLabel/getUpdateUserLabelMap/%d", int(userLabelId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUserLabel
		obj = auth.DELETE(fmt.Sprintf("v1/admin/userLabel/deleteUserLabel/%d", int(userLabelId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
