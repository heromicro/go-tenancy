package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func ProductList(auth *httpexpect.Expect, res map[string]interface{}, status int, message string, keys ...base.ResponseKeys) {
	url := "v1/admin/product/getProductList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", keys...)
}

func TestProductList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "total", Value: 0},
	}
	ProductList(auth, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}
func TestGetProductFilter(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/product/getProductFilter"
	pageKeys := []base.ResponseKeys{
		{
			{Key: "type", Value: 1},
			{Key: "name", Value: "出售中"},
			{Key: "count", Value: 0},
		},
		{
			{Key: "type", Value: 2},
			{Key: "name", Value: "仓库中"},
			{Key: "count", Value: 0},
		},
		{
			{Key: "type", Value: 6},
			{Key: "name", Value: "待审核"},
			{Key: "count", Value: 0},
		},
		{
			{Key: "type", Value: 7},
			{Key: "name", Value: "审核未通过"},
			{Key: "count", Value: 0},
		},
	}
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys...)
}

func TestProductProcess(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId uint

	adminAuth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	brandCategoryPid, _ := CreateBrandCategory(t, adminAuth, "箱包服饰_client", 0, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(t, adminAuth, "精品服饰_client", brandCategoryPid, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(t, adminAuth, "冈本_client", brandCategoryId, http.StatusOK, "创建成功")
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_client", 0, http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Errorf("添加分类失败")
		return
	}
	defer DeleteCategory(adminAuth, cateId, http.StatusOK, "删除成功")

	clientAuth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(clientAuth)

	shipTempId, _ = CreateShippingTemplate(clientAuth, "物流邮费模板_client", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Errorf("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(clientAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(clientAuth, "客户端数码产品_client", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Errorf("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(clientAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, _ := CreateProduct(clientAuth, cateId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
	if productId == 0 {
		t.Errorf("添加商品失败")
		return
	}
	defer DeleteProduct(clientAuth, productId, http.StatusOK, "删除成功")

	update := map[string]interface{}{
		"storeName": "领立裁腰带短袖连衣裙",
		"isHot":     2,
		"isBenefit": 2,
		"isBest":    2,
		"isNew":     2,
		"content":   "dsfsafasfasfas",
	}

	obj := adminAuth.PUT(fmt.Sprintf("v1/admin/product/updateProduct/%d", productId)).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")

	obj = adminAuth.POST("v1/admin/product/changeProductStatus").
		WithJSON(map[string]interface{}{"id": []uint{productId}, "status": 3}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	obj = adminAuth.POST("v1/admin/product/changeMutilProductStatus").
		WithJSON(map[string]interface{}{"id": []uint{productId}, "status": 3}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	obj = adminAuth.GET(fmt.Sprintf("v1/admin/product/getProductById/%d", productId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	product := obj.Value("data").Object()

	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(update["storeName"].(string))
	product.Value("isHot").Number().Equal(update["isHot"].(int))
	product.Value("isBenefit").Number().Equal(update["isBenefit"].(int))
	product.Value("isBest").Number().Equal(update["isBest"].(int))
	product.Value("isNew").Number().Equal(update["isNew"].(int))
	product.Value("content").String().Equal(update["content"].(string))
}
