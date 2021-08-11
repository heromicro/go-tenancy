package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientPatientList(t *testing.T) {
	auth := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/patient/getPatientList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(1)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "createdAt", "updatedAt", "name", "phone", "sex", "age", "locName", "bedNum", "hospitalNo", "disease", "sysTenancyId", "hospitalName")
}
