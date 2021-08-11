package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestDeviceOrderList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/device/order/getOrderList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 20}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestDeviceCheckOrder(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/device/cart/createCart").
		WithJSON(map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": "e2fe28308fd2", "productId": 1, "productType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
	cartId := obj.Value("data").Object().Value("id").Number().Raw()
	if cartId > 0 {
		obj = auth.POST("v1/device/order/checkOrder").
			WithJSON(map[string]interface{}{"cartIds": []uint{uint(cartId)}}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")
	}
}

func TestDeviceCreateOrder(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/device/cart/createCart").
		WithJSON(map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": "e2fe28308fd2", "productId": 1, "productType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
	cartId := obj.Value("data").Object().Value("id").Number().Raw()
	if cartId > 0 {
		obj = auth.POST("v1/device/order/createOrder").
			WithJSON(map[string]interface{}{"cartIds": []uint{uint(cartId)}, "orderType": 1, "remark": "fsdfsdf "}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")
		orderId := obj.Value("data").Object().Value("orderId").Number().Raw()
		if orderId > 0 {

			obj = auth.GET(fmt.Sprintf("v1/device/order/getOrderById/%d", int(orderId))).
				Expect().Status(http.StatusOK).JSON().Object()
			obj.Keys().ContainsOnly("status", "data", "message")
			obj.Value("status").Number().Equal(200)
			obj.Value("message").String().Equal("操作成功")

			obj = auth.GET(fmt.Sprintf("v1/device/order/payOrder/%d", int(orderId))).
				WithQuery("orderType", 1).
				Expect().Status(http.StatusOK).JSON().Object()
			obj.Keys().ContainsOnly("status", "data", "message")
			obj.Value("status").Number().Equal(200)
			obj.Value("message").String().Equal("获取成功")

			obj = auth.GET(fmt.Sprintf("v1/device/order/cancelOrder/%d", int(orderId))).
				Expect().Status(http.StatusOK).JSON().Object()
			obj.Keys().ContainsOnly("status", "data", "message")
			obj.Value("status").Number().Equal(200)
			obj.Value("message").String().Equal("操作成功")

		}
	}
}
