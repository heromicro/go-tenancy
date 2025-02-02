package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestDeviceProductList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/device/product/getProductList"
	res := map[string]interface{}{"page": 1, "pageSize": 10, "type": "0", "cateId": 0, "keyword": "", "tenancyCategoryId": 0}
	base.PostList(auth, url, res, http.StatusOK, "获取成功", base.PageKeys)
}

func TestDeviceProductDetail(t *testing.T) {

	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId uint
	var unique string
	var productType int32
	var adminAuth, tenancyAuth, deviceAuth *httpexpect.Expect

	adminAuth = base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	tenancyAuth, _ = base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(tenancyAuth)

	deviceAuth = base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(deviceAuth)

	brandCategoryPid, _ := CreateBrandCategory(t, adminAuth, "箱包服饰_product_detail", 0, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(t, adminAuth, "精品服饰_product_detail", brandCategoryPid, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(t, adminAuth, "冈本_product_detail", brandCategoryId, http.StatusOK, "创建成功")
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_product_detail", 0, http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Error("添加分类失败")
		return
	}
	defer DeleteCategory(adminAuth, cateId, http.StatusOK, "删除成功")

	shipTempId, _ = CreateShippingTemplate(tenancyAuth, "ship_temp_name_物流邮费模板", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(tenancyAuth, "客户端数码产品_device_order", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, productData := CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
	if productId == 0 {
		t.Errorf("添加商品失败 商品id:%d ", productId)
		return
	}
	defer DeleteProduct(tenancyAuth, productId, http.StatusOK, "删除成功")

	unique, productType = GetProduct(tenancyAuth, productId, productData)
	if len(unique) == 0 || productType == 0 {
		t.Errorf("添加商品失败规格:%+v,商品类型:%d", unique, productType)
	}

	ChangeProductIsShow(tenancyAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")

	url := fmt.Sprintf("v1/device/product/getProductById/%d", productId)
	keys := base.ResponseKeys{
		{Key: "id", Value: productId},
	}
	base.Get(deviceAuth, url, nil, http.StatusOK, "操作成功", keys)
}
