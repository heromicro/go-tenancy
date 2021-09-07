package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientProductList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/merchant/product/getProductList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}

func TestGetClientProductFilter(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.GET("v1/merchant/product/getProductFilter").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

}

func TestClientProductProcess(t *testing.T) {

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

	productId, _, _, _ := CreateProduct(clientAuth, cateId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
	if productId == 0 {
		t.Errorf("添加商品失败")
		return
	}
	defer DeleteProduct(clientAuth, productId, http.StatusOK, "删除成功")

	update := map[string]interface{}{
		"attrValue": []map[string]interface{}{
			{
				"image":        "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
				"barCode":      "",
				"brokerage":    1,
				"brokerageTwo": 1,
				"cost":         1,
				"detail": map[string]interface{}{
					"尺寸": "S",
				},
				"otPrice": 1,
				"price":   1,
				"stock":   1,
				"value0":  "S",
				"volume":  1,
				"weight":  1,
			},
		},
		"cateId":  183,
		"content": "<p>是的发生的发sadsdfsdfsdf</p>",
		"image":   "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
		"isGood":  g.StatusFalse,
		"keyword": "sdfdsfsdfsdf",
		"sliderImages": []string{
			"http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
			"http://127.0.0.1:8089/uploads/file/0701aa317da5a004fbf6111545678a6c_20210702150036.png",
		},
		"sort":              21321,
		"specType":          2,
		"storeInfo":         "的是否是否",
		"storeName":         "是防守打法发",
		"sysBrandId":        3,
		"tempId":            2,
		"tenancyCategoryId": []int{174},
		"unitName":          "放松的方式213123",
		"videoLink":         "sdfsdfsd11",
		"barCode":           "sdfsdfsd11",
	}
	{
		url := fmt.Sprintf("v1/merchant/product/updateProduct/%d", productId)
		base.Update(clientAuth, url, update, http.StatusOK, "更新成功")
	}

	keys := base.ResponseKeys{
		{Key: "id", Value: productId},
		{Key: "sort", Value: update["sort"]},
		{Key: "specType", Value: update["specType"]},
		{Key: "sysBrandId", Value: update["sysBrandId"]},
		{Key: "tenancyCategoryId", Value: update["tenancyCategoryId"]},
		{Key: "tempId", Value: update["tempId"]},
		{Key: "cateId", Value: update["cateId"]},
		{Key: "storeInfo", Value: update["storeInfo"]},
		{Key: "storeName", Value: update["storeName"]},
		{Key: "unitName", Value: update["unitName"]},
		{Key: "videoLink", Value: update["videoLink"]},
		{Key: "keyword", Value: update["keyword"]},
		{Key: "barCode", Value: update["barCode"]},
		{Key: "sliderImage", Value: update["sliderImage"]},
		{Key: "content", Value: update["content"]},
		{Key: "image", Value: update["image"]},
		{Key: "isGood", Value: update["isGood"]},
		{Key: "attrValue", Value: update["attrValue"]},
		{Key: "sliderImages", Value: update["sliderImages"]},
	}
	url := fmt.Sprintf("v1/merchant/product/getProductById/%d", productId)
	base.Get(clientAuth, url, nil, http.StatusOK, "操作成功", keys)

	ChangeProductIsShow(clientAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")
	ChangeProductIsShow(clientAuth, productId, g.StatusFalse, http.StatusOK, "设置成功")

	DeleteProduct(clientAuth, productId, http.StatusOK, "删除成功")
	defer RestoreProduct(clientAuth, productId, http.StatusOK, "操作成功")

}

func CreateProduct(auth *httpexpect.Expect, cateId, brandId, shipTempId, tenancyCategoryId uint, status int, message string) (uint, []string, int32, map[string]interface{}) {
	createProduct := map[string]interface{}{
		"attrValue": []map[string]interface{}{
			{
				"image":        "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
				"barCode":      "",
				"brokerage":    1,
				"brokerageTwo": 1,
				"cost":         1,
				"detail": map[string]interface{}{
					"尺寸": "S",
				},
				"otPrice": 1,
				"price":   1,
				"stock":   1,
				"value":   "S",
				"volume":  1,
				"weight":  1,
			},
		},
		"cateId":  cateId,
		"content": "<p>是的发生的发sad</p>",
		"image":   "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
		"isGood":  1,
		"keyword": "sdfdsfsdfsdf",
		"sliderImages": []string{
			"http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
			"http://127.0.0.1:8089/uploads/file/0701aa317da5a004fbf6111545678a6c_20210702150036.png",
		},
		"sort":              1,
		"specType":          1,
		"storeInfo":         "的是否是否",
		"storeName":         "是防守打法发",
		"sysBrandId":        brandId,
		"tempId":            shipTempId,
		"tenancyCategoryId": []uint{tenancyCategoryId},
		"unitName":          "放松的方式",
		"videoLink":         "sdfsdfsd",
		"barCode":           "sdfsdfsd",
	}
	res := base.ResponseKeys{
		{Key: "id", Value: uint(0)},
		{Key: "uniques", Value: []string{}},
		{Key: "productType", Value: 0},
	}
	url := "v1/merchant/product/createProduct"
	base.Create(auth, url, createProduct, res, status, message)
	return res.GetId(), res.GetStringArrayValue("uniques"), res.GetInt32Value("productType"), createProduct
}

func DeleteProduct(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/product/destoryProduct/%d", id)
	base.Delete(auth, url, status, message)
}

// 上架商品
func ChangeProductIsShow(auth *httpexpect.Expect, id uint, isShow, status int, message string) {
	url := "v1/merchant/product/changeProductIsShow"
	base.Post(auth, url, map[string]interface{}{"id": id, "isShow": isShow}, http.StatusOK, "设置成功")
}

func RestoreProduct(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/product/restoreProduct/%d", id)
	base.Get(auth, url, nil, status, message)
}
