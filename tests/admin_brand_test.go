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

func TestBrandList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	brandList(auth, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
}

func TestBrandProcess(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	createPid := map[string]interface{}{
		"cateName": "箱包服饰",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      0,
	}

	brandCategoryPid := CreateBrandCategory(auth, createPid, http.StatusOK, "创建成功")
	if brandCategoryPid == 0 {
		return
	}
	defer DeleteBrandCategory(auth, brandCategoryPid)
	createBrandCategory := map[string]interface{}{
		"cateName": "精品服饰",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      brandCategoryPid,
	}

	brandCategoryId := CreateBrandCategory(auth, createBrandCategory, http.StatusOK, "创建成功")
	if brandCategoryId == 0 {
		t.Errorf("添加品牌分类失败")
		return
	}
	defer DeleteBrandCategory(auth, brandCategoryId)
	createBrand := map[string]interface{}{
		"brandName":       "冈本",
		"status":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            1,
		"brandCategoryId": brandCategoryId,
	}
	brandId := CreateBrand(auth, createBrand, http.StatusOK, "创建成功")
	if brandId == 0 {
		t.Errorf("添加品牌失败")
		return
	}
	{
		pageRes := map[string]interface{}{"page": 1, "pageSize": 10, "brandCategoryId": brandCategoryPid}
		pageKeys := base.ResponseKeys{
			{Type: "int", Key: "pageSize", Value: 10},
			{Type: "int", Key: "page", Value: 1},
			{Type: "array", Key: "list", Value: nil},
			{Type: "int", Key: "total", Value: 0},
		}
		brandList(auth, pageRes, pageKeys, http.StatusOK, "获取成功")
	}
	{
		pageRes := map[string]interface{}{"page": 1, "pageSize": 10, "brandCategoryId": brandCategoryId}
		pageKeys := base.ResponseKeys{
			{Type: "int", Key: "pageSize", Value: 10},
			{Type: "int", Key: "page", Value: 1},
			{Type: "array", Key: "list", Value: []base.ResponseKeys{
				{
					{Type: "uint", Key: "id", Value: brandId},
					{Type: "int", Key: "status", Value: createBrand["status"]},
					{Type: "int", Key: "sort", Value: createBrand["sort"]},
					{Type: "uint", Key: "brandCategoryId", Value: createBrand["brandCategoryId"]},
					{Type: "string", Key: "brandName", Value: createBrand["brandName"]},
					{Type: "string", Key: "pic", Value: createBrand["pic"]},
					{Type: "string", Key: "createdAt", Value: createBrand["createdAt"]},
					{Type: "string", Key: "updatedAt", Value: createBrand["updatedAt"]},
				},
			},
			},
			{Type: "int", Key: "total", Value: 1},
		}
		brandList(auth, pageRes, pageKeys, http.StatusOK, "获取成功")
	}

	updateBrand := map[string]interface{}{
		"brandName":       "威尔刚",
		"status":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            2,
		"brandCategoryId": brandCategoryId,
	}
	url := fmt.Sprintf("v1/admin/brand/updateBrand/%d", brandId)
	base.Update(auth, url, updateBrand, http.StatusOK, "更新成功")

	{
		responseKeys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: brandId},
			{Type: "int", Key: "status", Value: updateBrand["status"]},
			{Type: "int", Key: "sort", Value: updateBrand["sort"]},
			{Type: "uint", Key: "brandCategoryId", Value: updateBrand["brandCategoryId"]},
			{Type: "string", Key: "brandName", Value: updateBrand["brandName"]},
			{Type: "string", Key: "pic", Value: updateBrand["pic"]},
		}
		url := fmt.Sprintf("v1/admin/brand/getBrandById/%d", brandId)
		base.GetById(auth, url, brandId, responseKeys, http.StatusOK, "操作成功")
	}

	{
		update := map[string]interface{}{
			"id":     brandId,
			"status": g.StatusTrue,
		}
		url := "v1/admin/brand/changeBrandStatus"
		base.Post(auth, url, update, http.StatusOK, "设置成功")
	}

	{
		url := "v1/admin/brand/getCreateBrandMap"
		base.Get(auth, url, http.StatusOK, "获取成功")
	}
	{
		url := fmt.Sprintf("v1/admin/brand/getUpdateBrandMap/%d", brandId)
		base.Get(auth, url, http.StatusOK, "获取成功")
	}
}

func TestBrandRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"brandName":       "",
		"status":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            2,
		"brandCategoryId": 0,
	}
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	msg := "Key: 'SysBrand.BrandName' Error:Field validation for 'BrandName' failed on the 'required' tag"
	brandId := CreateBrand(auth, data, response.BAD_REQUEST_ERROR, msg)
	if brandId == 0 {
		return
	}
	defer DeleteBrand(auth, brandId)
}

func brandList(auth *httpexpect.Expect, pageRes map[string]interface{}, pageKeys base.ResponseKeys, status int, message string) {
	url := "v1/admin/brand/getBrandList"
	base.PostList(auth, url, pageRes, pageKeys, status, message)
}

func CreateBrand(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/admin/brand/createBrand"
	res := base.IdKeys()
	base.Create(auth, url, create, res, status, message)
	return res.GetId()
}

func DeleteBrand(auth *httpexpect.Expect, id uint) {
	url := fmt.Sprintf("v1/admin/brand/deleteBrand/%d", id)
	defer base.Delete(auth, url, http.StatusOK, "删除成功")
}
