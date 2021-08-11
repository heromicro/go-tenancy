package base

import (
	"net/http"

	"github.com/gavv/httpexpect"
)

func createBrandCategory(auth *httpexpect.Expect, create map[string]interface{}) uint {
	obj := auth.POST("v1/admin/brandCategory/createBrandCategory").
		WithJSON(create).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
	return uint(obj.Value("data").Object().Value("id").Number().Raw())
}
