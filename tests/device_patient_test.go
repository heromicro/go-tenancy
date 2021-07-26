package tests

import (
	"net/http"
	"testing"
)

func TestDevicePatientList(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/device/patient/getPatientList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(-1)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(1)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "createdAt", "updatedAt", "name", "phone", "sex", "age", "locName", "bedNum", "hospitalNo", "disease", "sysTenancyId", "hospitalName")
}
