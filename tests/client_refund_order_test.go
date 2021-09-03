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
	base.PostList(auth, url, base.PageRes, pageKeys, http.StatusOK, "获取成功")
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
