package tests

import (
	"net/http"
	"testing"
)

func TestDeviceCartList(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/device/cart/getCartList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total")
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("sysTenancyId", "name", "Avatar", "productId", "products")
}

func TestDeviceCreateCart(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/device/cart/createCart").
		WithJSON(map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": "e2fe28308fd2", "productId": 1, "productType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
}
