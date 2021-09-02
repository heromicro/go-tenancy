package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientRefundOrderList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "stat", Value: nil},
		{Key: "total", Value: 0},
	}
	url := "v1/merchant/refundOrder/getRefundOrderList"
	base.PostList(auth, url, base.PageRes, pageKeys, http.StatusOK, "获取成功")
}

func TestClientRefundOrderRecord(t *testing.T) {
	orderId := 1
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderRecord/%d", orderId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
}

func TestClientRefundOrderRemark(t *testing.T) {
	orderId := 1
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderRemarkMap/%d", orderId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/refundOrder/remarkRefundOrder/%d", orderId)).
		WithJSON(map[string]interface{}{"mer_mark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestClientRefundOrderAudit(t *testing.T) {
	orderId := 1
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderMap/%d", orderId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/refundOrder/auditRefundOrder/%d", orderId)).
		WithJSON(map[string]interface{}{"status": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
