package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
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

	productId, uniques, productType, _ = CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
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

func TestDeviceOrderProcessForCancelOrder(t *testing.T) {
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

	shipTempId, _ = CreateShippingTemplate(tenancyAuth, "物流邮费模板_device_process", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(tenancyAuth, "device_order_cate_name", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, uniques, productType, _ = CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
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
	defer DeleteClientOrder(tenancyAuth, orderId, http.StatusOK, "删除成功")
	defer DeleteDeviceOrder(deviceAuth, orderId, http.StatusOK, "删除成功")
	getOrderByIdKeys := base.ResponseKeys{
		{Key: "id", Value: orderId},
	}
	base.GetById(deviceAuth, fmt.Sprintf("v1/device/order/getOrderById/%d", orderId), orderId, nil, getOrderByIdKeys, http.StatusOK, "操作成功")

	payOrderKeys := base.ResponseKeys{
		{Key: "qrcode", Value: "notempty"},
	}
	// 重新支付订单
	base.GetById(deviceAuth, fmt.Sprintf("v1/device/order/payOrder/%d", orderId), orderId, map[string]interface{}{"orderType": createOrderData["orderType"]}, payOrderKeys, http.StatusOK, "获取成功")

	// 取消订单
	base.Get(deviceAuth, fmt.Sprintf("v1/device/order/cancelOrder/%d", orderId), http.StatusOK, "操作成功")
}

func TestDeviceOrderProcessForCheckReturnOrder(t *testing.T) {
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

	shipTempId, _ = CreateShippingTemplate(tenancyAuth, "物流邮费模板_device_process", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(tenancyAuth, "device_order_cate_name", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, uniques, productType, _ = CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
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
	defer DeleteClientOrder(tenancyAuth, orderId, http.StatusOK, "删除成功")
	defer DeleteDeviceOrder(deviceAuth, orderId, http.StatusOK, "删除成功")

	// 申请退款
	data := map[string]interface{}{
		"ids": []uint{productId},
	}
	base.Post(deviceAuth, fmt.Sprintf("v1/device/order/checkRefundOrder/%d", orderId), data, http.StatusBadRequest, "操作失败:订单未付款,请取消订单")

	getOrderByIdKeys := base.ResponseKeys{
		{Key: "orderSn", Value: ""},
		{Key: "orderProduct",
			Value: []base.ResponseKeys{
				{
					{Key: "id", Value: 0},
				},
			},
		},
	}
	base.ScanById(deviceAuth, fmt.Sprintf("v1/device/order/getOrderById/%d", orderId), orderId, nil, getOrderByIdKeys, http.StatusOK, "操作成功")
	orderSn := getOrderByIdKeys.GetStringValue("orderSn")
	orderProducts := getOrderByIdKeys.GetResponseKeysValue("orderProduct")
	if len(orderProducts) == 0 {
		t.Error("添加订单失败:订单产品为空")
		return
	}
	orderProductId := orderProducts[0].GetId()
	changeData := map[string]interface{}{
		"status":   model.OrderStatusNoDeliver,
		"pay_type": model.PayTypeAlipay,
		"pay_time": time.Now(),
		"paid":     g.StatusTrue,
	}
	data = map[string]interface{}{
		"ids": []uint{orderProductId},
	}
	_, err := service.ChangeOrderPayNotifyByOrderSn(changeData, orderSn, "pay_success", "订单支付成功")
	if err != nil {
		t.Errorf("%s 订单支付失败%v", orderSn, err.Error())
	}

	// 申请退款
	base.Post(deviceAuth, fmt.Sprintf("v1/device/order/checkRefundOrder/%d", orderId), data, http.StatusOK, "操作成功")
}

func TestDeviceOrderProcessForReturnOrder(t *testing.T) {
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

	shipTempId, _ = CreateShippingTemplate(tenancyAuth, "物流邮费模板_device_process", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(tenancyAuth, "device_order_cate_name", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, uniques, productType, _ = CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
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
	defer DeleteClientOrder(tenancyAuth, orderId, http.StatusOK, "删除成功")
	defer DeleteDeviceOrder(deviceAuth, orderId, http.StatusOK, "删除成功")

	getOrderByIdKeys := base.ResponseKeys{
		{Key: "orderSn", Value: ""},
		{Key: "orderProduct",
			Value: []base.ResponseKeys{
				{
					{Key: "id", Value: 0},
				},
			},
		},
	}
	base.ScanById(deviceAuth, fmt.Sprintf("v1/device/order/getOrderById/%d", orderId), orderId, nil, getOrderByIdKeys, http.StatusOK, "操作成功")
	orderSn := getOrderByIdKeys.GetStringValue("orderSn")
	orderProducts := getOrderByIdKeys.GetResponseKeysValue("orderProduct")
	if len(orderProducts) == 0 {
		t.Error("添加订单失败:订单产品为空")
		return
	}
	orderProductId := orderProducts[0].GetId()
	changeData := map[string]interface{}{
		"status":   model.OrderStatusNoDeliver,
		"pay_type": model.PayTypeAlipay,
		"pay_time": time.Now(),
		"paid":     g.StatusTrue,
	}
	_, err := service.ChangeOrderPayNotifyByOrderSn(changeData, orderSn, "pay_success", "订单支付成功")
	if err != nil {
		t.Errorf("%s 订单支付失败%v", orderSn, err.Error())
	}
	data := map[string]interface{}{
		"ids":           []uint{orderProductId},
		"refundMessage": "地址错了",
		"RefundPrice":   1.0,
		"RefundType":    1,
		"Num":           1,
		"Mark":          "",
	}
	// 提交退款
	base.Post(deviceAuth, fmt.Sprintf("v1/device/order/refundOrder/%d", orderId), data, http.StatusOK, "操作成功")
}

func CreateOrder(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/device/order/createOrder"
	keys := base.IdKeys()
	base.Create(auth, url, create, keys, status, message)
	return keys.GetId()
}

func DeleteDeviceOrder(auth *httpexpect.Expect, id uint, status int, message string) {
	url := fmt.Sprintf("v1/device/order/deleteOrder/%d", id)
	base.Delete(auth, url, status, message)
}
