package source

import (
	"encoding/json"
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

var Order = new(order)

type order struct{}

var groupOrders = []model.GroupOrder{
	{CUserId: 7, GroupOrderSn: g.CreateOrderSn("G"), TotalPostage: 20.00, TotalPrice: 50.00, TotalNum: 1, RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", PayPrice: 50.00, PayPostage: 30.00, Cost: 5.00, Paid: g.StatusFalse, PayTime: &now, PayType: model.PayTypeWx, IsRemind: g.StatusTrue},
	{CUserId: 7, GroupOrderSn: g.CreateOrderSn("G"), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: g.StatusFalse, PayType: model.PayTypeWx, IsRemind: g.StatusTrue, TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now()}}},
	{CUserId: 7, GroupOrderSn: g.CreateOrderSn("G"), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: g.StatusTrue, PayType: model.PayTypeWx, IsRemind: g.StatusTrue},
	{CUserId: 7, GroupOrderSn: g.CreateOrderSn("G"), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: g.StatusTrue, PayType: model.PayTypeWx, IsRemind: g.StatusTrue},
	{CUserId: 7, GroupOrderSn: g.CreateOrderSn("G"), TotalPostage: 4.00, TotalPrice: 352.00, TotalNum: 4, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 356.00, PayPostage: 4.00, Cost: 400.00, Paid: g.StatusTrue, PayType: model.PayTypeWx, IsRemind: g.StatusTrue},
}

var orders = []model.Order{
	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 1, ReconciliationId: 0, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", TotalNum: 10, TotalPrice: 20.00, TotalPostage: 30.00, PayPrice: 50.00, PayPostage: 30.00, CommissionRate: 15.00, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusRefund, DeliveryType: model.DeliverTypeFH, DeliveryName: "delivery_name", DeliveryID: "delivery_id", Mark: "mark", Remark: "remark", AdminMark: "admin_mark", ActivityType: model.GeneralSale, Cost: 5.00, IsCancel: g.StatusFalse}},

	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 2, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusFalse, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusNoReceive, Cost: 100.00, IsCancel: g.StatusTrue}},

	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 2, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusNoReceive, Cost: 100.00, IsCancel: g.StatusFalse}},

	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 3, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusNoComment, Cost: 100.00, IsCancel: g.StatusFalse}},

	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 4, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusFinish, Cost: 100.00, IsCancel: g.StatusFalse}},

	{CUserId: 7, SysTenancyId: 1, GroupOrderId: 5, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.GeneralSale), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 4, TotalPrice: 352.00, TotalPostage: 4.00, PayPrice: 356.00, PayPostage: 4.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: &now, PayType: model.PayTypeWx, Status: model.OrderStatusNoDeliver, Cost: 400.00, IsCancel: g.StatusFalse}},
}

func getOrderProducts() []model.OrderProduct {
	cartInfo := request.CartInfo{
		Product: request.CartInfoProduct{
			Image:     "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg",
			StoreName: "领立裁腰带短袖连衣裙",
		},
		ProductAttr: request.CartInfoProductAttr{
			Price: 50.00,
			Sku:   "L",
		},
	}
	ci, _ := json.Marshal(&cartInfo)

	orderProducts := []model.OrderProduct{
		{OrderId: 1, CUserId: 7, CartId: 1, ProductId: 1, CartInfo: string(ci), BaseOrderProduct: model.BaseOrderProduct{ProductSku: "L", IsRefund: 0, ProductNum: 12, ProductType: model.GeneralSale, RefundNum: 5, IsReply: g.StatusFalse, ProductPrice: 50.00}},

		{OrderId: 1, CUserId: 7, CartId: 1, ProductId: 2, CartInfo: string(ci), BaseOrderProduct: model.BaseOrderProduct{ProductSku: "L", IsRefund: 0, ProductNum: 12, ProductType: model.GeneralSale, RefundNum: 5, IsReply: g.StatusFalse, ProductPrice: 50.00}},

		{OrderId: 2, CUserId: 7, CartId: 2, ProductId: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":99,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 1, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderId: 3, CUserId: 7, CartId: 3, ProductId: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":98,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 1, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 0, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderId: 4, CUserId: 7, CartId: 4, ProductId: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":97,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 1, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderId: 5, CUserId: 7, CartId: 5, ProductId: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":96,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 4, ProductType: model.GeneralSale, RefundNum: 4, IsReply: g.StatusFalse, ProductPrice: 352.00}},
	}
	return orderProducts
}

var orderStatus = []model.OrderStatus{
	{OrderId: 2, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderId: 2, ChangeType: "cancel", ChangeMessage: "取消订单[自动]", ChangeTime: time.Now()},
	{OrderId: 3, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderId: 3, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
	{OrderId: 4, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderId: 4, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
	{OrderId: 5, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderId: 5, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
}

var orderCarts = []model.CartOrder{
	{OrderId: 1, CartId: 1},
	{OrderId: 2, CartId: 2},
	{OrderId: 3, CartId: 3},
	{OrderId: 4, CartId: 4},
	{OrderId: 5, CartId: 5},
	{OrderId: 6, CartId: 6},
	{OrderId: 7, CartId: 7},
}

//@description: orders 表数据初始化
func (a *order) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Order{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> orders 表的初始数据已存在!")
			return nil
		}
		if err := tx.Model(&model.GroupOrder{}).Create(&groupOrders).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&orders).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		orderProducts := getOrderProducts()
		if err := tx.Model(&model.OrderProduct{}).Create(&orderProducts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.OrderStatus{}).Create(&orderStatus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.CartOrder{}).Create(&orderCarts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> orders 表初始数据成功!")
		return nil
	})
}
