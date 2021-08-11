package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestAdminUserList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/user/getAdminList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
}

func TestAdminLoginUser(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	// changePassword success
	obj := auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("修改失败，原密码与当前账户不符")

	// changeProfile success
	obj = auth.POST("v1/admin/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "admin", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changeProfile success
	obj = auth.POST("v1/admin/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "超级管理员", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword success
	obj = auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "456789", "newPassword": "123456", "confirmPassword": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")
}

func TestAdminUserProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	// registerAdminMap
	obj := auth.GET("v1/admin/user/registerAdminMap").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	obj = auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "chindeo11", "password": "123456", "ConfirmPassword": "123456", "authorityId": []string{source.AdminAuthorityId}}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("注册成功")

	user := obj.Value("data").Object()
	userId := user.Value("user_id").Number().Raw()
	if userId > 0 {
		// setUserAuthority
		obj = auth.POST("v1/admin/user/setUserAuthority").
			WithJSON(map[string]interface{}{"id": userId, "authorityId": source.AdminAuthorityId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("修改成功")

		// setUserInfo
		obj = auth.PUT(fmt.Sprintf("v1/admin/user/setUserInfo/%d", int(userId))).
			WithJSON(map[string]interface{}{
				"email":       "admin@admin.com",
				"phone":       "13800138001",
				"nickName":    "超级管理员",
				"username":    "chindeo",
				"hreaderImg":  "http://qmplusimg.henrongyi.top/head.png",
				"status":      2,
				"authorityId": []string{"999"}}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// changeUserStatus
		obj = auth.POST("v1/admin/user/changeUserStatus", int(userId)).
			WithJSON(map[string]interface{}{
				"id":     userId,
				"status": 2}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("修改成功")

		// updateAdminMap
		obj = auth.GET(fmt.Sprintf("v1/admin/user/updateAdminMap/%d", int(userId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// changePasswordMap
		obj = auth.GET(fmt.Sprintf("v1/admin/user/changePasswordMap/%d", int(userId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUser
		obj = auth.DELETE("v1/admin/user/deleteUser").
			WithJSON(map[string]interface{}{"id": userId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")

	}

}

func TestAdminUserRegisterError(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "ConfirmPassword": "123456", "authorityId": []string{source.AdminAuthorityId}}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名已注册")

}

func TestAdminUserRegisterAuthorityIdEmpty(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "admin_authrity_id_empty", "password": "123456", "ConfirmPassword": "123456", "authorityId": nil}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'Register.AuthorityId' Error:Field validation for 'AuthorityId' failed on the 'required' tag")

}
