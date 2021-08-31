package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
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

	brandCategoryPid, _ := CreateBrandCategory(auth, "箱包服饰", 0, http.StatusOK, "创建成功")
	if brandCategoryPid == 0 {
		return
	}
	defer DeleteBrandCategory(auth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(auth, "精品服饰", brandCategoryPid, http.StatusOK, "创建成功")
	if brandCategoryId == 0 {
		t.Errorf("添加品牌分类失败")
		return
	}
	defer DeleteBrandCategory(auth, brandCategoryId)
	brandId, createBrand := CreateBrand(auth, "冈本", brandCategoryId, http.StatusOK, "创建成功")
	if brandId == 0 {
		t.Errorf("添加品牌失败")
		return
	}
	{
		pageRes := map[string]interface{}{"page": 1, "pageSize": 10, "brandCategoryId": brandCategoryPid}
		pageKeys := base.ResponseKeys{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: nil},
			{Key: "total", Value: 0},
		}
		brandList(auth, pageRes, pageKeys, http.StatusOK, "获取成功")
	}
	{
		pageRes := map[string]interface{}{"page": 1, "pageSize": 10, "brandCategoryId": brandCategoryId}
		pageKeys := base.ResponseKeys{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []base.ResponseKeys{
				{
					{Key: "id", Value: brandId},
					{Key: "status", Value: createBrand["status"]},
					{Key: "sort", Value: createBrand["sort"]},
					{Key: "brandCategoryId", Value: createBrand["brandCategoryId"]},
					{Key: "brandName", Value: createBrand["brandName"]},
					{Key: "pic", Value: createBrand["pic"]},
					{Key: "createdAt", Value: createBrand["createdAt"]},
					{Key: "updatedAt", Value: createBrand["updatedAt"]},
				},
			},
			},
			{Key: "total", Value: 1},
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
			{Key: "id", Value: brandId},
			{Key: "status", Value: updateBrand["status"]},
			{Key: "sort", Value: updateBrand["sort"]},
			{Key: "brandCategoryId", Value: updateBrand["brandCategoryId"]},
			{Key: "brandName", Value: updateBrand["brandName"]},
			{Key: "pic", Value: updateBrand["pic"]},
		}
		url := fmt.Sprintf("v1/admin/brand/getBrandById/%d", brandId)
		base.GetById(auth, url, brandId, nil, responseKeys, http.StatusOK, "操作成功")
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

	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	msg := "Key: 'SysBrand.BrandName' Error:Field validation for 'BrandName' failed on the 'required' tag"
	brandId, _ := CreateBrand(auth, "", 0, http.StatusBadRequest, msg)
	if brandId == 0 {
		return
	}
	defer DeleteBrand(auth, brandId)
}

func brandList(auth *httpexpect.Expect, pageRes map[string]interface{}, pageKeys base.ResponseKeys, status int, message string) {
	url := "v1/admin/brand/getBrandList"
	base.PostList(auth, url, pageRes, pageKeys, status, message)
}

func CreateBrand(auth *httpexpect.Expect, brandName string, brandCategoryId uint, status int, message string) (uint, map[string]interface{}) {
	create := map[string]interface{}{
		"brandName":       brandName,
		"status":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            1,
		"brandCategoryId": brandCategoryId,
	}
	url := "v1/admin/brand/createBrand"
	res := base.IdKeys()
	base.Create(auth, url, create, res, status, message)
	return res.GetId(), create
}

func DeleteBrand(auth *httpexpect.Expect, id uint) {
	url := fmt.Sprintf("v1/admin/brand/deleteBrand/%d", id)
	defer base.Delete(auth, url, http.StatusOK, "删除成功")
}
