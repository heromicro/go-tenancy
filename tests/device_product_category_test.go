package tests

import (
	"net/http"
	"testing"
)

func TestDeviceProductCategoryList(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/device/productCategory/getProductCategoryList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Array()
	first := data.First().Object()
	first.Keys().ContainsOnly("id",
		"createdAt",
		"updatedAt",
		"pid",
		"path",
		"sort",
		"pic",
		"status",
		"children",
		"cateName",
		"level")
	first.Value("id").Number().Ge(0)

}
