package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestCategoryList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/productCategory/getProductCategoryList"
	base.GetList(auth, url, 0, nil, http.StatusOK, "获取成功")
}

func TestCategorySelect(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/productCategory/getProductCategorySelect"
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func TestCategoryProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	categoryId, data := CreateCategory(auth, "数码产品", http.StatusOK, "创建成功")
	if categoryId == 0 {
		t.Errorf("添加分类失败")
		return
	}
	defer DeleteCategory(auth, categoryId, http.StatusOK, "删除成功")
	{
		keys := base.ResponseKeys{
			{Key: "id", Value: categoryId},
			{Key: "cateName", Value: data["cateName"]},
			{Key: "status", Value: data["status"]},
			{Key: "path", Value: data["path"]},
			{Key: "sort", Value: data["sort"]},
			{Key: "pid", Value: data["pid"]},
			{Key: "pic", Value: data["pic"]},
			{Key: "level", Value: data["level"]},
		}
		url := "v1/admin/productCategory/getProductCategoryList"
		base.GetList(auth, url, 0, keys, http.StatusOK, "获取成功")
	}

	update := map[string]interface{}{
		"cateName": "家电",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	{
		url := fmt.Sprintf("v1/admin/productCategory/updateProductCategory/%d", categoryId)
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}

	{
		url := fmt.Sprintf("v1/admin/productCategory/getProductCategoryById/%d", categoryId)
		keys := base.ResponseKeys{
			{Key: "id", Value: categoryId},
			{Key: "cateName", Value: update["cateName"]},
			{Key: "status", Value: update["status"]},
			{Key: "path", Value: update["path"]},
			{Key: "sort", Value: update["sort"]},
			{Key: "pid", Value: update["pid"]},
			{Key: "pic", Value: update["pic"]},
			{Key: "level", Value: update["level"]},
		}
		base.GetById(auth, url, categoryId, nil, keys, http.StatusOK, "操作成功")
	}

	{
		url := "v1/admin/productCategory/changeProductCategoryStatus"
		base.Post(auth, url, map[string]interface{}{"id": categoryId, "status": g.StatusTrue}, http.StatusOK, "设置成功")
	}

	{
		url := "v1/admin/productCategory/getCreateProductCategoryMap"
		base.Get(auth, url, http.StatusOK, "获取成功")
	}

	{
		url := fmt.Sprintf("v1/admin/productCategory/getUpdateProductCategoryMap/%d", categoryId)
		base.Get(auth, url, http.StatusOK, "获取成功")
	}

}

func TestCategoryRegisterError(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	msg := "Key: 'ProductCategory.BaseProductCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	categoryId, _ := CreateCategory(auth, "", http.StatusBadRequest, msg)
	if categoryId == 0 {
		return
	}
	defer DeleteCategory(auth, categoryId, http.StatusOK, "删除成功")
}

func CreateCategory(auth *httpexpect.Expect, cateName string, status int, message string) (uint, map[string]interface{}) {
	create := map[string]interface{}{
		"cateName": cateName,
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	url := "v1/admin/productCategory/createProductCategory"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId(), create
}

func DeleteCategory(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/admin/productCategory/deleteProductCategory/%d", id)
	base.Delete(auth, url, status, message)
}
