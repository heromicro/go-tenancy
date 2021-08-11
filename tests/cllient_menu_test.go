package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientMenu(t *testing.T) {
	auth := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/merchant/menu/getMenu").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Array()
	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"authoritys",
		"children",
		"createdAt",
		"hidden",
		"icon",
		"id",
		"is_menu",
		"is_tenancy",
		"menu_id",
		"menu_name",
		"params",
		"path",
		"pid",
		"route",
		"sort",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)
}
