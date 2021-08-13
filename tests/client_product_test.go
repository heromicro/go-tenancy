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
	params := []base.Param{
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1"}, length: 3},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "keyword": "领立"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "isGiftBag": "1"}, length: 0},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "cateId": 185}, length: 0},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "tenancyCategoryId": 174}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "2"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "3"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "4"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "5"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "6"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "7"}, length: 1},
	}
	for _, param := range params {
		fmt.Print(param)
		url := "v1/merchant/product/getProductList"
		base.PostList(auth, url, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
	}
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
	{

		auth := base.BaseWithLoginTester(t)
		defer base.BaseLogOut(auth)

		createPid := map[string]interface{}{
			"cateName": "箱包服饰_client",
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
			"cateName": "精品服饰_client",
			"status":   g.StatusTrue,
			"path":     "http://qmplusimg.henrongyi.top/head.png",
			"sort":     1,
			"level":    1,
			"pid":      brandCategoryPid,
		}

		brandCategoryId := CreateBrandCategory(auth, createBrandCategory, http.StatusOK, "创建成功")
		if brandCategoryId == 0 {
			return
		}
		defer DeleteBrandCategory(auth, brandCategoryId)

		createBrand := map[string]interface{}{
			"brandName":       "冈本_client",
			"status":          g.StatusTrue,
			"pic":             "http://qmplusimg.henrongyi.top/head.png",
			"sort":            1,
			"brandCategoryId": brandCategoryId,
		}
		brandId = CreateBrand(auth, createBrand, http.StatusOK, "创建成功")
		if brandId == 0 {
			return
		}
		defer DeleteBrand(auth, brandId)

		{
			data := map[string]interface{}{
				"cateName": "数码产品_client",
				"status":   g.StatusTrue,
				"path":     "http://qmplusimg.henrongyi.top/head.png",
				"sort":     1,
				"level":    1,
				"pid":      1,
				"pic":      "http://qmplusimg.henrongyi.top/head.png",
			}

			cateId = CreateCategory(auth, data, http.StatusOK, "创建成功")
			if cateId == 0 {
				return
			}
			defer DeleteCategory(auth, cateId, http.StatusOK, "删除成功")
		}
	}

	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	{
		create := map[string]interface{}{
			"name":       "物流邮费模板_client",
			"type":       2,
			"appoint":    2,
			"undelivery": 2,
			"isDefault":  1,
			"sort":       2,
		}
		shipTempId = CreateShippingTemplate(auth, create, http.StatusOK, "创建成功")
		if shipTempId == 0 {
			return
		}
		defer DeleteShippingTemplate(auth, shipTempId, http.StatusOK, "删除成功")
	}

	{
		data := map[string]interface{}{
			"cateName": "客户端数码产品_client",
			"status":   g.StatusTrue,
			"path":     "http://qmplusimg.henrongyi.top/head.png",
			"sort":     1,
			"level":    1,
			"pid":      1,
			"pic":      "http://qmplusimg.henrongyi.top/head.png",
		}

		tenancyCategoryId = ClientCreateCategory(auth, data, http.StatusOK, "创建成功")
		if tenancyCategoryId == 0 {
			return
		}
		defer DeleteClientCategory(auth, tenancyCategoryId, http.StatusOK, "删除成功")
	}

	data := map[string]interface{}{
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
		"cateId":    cateId,
		"content":   "<p>是的发生的发sad</p>",
		"image":     "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
		"isGiftBag": 2,
		"isGood":    1,
		"keyword":   "sdfdsfsdfsdf",
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

	productId, _, _ := CreateProduct(auth, data, http.StatusOK, "创建成功")
	if productId == 0 {
		return
	}
	defer DeleteProduct(auth, productId, http.StatusOK, "删除成功")

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
		"cateId":    183,
		"content":   "<p>是的发生的发sadsdfsdfsdf</p>",
		"image":     "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
		"isGiftBag": g.StatusTrue,
		"isGood":    g.StatusFalse,
		"keyword":   "sdfdsfsdfsdf",
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
		base.Update(auth, url, update, http.StatusOK, "更新成功")
	}

	keys := base.ResponseKeys{
		{Type: "uint", Key: "id", Value: productId},
		{Type: "int", Key: "sort", Value: update["sort"]},
		{Type: "int", Key: "specType", Value: update["specType"]},
		{Type: "int", Key: "sysBrandId", Value: update["sysBrandId"]},
		{Type: "array", Key: "tenancyCategoryId", Value: update["tenancyCategoryId"]},
		{Type: "int", Key: "tempId", Value: update["tempId"]},
		{Type: "int", Key: "cateId", Value: update["cateId"]},
		{Type: "string", Key: "storeInfo", Value: update["storeInfo"]},
		{Type: "string", Key: "storeName", Value: update["storeName"]},
		{Type: "string", Key: "unitName", Value: update["unitName"]},
		{Type: "string", Key: "videoLink", Value: update["videoLink"]},
		{Type: "string", Key: "keyword", Value: update["keyword"]},
		{Type: "string", Key: "barCode", Value: update["barCode"]},
		{Type: "string", Key: "sliderImage", Value: update["sliderImage"]},
		{Type: "string", Key: "content", Value: update["content"]},
		{Type: "string", Key: "image", Value: update["image"]},
		{Type: "int", Key: "isGiftBag", Value: update["isGiftBag"]},
		{Type: "int", Key: "isGood", Value: update["isGood"]},
		{Type: "array", Key: "attrValue", Value: update["attrValue"]},
		{Type: "array", Key: "sliderImages", Value: update["sliderImages"]},
	}
	url := fmt.Sprintf("v1/merchant/product/getProductById/%d", productId)
	base.GetById(auth, url, productId, keys, http.StatusOK, "操作成功")

	{
		url := "v1/merchant/product/changeProductIsShow"
		base.Post(auth, url, map[string]interface{}{"id": productId, "isShow": 1}, http.StatusOK, "设置成功")
	}
	{
		url := "v1/merchant/product/changeProductIsShow"
		base.Post(auth, url, map[string]interface{}{"id": productId, "isShow": 1}, http.StatusOK, "设置成功")
	}

	DeleteProduct(auth, productId, http.StatusOK, "删除成功")
	defer RestoreProduct(auth, productId, http.StatusOK, "操作成功")

}

func CreateProduct(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) (uint, []string, int) {
	res := base.ResponseKeys{
		{Type: "uint", Key: "id", Value: uint(0)},
		{Type: "array", Key: "uniques", Value: []string{}},
		{Type: "int", Key: "productType", Value: 0},
	}
	url := "v1/merchant/product/createProduct"
	base.Create(auth, url, create, res, status, message)
	return res.GetId(), res.GetStringArrayValue("uniques"), res.GetIntValue("int")
}

func DeleteProduct(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/product/destoryProduct/%d", id)
	base.Delete(auth, url, status, message)
}

func RestoreProduct(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/merchant/product/restoreProduct/%d", id)
	base.Get(auth, url, status, message)
}
