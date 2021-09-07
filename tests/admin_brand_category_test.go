package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestBrandCategoryList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/brandCategory/getBrandCategoryList"
	base.GetList(auth, url, http.StatusOK, "获取成功")
}

func TestBrandCategoryProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	brandCategoryId, create := CreateBrandCategory(t, auth, "数码产品", 0, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(auth, brandCategoryId)
	{
		rkeys := []base.ResponseKeys{
			{
				{Key: "id", Value: brandCategoryId},
				{Key: "pid", Value: create["pid"]},
				{Key: "status", Value: create["status"]},
				{Key: "sort", Value: create["sort"]},
				{Key: "level", Value: create["level"]},
				{Key: "cateName", Value: create["cateName"]},
				{Key: "path", Value: create["path"]},
				{Key: "children", Value: nil},
			},
			nil,
		}
		url := "v1/admin/brandCategory/getBrandCategoryList"
		base.GetList(auth, url, http.StatusOK, "获取成功", rkeys...)
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
			base.Update(auth, url, update, http.StatusOK, "更新成功")
		}

		{
			rkeys := base.ResponseKeys{
				{Key: "id", Value: brandCategoryId},
				{Key: "pid", Value: update["pid"]},
				{Key: "status", Value: update["status"]},
				{Key: "sort", Value: update["sort"]},
				{Key: "level", Value: update["level"]},
				{Key: "cateName", Value: update["cateName"]},
				{Key: "path", Value: update["path"]},
				{Key: "createdAt", Value: update["createdAt"]},
				{Key: "updatedAt", Value: update["updatedAt"]},
				{Key: "children", Value: nil},
			}
			url := fmt.Sprintf("v1/admin/brandCategory/getBrandCategoryById/%d", brandCategoryId)
			base.Get(auth, url, nil, http.StatusOK, "操作成功", rkeys)
		}

		{
			update := map[string]interface{}{"id": brandCategoryId, "status": g.StatusTrue}
			url := "v1/admin/brandCategory/changeBrandCategoryStatus"
			base.Post(auth, url, update, http.StatusOK, "设置成功")
		}

		{
			url := "v1/admin/brandCategory/getCreateBrandCategoryMap"
			base.Get(auth, url, nil, http.StatusOK, "获取成功")
		}

		{
			url := fmt.Sprintf("v1/admin/brandCategory/getUpdateBrandCategoryMap/%d", int(brandCategoryId))
			base.Get(auth, url, nil, http.StatusOK, "获取成功")
		}
	}

}

func TestBrandCategoryRegisterError(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	messge := "Key: 'SysBrandCategory.BaseBrandCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	brandCategoryId, _ := CreateBrandCategory(t, auth, "", 0, http.StatusBadRequest, messge)
	if brandCategoryId > 0 {
		defer DeleteBrandCategory(auth, brandCategoryId)
	}
}

func CreateBrandCategory(t *testing.T, auth *httpexpect.Expect, cateName string, pid uint, status int, message string) (uint, map[string]interface{}) {
	create := map[string]interface{}{
		"cateName": cateName,
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      pid,
	}
	url := "v1/admin/brandCategory/createBrandCategory"
	res := base.IdKeys()
	base.Create(auth, url, create, res, status, message)
	brandCategoryId := res.GetId()
	if brandCategoryId == 0 && cateName != "" {
		t.Fatal("品牌分类添加错误")
	}
	return brandCategoryId, create
}

func DeleteBrandCategory(auth *httpexpect.Expect, id uint) {
	url := fmt.Sprintf("v1/admin/brandCategory/deleteBrandCategory/%d", id)
	defer base.Delete(auth, url, http.StatusOK, "删除成功")
}
