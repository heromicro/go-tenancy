package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestProductList(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/admin/product/getProductList"
	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "total", Value: 0},
	}
	base.PostList(auth, url, base.PageRes, pageKeys, http.StatusOK, "获取成功")
}
func TestGetProductFilter(t *testing.T) {
	auth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/admin/product/getProductFilter"
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func TestProductProcess(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId uint

	adminAuth := base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	brandCategoryPid, _ := CreateBrandCategory(adminAuth, "箱包服饰_client", 0, http.StatusOK, "创建成功")
	if brandCategoryPid == 0 {
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(adminAuth, "精品服饰_client", brandCategoryPid, http.StatusOK, "创建成功")
	if brandCategoryId == 0 {
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(adminAuth, "冈本_client", brandCategoryId, http.StatusOK, "创建成功")
	if brandId == 0 {
		t.Errorf("添加品牌失败")
		return
	}
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

	productId, _, _, _ := CreateProduct(clientAuth, cateId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
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
