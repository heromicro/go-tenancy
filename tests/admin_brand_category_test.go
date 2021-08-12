package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestBrandCategoryList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/brandCategory/getBrandCategoryList"
	base.GetList(auth, url, 0, nil, nil, http.StatusOK, "获取成功")
}

func TestBrandCategoryProcess(t *testing.T) {
	create := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusFalse,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      0,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/brandCategory/createBrandCategory"
	brandCategoryId := base.Create(auth, url, create, http.StatusOK, "创建成功")

	rkeys := base.ResponseKeys{
		"id":        "number",
		"pid":       "number",
		"cateName":  "string",
		"status":    "number",
		"path":      "number",
		"sort":      "number",
		"level":     "number",
		"children":  "object",
		"createdAt": "string",
		"updatedAt": "string",
	}
	{
		url := "v1/admin/brandCategory/getBrandCategoryList"
		base.GetList(auth, url, brandCategoryId, create, rkeys, http.StatusOK, "获取成功")
	}

	if brandCategoryId > 0 {
		update := map[string]interface{}{
			"cateName": "家电",
			"status":   g.StatusTrue,
			"path":     "http://qmplusimg.henrongyi.top/head.png",
			"sort":     2,
			"level":    1,
			"pid":      1,
		}
		{
			url := fmt.Sprintf("v1/admin/brandCategory/updateBrandCategory/%d", brandCategoryId)
			base.Update(auth, url, update, brandCategoryId, http.StatusOK, "更新成功")
		}

		{
			url := fmt.Sprintf("v1/admin/brandCategory/getBrandCategoryById/%d", brandCategoryId)
			base.Get(auth, url, update, brandCategoryId, rkeys, http.StatusOK, "操作成功")
		}

		{
			update := map[string]interface{}{"id": brandCategoryId, "status": g.StatusTrue}
			url := "v1/admin/brandCategory/changeBrandCategoryStatus"
			base.Post(auth, url, update, rkeys, http.StatusOK, "设置成功")
		}

		{
			url := "v1/admin/brandCategory/getCreateBrandCategoryMap"
			base.Get(auth, url, update, 0, rkeys, http.StatusOK, "获取成功")
		}

		{
			url := fmt.Sprintf("v1/admin/brandCategory/getUpdateBrandCategoryMap/%d", int(brandCategoryId))
			base.Get(auth, url, update, 0, rkeys, http.StatusOK, "获取成功")
		}

		{
			url := fmt.Sprintf("v1/admin/brandCategory/deleteBrandCategory/%d", int(brandCategoryId))
			base.Delete(auth, url, update, rkeys, http.StatusOK, "删除成功")
		}
	}
}

func TestBrandCategoryRegisterError(t *testing.T) {
	create := map[string]interface{}{
		"cateName": "",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/brandCategory/createBrandCategory"
	messge := "Key: 'SysBrandCategory.BaseBrandCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	base.Create(auth, url, create, response.BAD_REQUEST_ERROR, messge)
}
