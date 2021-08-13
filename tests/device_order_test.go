package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestDeviceOrderList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/device/order/getOrderList"
	base.PostList(auth, url, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
}

func TestDeviceCheckOrder(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId, productId uint
	var uniques []string
	var productType int
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
	{
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

		productId, uniques, productType = CreateProduct(auth, data, http.StatusOK, "创建成功")
		if productId == 0 || len(uniques) == 0 || productType == 0 {
			return
		}

		defer DeleteProduct(auth, productId, http.StatusOK, "删除成功")
	}

	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	create := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": uniques[0], "productId": productId, "productType": productType}
	cartId := CreateCart(auth, create, http.StatusOK, "创建成功")
	if cartId == 0 {
		return
	}
	defer DeleteCart(auth, map[string]interface{}{"ids": []uint{cartId}}, http.StatusOK, "删除成功")

	{
		url := "v1/device/order/checkOrder"
		base.Post(auth, url, map[string]interface{}{"cartIds": []uint{uint(cartId)}}, http.StatusOK, "获取成功")
	}
}

func TestDeviceOrderProcess(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId uint
	var uniques []string
	var productType int
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
	{
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

		productId, uniques, productType = CreateProduct(auth, data, http.StatusOK, "创建成功")
		if productId == 0 || len(uniques) == 0 || productType == 0 {
			return
		}

		defer DeleteProduct(auth, productId, http.StatusOK, "删除成功")
	}

	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	{
		create := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": uniques[0], "productId": productId, "productType": productType}
		cartId = CreateCart(auth, create, http.StatusOK, "创建成功")
		if cartId == 0 {
			return
		}
		defer DeleteCart(auth, map[string]interface{}{"ids": []uint{cartId}}, http.StatusOK, "删除成功")
	}

	create := map[string]interface{}{"cartIds": []uint{uint(cartId)}, "orderType": 1, "remark": "fsdfsdf "}
	orderId := CreateOrder(auth, create, http.StatusOK, "创建成功")
	if orderId == 0 {
		return
	}

	{
		url := fmt.Sprintf("v1/device/order/getOrderById/%d", orderId)
		keys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: orderId},
		}
		base.GetById(auth, url, orderId, keys, http.StatusOK, "操作成功")
	}

	{
		url := fmt.Sprintf("v1/device/order/payOrder/%d", orderId)
		keys := base.ResponseKeys{
			{Type: "uint", Key: "id", Value: orderId},
			{Type: "string", Key: "qrcode", Value: "21312"},
		}
		base.GetById(auth, url, orderId, keys, http.StatusOK, "操作成功")
	}

	{
		url := fmt.Sprintf("v1/device/order/cancelOrder/%d", orderId)
		base.Get(auth, url, http.StatusOK, "操作成功")
	}
}

func CreateOrder(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/device/order/createOrder"
	keys := base.IdKeys
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}
