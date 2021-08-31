package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientCategoryList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/productCategory/getProductCategoryList"
	base.GetList(auth, url, 0, nil, http.StatusOK, "获取成功")
}

func TestClientCategorySelect(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/productCategory/getProductCategorySelect"
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func TestGetAdminCategorySelect(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/merchant/productCategory/getAdminProductCategorySelect"
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func TestClientCategoryProcess(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	cateId, _ := ClientCreateCategory(auth, "数码产品", 0, http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Errorf("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(auth, cateId, http.StatusOK, "删除成功")

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
		url := fmt.Sprintf("v1/merchant/productCategory/updateProductCategory/%d", cateId)
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}

	{
		url := fmt.Sprintf("v1/merchant/productCategory/getProductCategoryById/%d", cateId)
		keys := base.ResponseKeys{
			{Key: "id", Value: cateId},
			{Key: "cateName", Value: update["cateName"]},
			{Key: "path", Value: update["path"]},
			{Key: "pic", Value: update["pic"]},
			{Key: "status", Value: update["status"]},
			{Key: "sort", Value: update["sort"]},
			{Key: "pid", Value: update["pid"]},
			{Key: "level", Value: update["level"]},
		}
		base.GetById(auth, url, cateId, nil, keys, http.StatusOK, "操作成功")
	}

	{
		url := "v1/merchant/productCategory/changeProductCategoryStatus"
		data := map[string]interface{}{
			"id":     cateId,
			"status": g.StatusTrue,
		}
		base.Post(auth, url, data, http.StatusOK, "设置成功")
	}

	{
		url := "v1/merchant/productCategory/getCreateProductCategoryMap"
		base.Get(auth, url, http.StatusOK, "获取成功")
	}
	{
		url := fmt.Sprintf("v1/merchant/productCategory/getUpdateProductCategoryMap/%d", cateId)
		base.Get(auth, url, http.StatusOK, "获取成功")
	}

}

func TestClientCategoryRegisterError(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	msg := "Key: 'ProductCategory.BaseProductCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag"
	cateId, _ := ClientCreateCategory(auth, "", 0, http.StatusBadRequest, msg)
	if cateId == 0 {
		return
	}
	defer DeleteClientCategory(auth, cateId, http.StatusOK, "删除成功")
}

func ClientCreateCategory(auth *httpexpect.Expect, cateName string, pid uint, status int, message string) (uint, map[string]interface{}) {
	create := map[string]interface{}{
		"cateName": cateName,
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      pid,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	url := "v1/merchant/productCategory/createProductCategory"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId(), create
}

func DeleteClientCategory(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/productCategory/deleteProductCategory/%d", id)
	base.Delete(auth, url, status, message)
}
