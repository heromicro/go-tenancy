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
	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId uint
	var uniques []string
	var productType int32
	var adminAuth, tenancyAuth, deviceAuth *httpexpect.Expect

	adminAuth = base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	tenancyAuth, _ = base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(tenancyAuth)

	deviceAuth = base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(deviceAuth)

	brandCategoryPid, _ := CreateBrandCategory(adminAuth, "箱包服饰_device_check_order", 0, http.StatusOK, "创建成功")
	if brandCategoryPid == 0 {
		t.Error("添加品牌父分类失败")
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(adminAuth, "精品服饰_device_check_order", brandCategoryPid, http.StatusOK, "创建成功")
	if brandCategoryId == 0 {
		t.Error("添加品牌分类失败")
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(adminAuth, "冈本_device_check_order", brandCategoryId, http.StatusOK, "创建成功")
	if brandId == 0 {
		t.Error("添加品牌失败")
		return
	}
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_device_check_order", http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Error("添加分类失败")
		return
	}
	defer DeleteCategory(adminAuth, cateId, http.StatusOK, "删除成功")

	createShipTemp := map[string]interface{}{
		"name":       "物流邮费模板_client",
		"type":       2,
		"appoint":    2,
		"undelivery": 2,
		"isDefault":  1,
		"sort":       2,
	}
	shipTempId = CreateShippingTemplate(tenancyAuth, createShipTemp, http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	createTenancyCategory := map[string]interface{}{
		"cateName": "客户端数码产品_client",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}

	tenancyCategoryId = ClientCreateCategory(tenancyAuth, createTenancyCategory, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

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

	productId, uniques, productType = CreateProduct(tenancyAuth, createProduct, http.StatusOK, "创建成功")
	if productId == 0 || len(uniques) == 0 || productType == 0 {
		t.Errorf("添加商品失败 商品id:%d 规格:%+v,商品类型:%d", productId, uniques, productType)
		return
	}
	defer DeleteProduct(tenancyAuth, productId, http.StatusOK, "删除成功")

	ChangeProductIsShow(tenancyAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")

	createCart := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": uniques[0], "productId": productId, "productType": productType}
	cartId = CreateCart(deviceAuth, createCart, http.StatusOK, "创建成功")
	if cartId == 0 {
		t.Error("添加购物车失败")
		return
	}
	defer DeleteCart(deviceAuth, map[string]interface{}{"ids": []uint{cartId}}, http.StatusOK, "操作成功")

	url := "v1/device/order/checkOrder"
	base.Post(deviceAuth, url, map[string]interface{}{"ids": []uint{uint(cartId)}}, http.StatusOK, "获取成功")
}

func TestDeviceOrderProcess(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId, orderId uint
	var uniques []string
	var productType int32
	var adminAuth, tenancyAuth, deviceAuth *httpexpect.Expect

	adminAuth = base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	tenancyAuth, _ = base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(tenancyAuth)

	deviceAuth = base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(deviceAuth)
	brandCategoryPid, _ := CreateBrandCategory(adminAuth, "箱包服饰_device_process", 0, http.StatusOK, "创建成功")
	if brandCategoryPid == 0 {
		t.Error("添加品牌分类父分类失败")
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(adminAuth, "精品服饰_device_process", brandCategoryPid, http.StatusOK, "创建成功")
	if brandCategoryId == 0 {
		t.Error("添加品牌分类失败")
		return
	}
	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(adminAuth, "冈本_device_process", brandCategoryId, http.StatusOK, "创建成功")
	if brandId == 0 {
		t.Error("添加品牌失败")
		return
	}
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_device_process", http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Error("添加分类失败")
		return
	}
	defer DeleteCategory(adminAuth, cateId, http.StatusOK, "删除成功")

	shippingTemplateCreate := map[string]interface{}{
		"name":       "物流邮费模板_client",
		"type":       2,
		"appoint":    2,
		"undelivery": 2,
		"isDefault":  1,
		"sort":       2,
	}
	shipTempId = CreateShippingTemplate(tenancyAuth, shippingTemplateCreate, http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	shippingTemplateData := map[string]interface{}{
		"cateName": "客户端数码产品_client",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	tenancyCategoryId = ClientCreateCategory(tenancyAuth, shippingTemplateData, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

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

	productId, uniques, productType = CreateProduct(tenancyAuth, data, http.StatusOK, "创建成功")
	if productId == 0 || len(uniques) == 0 || productType == 0 {
		t.Errorf("添加商品失败 商品id:%d 规格:%+v,商品类型:%d", productId, uniques, productType)
		return
	}
	defer DeleteProduct(tenancyAuth, productId, http.StatusOK, "删除成功")

	ChangeProductIsShow(tenancyAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")

	createCartData := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": uniques[0], "productId": productId, "productType": productType}
	cartId = CreateCart(deviceAuth, createCartData, http.StatusOK, "创建成功")
	if cartId == 0 {
		t.Error("添加购物车失败")
		return
	}
	defer DeleteCart(deviceAuth, map[string]interface{}{"ids": []uint{cartId}}, http.StatusOK, "操作成功")

	createOrderData := map[string]interface{}{"cartIds": []uint{uint(cartId)}, "orderType": 1, "remark": "fsdfsdf "}
	orderId = CreateOrder(deviceAuth, createOrderData, http.StatusOK, "获取成功")
	if orderId == 0 {
		t.Error("添加订单失败")
		return
	}

	getOrderByIdKeys := base.ResponseKeys{
		{Type: "uint", Key: "id", Value: orderId},
	}
	base.GetById(deviceAuth, fmt.Sprintf("v1/device/order/getOrderById/%d", orderId), orderId, nil, getOrderByIdKeys, http.StatusOK, "操作成功")

	payOrderKeys := base.ResponseKeys{
		{Type: "notempty", Key: "qrcode", Value: ""},
	}

	// 重新支付订单
	base.GetById(deviceAuth, fmt.Sprintf("v1/device/order/payOrder/%d", orderId), orderId, map[string]interface{}{"orderType": createOrderData["orderType"]}, payOrderKeys, http.StatusOK, "获取成功")

	// 取消订单
	base.Get(deviceAuth, fmt.Sprintf("v1/device/order/cancelOrder/%d", orderId), http.StatusOK, "操作成功")
}

func CreateOrder(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/device/order/createOrder"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}
