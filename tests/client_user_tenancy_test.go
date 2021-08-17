package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestTenancyUserList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/user/getAdminList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(1)
}

func TestTenancyLoginUser(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	// changePassword success
	obj := auth.POST("v1/merchant/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("v1/merchant/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(http.StatusBadRequest)
	obj.Value("message").String().Equal("修改失败，原密码与当前账户不符")

	// changeProfile success
	obj = auth.POST("v1/merchant/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "admin", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changeProfile success
	obj = auth.POST("v1/merchant/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "商户管理员", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword success
	obj = auth.POST("v1/merchant/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "456789", "newPassword": "123456", "confirmPassword": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")
}

func TestTenancyUserProcess(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "admin1111", "password": "123456", "ConfirmPassword": "123456", "authorityId": []string{source.TenancyAuthorityId}}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("注册成功")

	user := obj.Value("data").Object()
	user.Value("user_id").Number().Ge(0)
	userId := user.Value("user_id").Number().Raw()
	if userId > 0 {

		// setTenancyInfo
		obj = auth.PUT(fmt.Sprintf("v1/merchant/user/setUserInfo/%d", int(userId))).
			WithJSON(map[string]interface{}{
				"email":       "213213@admin.com",
				"phone":       "13800138222",
				"nickName":    "超级管理员11",
				"username":    "chindeo11",
				"hreaderImg":  "http://qmplusimg.henrongyi.top/head.png",
				"status":      2,
				"authorityId": []string{"998"}}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// updateAdminMap
		obj = auth.GET(fmt.Sprintf("v1/merchant/user/updateAdminMap/%d", int(userId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// changePasswordMap
		obj = auth.GET(fmt.Sprintf("v1/merchant/user/changePasswordMap/%d", int(userId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// changeLoginPasswordMap
		obj = auth.GET("v1/merchant/user/changeLoginPasswordMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// changeLoginPasswordMap
		obj = auth.GET("v1/merchant/user/changeProfileMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUser
		obj = auth.DELETE("v1/merchant/user/deleteUser").
			WithJSON(map[string]interface{}{"id": userId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}
