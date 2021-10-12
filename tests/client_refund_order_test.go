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

func TestClientRefundOrderList(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	pageKeys := base.ResponseKeys{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "stat", Value: nil},
		{Key: "total", Value: 0},
	}
	url := "v1/merchant/refundOrder/getRefundOrderList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}

func TestClientRefundOrderRecord(t *testing.T) {
	orderId := 1
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderRecord/%d", orderId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
}

func TestClientRefundOrderRemark(t *testing.T) {
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
		t.Errorf("添加商品失败 商品id:%d ", productId)
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
	_, err := service.ChangeOrderPayNotifyByOrderSn(changeData, orderSn, model.ChangeTypePaySuccess)
	if err != nil {
		t.Errorf("%s 订单支付失败%v", orderSn, err.Error())
	}
	refundOrder := CreateRefundOrder(deviceAuth, orderId, []uint{orderProductId}, http.StatusOK, "操作成功")
	if refundOrder == 0 {
		t.Error("添加提交退款申请失败")
		return
	}
	obj := tenancyAuth.GET(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderRemarkMap/%d", refundOrder)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = tenancyAuth.POST(fmt.Sprintf("v1/merchant/refundOrder/remarkRefundOrder/%d", refundOrder)).
		WithJSON(map[string]interface{}{"mer_mark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestClientRefundOrderAudit(t *testing.T) {
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
		t.Errorf("添加商品失败 商品id:%d ", productId)
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
	_, err := service.ChangeOrderPayNotifyByOrderSn(changeData, orderSn, model.ChangeTypePaySuccess)
	if err != nil {
		t.Errorf("%s 订单支付失败%v", orderSn, err.Error())
	}
	refundOrder := CreateRefundOrder(deviceAuth, orderId, []uint{orderProductId}, http.StatusOK, "操作成功")
	if refundOrder == 0 {
		t.Error("添加提交退款申请失败")
		return
	}
	obj := tenancyAuth.GET(fmt.Sprintf("v1/merchant/refundOrder/getRefundOrderMap/%d", refundOrder)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = tenancyAuth.POST(fmt.Sprintf("v1/merchant/refundOrder/auditRefundOrder/%d", refundOrder)).
		WithJSON(map[string]interface{}{"status": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
