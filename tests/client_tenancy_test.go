package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestGetTenancyInfo(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/merchant/tenancy/getTenancyInfo").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestUpdateClientTenancy(t *testing.T) {
	auth, tenancyId := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	data := map[string]interface{}{
		"avatar": "http://127.0.0.1:8089/uploads/file/49989c75324ef71956c91e79ae49b10d.jpg",
		"banner": "http://127.0.0.1:8089/uploads/def/20200908/c7837d662fd8bd31a8461f7f32e138ce.jpg",
		"info":   "多商户解决方案是为了有效的管理自己的平台及平台入驻商家而提出的方案。",
		"state":  1,
		"tele":   "15109234132",
	}
	obj := auth.PUT(fmt.Sprintf("v1/merchant/tenancy/updateTenancy/%d", tenancyId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
}
