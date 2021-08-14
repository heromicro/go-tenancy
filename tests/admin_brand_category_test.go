package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestBrandCategoryList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/brandCategory/getBrandCategoryList"
	base.GetList(auth, url, 0, nil, http.StatusOK, "获取成功")
}

func TestBrandCategoryProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	create := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusFalse,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      0,
	}

	brandCategoryId := CreateBrandCategory(auth, create, http.StatusOK, "创建成功")
	if brandCategoryId > 0 {
		defer DeleteBrandCategory(auth, brandCategoryId)
		{
			rkeys := base.ResponseKeys{
				{Type: "uint", Key: "id", Value: brandCategoryId},
				{Type: "int", Key: "pid", Value: create["pid"]},
				{Type: "int", Key: "status", Value: create["status"]},
				{Type: "int", Key: "sort", Value: create["sort"]},
				{Type: "int", Key: "level", Value: create["level"]},
				{Type: "string", Key: "cateName", Value: create["cateName"]},
				{Type: "string", Key: "path", Value: create["path"]},
				{Type: "object", Key: "children", Value: nil},
			}
			url := "v1/admin/brandCategory/getBrandCategoryList"
			base.GetList(auth, url, brandCategoryId, rkeys, http.StatusOK, "获取成功")
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
			rkeys := base.ResponseKeys{
				{Type: "uint", Key: "id", Value: brandCategoryId},
				{Type: "int", Key: "pid", Value: update["pid"]},
				{Type: "int", Key: "status", Value: update["status"]},
				{Type: "int", Key: "sort", Value: update["sort"]},
				{Type: "int", Key: "level", Value: update["level"]},
				{Type: "string", Key: "cateName", Value: update["cateName"]},
				{Type: "string", Key: "path", Value: update["path"]},
				{Type: "string", Key: "createdAt", Value: update["createdAt"]},
				{Type: "string", Key: "updatedAt", Value: update["updatedAt"]},
				{Type: "object", Key: "children", Value: nil},
			}
			{
				url := fmt.Sprintf("v1/admin/brandCategory/updateBrandCategory/%d", brandCategoryId)
				base.Update(auth, url, update, http.StatusOK, "更新成功")
			}

			{
				url := fmt.Sprintf("v1/admin/brandCategory/getBrandCategoryById/%d", brandCategoryId)
				base.GetById(auth, url, brandCategoryId, rkeys, http.StatusOK, "操作成功")
			}

			{
				update := map[string]interface{}{"id": brandCategoryId, "status": g.StatusTrue}
				url := "v1/admin/brandCategory/changeBrandCategoryStatus"
				base.Post(auth, url, update, http.StatusOK, "设置成功")
			}

			{
				url := "v1/admin/brandCategory/getCreateBrandCategoryMap"
				base.Get(auth, url, http.StatusOK, "获取成功")
			}

			{
				url := fmt.Sprintf("v1/admin/brandCategory/getUpdateBrandCategoryMap/%d", int(brandCategoryId))
				base.Get(auth, url, http.StatusOK, "获取成功")
			}
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
	messge := "Key: 'SysBrandCategory.BaseBrandCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	brandCategoryId := CreateBrandCategory(auth, create, response.BAD_REQUEST_ERROR, messge)
	if brandCategoryId > 0 {
		defer DeleteBrandCategory(auth, brandCategoryId)
	}
}

func CreateBrandCategory(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/admin/brandCategory/createBrandCategory"
	res := base.IdKeys()
	base.Create(auth, url, create, res, status, message)
	return res.GetId()
}

func DeleteBrandCategory(auth *httpexpect.Expect, id uint) {
	url := fmt.Sprintf("v1/admin/brandCategory/deleteBrandCategory/%d", id)
	defer base.Delete(auth, url, http.StatusOK, "删除成功")
}
