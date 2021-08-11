package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestBrandCategoryList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/admin/brandCategory/getBrandCategoryList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestBrandCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusFalse,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      0,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brandCategory/createBrandCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")
	brandCategoryId := obj.Value("data").Object().Value("id").Number().Raw()

	obj = auth.GET("v1/admin/brandCategory/getBrandCategoryList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(1)
	first := obj.Value("data").Array().First().Object()
	first.Keys().ContainsOnly("id", "pid", "cateName", "status", "path", "sort", "level", "children", "createdAt", "updatedAt")
	first.Value("id").Number().Equal(brandCategoryId)
	first.Value("cateName").String().Equal(data["cateName"].(string))
	first.Value("status").Number().Equal(data["status"].(int))
	first.Value("path").String().Equal(data["path"].(string))
	first.Value("sort").Number().Equal(data["sort"].(int))
	first.Value("pid").Number().Equal(data["pid"].(int))
	first.Value("level").Number().Equal(data["level"].(int))

	if brandCategoryId > 0 {
		update := map[string]interface{}{
			"cateName": "家电",
			"status":   g.StatusTrue,
			"path":     "http://qmplusimg.henrongyi.top/head.png",
			"sort":     2,
			"level":    1,
			"pid":      1,
		}

		obj = auth.PUT(fmt.Sprintf("v1/admin/brandCategory/updateBrandCategory/%d", int(brandCategoryId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/brandCategory/getBrandCategoryById/%d", int(brandCategoryId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		brandCategory := obj.Value("data").Object()

		brandCategory.Value("id").Number().Ge(0)
		brandCategory.Value("cateName").String().Equal(update["cateName"].(string))
		brandCategory.Value("status").Number().Equal(update["status"].(int))
		brandCategory.Value("path").String().Equal(update["path"].(string))
		brandCategory.Value("sort").Number().Equal(update["sort"].(int))
		brandCategory.Value("pid").Number().Equal(update["pid"].(int))
		brandCategory.Value("level").Number().Equal(update["level"].(int))

		obj = auth.POST("v1/admin/brandCategory/changeBrandCategoryStatus").
			WithJSON(map[string]interface{}{
				"id":     brandCategoryId,
				"status": g.StatusTrue,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		obj = auth.GET("v1/admin/brandCategory/getCreateBrandCategoryMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/brandCategory/getUpdateBrandCategoryMap/%d", int(brandCategoryId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteBrandCategory
		obj = auth.DELETE(fmt.Sprintf("v1/admin/brandCategory/deleteBrandCategory/%d", int(brandCategoryId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}

func TestBrandCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/admin/brandCategory/createBrandCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'SysBrandCategory.BaseBrandCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag")

}
