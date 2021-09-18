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

func TestDeviceRefundOrderList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/device/refundOrder/getRefundOrderList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}

func TestDeviceRefundOrderProcess(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId, orderId uint
	var unique string
	var productType int32
	var adminAuth, tenancyAuth, deviceAuth *httpexpect.Expect

	adminAuth = base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	tenancyAuth, _ = base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(tenancyAuth)

	deviceAuth = base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(deviceAuth)
	brandCategoryPid, _ := CreateBrandCategory(t, adminAuth, "箱包服饰_device_process", 0, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(t, adminAuth, "精品服饰_device_process", brandCategoryPid, http.StatusOK, "创建成功")

	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(t, adminAuth, "冈本_device_process", brandCategoryId, http.StatusOK, "创建成功")
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_device_process", 0, http.StatusOK, "创建成功")
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

	productId, productData := CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
	if productId == 0 {
		t.Errorf("添加商品失败 商品id:%d", productId)
		return
	}
	defer DeleteProduct(tenancyAuth, productId, http.StatusOK, "删除成功")

	unique, productType = GetProduct(tenancyAuth, productId, productData)
	if len(unique) == 0 || productType == 0 {
		t.Errorf("添加商品失败规格:%+v,商品类型:%d", unique, productType)
	}

	ChangeProductIsShow(tenancyAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")

	createCartData := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": unique, "productId": productId, "productType": productType}
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
	base.ScanById(deviceAuth, fmt.Sprintf("v1/device/order/getOrderById/%d", orderId), nil, http.StatusOK, "操作成功", getOrderByIdKeys)
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
	refundOrder := CreateRefundOrder(deviceAuth, orderId, []uint{orderProductId}, http.StatusOK, "操作成功")
	if refundOrder == 0 {
		t.Error("添加提交退款申请失败")
		return
	}
	base.Get(deviceAuth, fmt.Sprintf("v1/device/refundOrder/getRefundOrderById/%d", refundOrder), nil, http.StatusOK, "操作成功")
	url := "v1/device/refundOrder/getRefundOrderList"
	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "total", Value: 1},
		{Key: "list", Value: []base.ResponseKeys{
			{{Key: "id", Value: 0, Type: "ge"}},
		}},
		{Key: "stat", Value: base.ResponseKeys{
			{Key: "agree", Value: 0},
			{Key: "all", Value: 1},
			{Key: "audit", Value: 0},
			{Key: "backgood", Value: 0},
			{Key: "end", Value: 0},
			{Key: "refuse", Value: 0},
		}},
	}
	base.PostList(deviceAuth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}

func CreateRefundOrder(auth *httpexpect.Expect, orderId uint, ids []uint, status int, message string) uint {
	data := map[string]interface{}{
		"ids":           ids,
		"refundMessage": "地址错了",
		"refundPrice":   1.0,
		"refundType":    1,
		"num":           1,
		"mark":          "",
	}
	keys := base.IdKeys()
	// 提交退款
	base.Create(auth, fmt.Sprintf("v1/device/order/refundOrder/%d", orderId), data, keys, http.StatusOK, "操作成功")
	return keys.GetId()
}
