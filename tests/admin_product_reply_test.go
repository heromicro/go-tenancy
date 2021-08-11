package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestProductReplyList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/productReply/getProductReplyList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "", "nickname": ""}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(0)
}

func TestProductReply(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/admin/productReply/replyMap").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST("v1/admin/productReply/reply").
		WithJSON(map[string]interface{}{
			"avatar":       "http://127.0.0.1:8089/uploads/file/bb2bfacd349a34e0644ba3b31ceb03a8_20210810153341.jpg",
			"comment":      "123213123",
			"nickname":     "2312312",
			"pic":          []string{},
			"postageScore": 5,
			"productId": map[string]interface{}{
				"src": "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg",
				"id":  1,
			},
			"productScore": 5,
			"serviceScore": 5,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	productReplyId := obj.Value("data").Object().Value("id").Number().Raw()
	if productReplyId > 0 {
		obj = auth.DELETE(fmt.Sprintf("v1/admin/productReply/deleteProductReply/%d", int(productReplyId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
	}

}
