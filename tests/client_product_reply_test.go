package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientProductReplyList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "", "nickname": ""}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": ""}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "2", "nickname": ""}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": "B"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "yesterday", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 1, "keyword": "1", "nickname": "C"}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 2, "keyword": "1", "nickname": "C"}, length: 0},
	}
	for _, param := range params {
		productReplyClientlist(t, param.args, param.length)
	}
}

func productReplyClientlist(t *testing.T, params map[string]interface{}, length int) {
	auth := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/productReply/getProductReplyList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
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
			"productScore",
			"serviceScore",
			"postageScore",
			"rate",
			"comment",
			"pics",
			"merchantReplyContent",
			"merchantReplyTime",
			"isReply",
			"isVirtual",
			"avatar",
			"productId",
			"sysUserId",
			"nickname",
			"storeName",
			"image",
			"images",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestClientProductReply(t *testing.T) {
	auth := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/merchant/productReply/replyMap/1").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST("v1/merchant/productReply/reply/1").
		WithJSON(map[string]interface{}{"content": "pageSize"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

}
