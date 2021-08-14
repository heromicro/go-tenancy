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
	data := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	categoryId := CreateCategory(auth, data, http.StatusOK, "创建成功")
	if categoryId == 0 {
		t.Errorf("添加分类失败")
		return
	}
	defer DeleteCategory(auth, categoryId, http.StatusOK, "删除成功")
	{
		keys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: categoryId},
			{Type: "string", Key: "cateName", Value: data["cateName"]},
			{Type: "int", Key: "status", Value: data["status"]},
			{Type: "string", Key: "path", Value: data["path"]},
			{Type: "int", Key: "sort", Value: data["sort"]},
			{Type: "int", Key: "pid", Value: data["pid"]},
			{Type: "string", Key: "pic", Value: data["pic"]},
			{Type: "int", Key: "level", Value: data["level"]},
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
			{Type: "uint", Key: "id", Value: categoryId},
			{Type: "string", Key: "cateName", Value: update["cateName"]},
			{Type: "int", Key: "status", Value: update["status"]},
			{Type: "string", Key: "path", Value: update["path"]},
			{Type: "int", Key: "sort", Value: update["sort"]},
			{Type: "int", Key: "pid", Value: update["pid"]},
			{Type: "string", Key: "pic", Value: update["pic"]},
			{Type: "int", Key: "level", Value: update["level"]},
		}
		base.GetById(auth, url, categoryId, keys, http.StatusOK, "操作成功")
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
	data := map[string]interface{}{
		"cateName": "",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	msg := "Key: 'ProductCategory.BaseProductCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	categoryId := CreateCategory(auth, data, response.BAD_REQUEST_ERROR, msg)
	if categoryId == 0 {
		return
	}
	defer DeleteCategory(auth, categoryId, http.StatusOK, "删除成功")
}

func CreateCategory(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/admin/productCategory/createProductCategory"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}

func DeleteCategory(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/admin/productCategory/deleteProductCategory/%d", id)
	base.Delete(auth, url, status, message)
}
