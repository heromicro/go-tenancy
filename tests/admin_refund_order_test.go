package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestRefundOrderList(t *testing.T) {
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
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}

func refundOrderlist(t *testing.T, params map[string]interface{}, length int) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/refundOrder/getRefundOrderList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize", "stat")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(length)

	if length > 0 {
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(

			"id",
			"createdAt",
			"updatedAt",
			"refundOrderSn",
			"deliveryType",
			"deliveryId",
			"deliveryMark",
			"deliveryPics",
			"deliveryPhone",
			"merDeliveryUser",
			"merDeliveryAddress",
			"phone",
			"mark",
			"merMark",
			"adminMark",
			"pics",
			"refundType",
			"refundMessage",
			"refundPrice",
			"refundNum",
			"failMessage",
			"status",
			"statusTime",
			"isDel",
			"isSystemDel",
			"orderSn",
			"userNickName",
			"tenancyName",
			"isTrader",
			"reconciliationId",
			"orderId",
			"sysUserId",
			"sysTenancyId",
			"activityType",
			"refundProduct",
		)
		first.Value("id").Number().Ge(0)
	}
}
