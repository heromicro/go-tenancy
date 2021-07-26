package tests

import (
	"net/http"
	"testing"
)

func TestDeviceCheckOrder(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/device/order/checkOrder").
		WithJSON(map[string]interface{}{"cartIds": []uint{1}}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}
