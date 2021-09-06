package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestUserGroupList(t *testing.T) {

	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/userGroup/getUserGroupList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}

func TestUserGroupProcess(t *testing.T) {
	data := map[string]interface{}{
		"groupName": "sdfsdfs34234",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/userGroup/createUserGroup").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	userGroup := obj.Value("data").Object()
	userGroup.Value("id").Number().Ge(0)
	userGroup.Value("groupName").String().Equal(data["groupName"].(string))
	userGroupId := userGroup.Value("id").Number().Raw()
	if userGroupId > 0 {
		update := map[string]interface{}{
			"groupName": "sdfsdfs213213",
		}
		obj = auth.PUT(fmt.Sprintf("v1/admin/userGroup/updateUserGroup/%d", int(userGroupId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		userGroup = obj.Value("data").Object()

		userGroup.Value("id").Number().Ge(0)
		userGroup.Value("groupName").String().Equal(update["groupName"].(string))

		obj = auth.GET("v1/admin/userGroup/getCreateUserGroupMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/userGroup/getUpdateUserGroupMap/%d", int(userGroupId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUserGroup
		obj = auth.DELETE(fmt.Sprintf("v1/admin/userGroup/deleteUserGroup/%d", int(userGroupId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
