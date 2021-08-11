package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestBrandList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brand/getBrandList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
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

func TestBrandListWithCategoryId(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brand/getBrandList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10, "brandCategoryId": 2}).
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

type brandProcess struct {
	create map[string]interface{}
	update map[string]interface{}
}

func TestBrandProcess(t *testing.T) {
	params := []brandProcess{
		{
			create: map[string]interface{}{
				"brandName":       "冈本",
				"status":          g.StatusTrue,
				"pic":             "http://qmplusimg.henrongyi.top/head.png",
				"sort":            1,
				"brandCategoryId": 1,
			},
			update: map[string]interface{}{
				"brandName":       "威尔刚",
				"status":          g.StatusTrue,
				"pic":             "http://qmplusimg.henrongyi.top/head.png",
				"sort":            2,
				"brandCategoryId": 1,
			},
		},
	}
	for _, param := range params {
		baseBrandProcess(t, param)
	}
}

func baseBrandProcess(t *testing.T, param brandProcess) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brand/createBrand").
		WithJSON(param.create).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	obj = auth.POST("v1/admin/brand/getBrandList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(2)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "brandName", "status", "pic", "sort", "brandCategoryId", "createdAt", "updatedAt")

	brand := obj.Value("data").Object()
	brand.Value("id").Number().Ge(0)
	brand.Value("brandName").String().Equal(param.create["brandName"].(string))
	brand.Value("status").Number().Equal(param.create["status"].(int))
	brand.Value("pic").String().Equal(param.create["pic"].(string))
	brand.Value("sort").Number().Equal(param.create["sort"].(int))
	brandId := brand.Value("id").Number().Raw()
	if brandId > 0 {
		obj = auth.PUT(fmt.Sprintf("v1/admin/brand/updateBrand/%d", int(brandId))).
			WithJSON(param.update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/brand/getBrandById/%d", int(brandId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		brand = obj.Value("data").Object()

		brand.Value("id").Number().Ge(0)
		brand.Value("brandName").String().Equal(param.update["brandName"].(string))
		brand.Value("status").Number().Equal(param.update["status"].(int))
		brand.Value("pic").String().Equal(param.update["pic"].(string))
		brand.Value("sort").Number().Equal(param.update["sort"].(int))

		obj = auth.POST("v1/admin/brand/changeBrandStatus").
			WithJSON(map[string]interface{}{
				"id":     brandId,
				"status": g.StatusTrue,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		obj = auth.GET("v1/admin/brand/getCreateBrandMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/brand/getUpdateBrandMap/%d", int(brandId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteBrand
		obj = auth.DELETE(fmt.Sprintf("v1/admin/brand/deleteBrand/%d", int(brandId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}

func TestBrandRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"brandName":       "",
		"status":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            2,
		"brandCategoryId": 1,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brand/createBrand").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'SysBrand.BrandName' Error:Field validation for 'BrandName' failed on the 'required' tag")

}
