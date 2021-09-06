package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/skip2/go-qrcode"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils/param"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// DeliveryOrderMap 发货表单
func DeliveryOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := `{"rule":[{"type":"radio","field":"deliveryType","value":1,"title":"发货类型","props":{},"control":[{"value":1,"rule":[{"type":"select","field":"deliveryName","value":"","title":"快递名称","props":{"multiple":false,"placeholder":"请选择快递名称"},"options":[]},{"type":"input","field":"deliveryId","value":"","title":"快递单号","props":{"type":"text","placeholder":"请输入快递单号"},"validate":[{"message":"请输入快递单号","required":true,"type":"string","trigger":"change"}]}]},{"value":2,"rule":[{"type":"input","field":"deliveryName","value":"","title":"送货人姓名","props":{"type":"text","placeholder":"请输入送货人姓名"},"validate":[{"message":"请输入送货人姓名","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"deliveryId","value":"","title":"手机号","props":{"type":"text","placeholder":"请输入手机号"},"validate":[{"message":"请输入手机号","required":true,"type":"string","trigger":"change"}]}]},{"value":3,"rule":[]}],"options":[{"value":1,"label":"发货"},{"value":2,"label":"送货"},{"value":3,"label":"无需物流"}]}],"action":"","method":"POST","title":"添加发货信息","config":{}}`
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/deliveryOrder", id), ctx)

	opts, err := GetExpressOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Controls[0].Rule[0].Options = opts
	return form, err
}

// GetOrderRemarkMap 备注表单
func GetOrderRemarkMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	order, err := GetOrderRemarkAndUpdateByID(id, ctx)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"remark","value":"%s","title":"备注","props":{"type":"text","placeholder":"请输入备注"},"validate":[{"message":"请输入备注","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"修改备注","config":{}}`, order.Remark)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/remarkOrder", id), ctx)
	return form, err
}

// GetEditOrderMap 编辑表单
func GetEditOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	order, err := GetOrderRemarkAndUpdateByID(id, ctx)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"inputNumber","field":"totalPrice","value":%02f,"title":"订单总价","props":{"placeholder":"请输入订单总价"},"validate":[{"message":"请输入订单总价","required":true,"type":"number","trigger":"change"}]},{"type":"inputNumber","field":"payPrice","value":%02f,"title":"实际支付金额","props":{"placeholder":"请输入实际支付金额"},"validate":[{"message":"请输入实际支付金额","required":true,"type":"number","trigger":"change"}]},{"type":"inputNumber","field":"totalPostage","value":%02f,"title":"订单邮费","props":{"placeholder":"请输入订单邮费"},"validate":[{"message":"请输入订单邮费","required":true,"type":"number","trigger":"change"}]}],"action":"","method":"POST","title":"修改订单","config":{}}`, order.TotalPrice, order.PayPrice, order.TotalPostage)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/updateOrder", id), ctx)
	return form, err
}

// GetOrderRemarkAndUpdateByID 获取订单备注和价格
func GetOrderRemarkAndUpdateByID(id uint, ctx *gin.Context) (request.OrderRemarkAndUpdate, error) {
	var order request.OrderRemarkAndUpdate
	db := g.TENANCY_DB.Model(&model.Order{}).Select("remark,total_price,pay_price,total_postage").Where("id = ?", id)
	err := db.First(&order).Error
	return order, err
}

// getOrderCount 订单数量
func getOrderCount(name string, ctx *gin.Context) (int64, error) {
	var count int64
	wheres := getOrderConditions()
	for _, where := range wheres {
		if where.Name == name {
			db := g.TENANCY_DB.Model(&model.Order{})
			if where.Conditions != nil && len(where.Conditions) > 0 {
				for key, cn := range where.Conditions {
					if cn == nil {
						db = db.Where(key)
					} else {
						db = db.Where(fmt.Sprintf("%s = ?", key), cn)
					}
				}
			}
			if multi.IsTenancy(ctx) {
				db = CheckTenancyId(db, multi.GetTenancyId(ctx), "")
			}

			err := db.Count(&count).Error
			if err != nil {
				return count, err
			}
		}
	}

	return count, nil
}

// GetFilter 订单过滤数量
func GetFilter(ctx *gin.Context) ([]map[string]interface{}, error) {
	charts := []map[string]interface{}{
		{"count": 0, "orderType": "", "title": "全部"},
		{"count": 0, "orderType": "1", "title": "普通订单"},
		{"count": 0, "orderType": "2", "title": "核销订单"},
	}

	for _, chart := range charts {
		if chart["orderType"] == "" {
			continue
		}
		var count int64
		db := g.TENANCY_DB.Model(&model.Order{})
		if multi.IsTenancy(ctx) {
			db = CheckTenancyId(db, multi.GetTenancyId(ctx), "")
		}

		err := db.Where("order_type = ?", chart["orderType"]).Count(&count).Error
		if err != nil {
			return nil, err
		} else {
			chart["count"] = count
		}
	}

	return charts, nil
}

// GetChart 订单分类抬头
func GetChart(ctx *gin.Context) (map[string]interface{}, error) {
	charts := map[string]interface{}{
		"all":        0,
		"complete":   0,
		"del":        0,
		"refund":     0,
		"statusAll":  0,
		"unevaluate": 0,
		"unpaid":     0,
		"unshipped":  0,
		"untake":     0,
	}
	for name, _ := range charts {
		if cc, err := getOrderCount(name, ctx); err != nil {
			return nil, err
		} else {
			charts[name] = cc
		}

	}

	return charts, nil
}

//1: 未支付 2: 未发货 3: 待收货 4: 待评价 5: 交易完成 6: 已退款 7: 已删除
// getOrderConditions
func getOrderConditions() []response.OrderCondition {
	conditions := []response.OrderCondition{
		{Name: "all", Type: "0", Conditions: nil},
		{Name: "unpaid", Type: "1", Conditions: map[string]interface{}{"paid": g.StatusFalse, "is_cancel": g.StatusFalse}},
		{Name: "unshipped", Type: "2", Conditions: map[string]interface{}{"paid": g.StatusTrue, "status": model.OrderStatusNoDeliver, "is_cancel": g.StatusTrue}},
		{Name: "untake", Type: "3", Conditions: map[string]interface{}{"status": model.OrderStatusNoReceive, "is_cancel": g.StatusFalse}},
		{Name: "unevaluate", Type: "4", Conditions: map[string]interface{}{"status": model.OrderStatusNoComment, "is_cancel": g.StatusFalse}},
		{Name: "complete", Type: "5", Conditions: map[string]interface{}{"status": model.OrderStatusFinish, "is_cancel": g.StatusFalse}},
		{Name: "refund", Type: "6", Conditions: map[string]interface{}{"status": model.OrderStatusRefund, "is_cancel": g.StatusFalse}},
		{Name: "del", Type: "7", Conditions: map[string]interface{}{"is_cancel": g.StatusTrue}},
	}
	return conditions
}

// getOrderConditionByStatus 获取查询条件
func getOrderConditionByStatus(status string) response.OrderCondition {
	conditions := getOrderConditions()
	for _, condition := range conditions {
		if condition.Type == status {
			return condition
		}
	}
	return conditions[0]
}

// GetOrderByOrderId 获取订单
func GetOrderByOrderId(req request.GetById) (model.Order, error) {
	var order model.Order
	db := g.TENANCY_DB.Model(&model.Order{}).
		Where("id = ?", req.Id)
	db = CheckTenancyIdAndUserId(db, req, "")
	err := db.First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

// GetOrderDetailById 订单详情，包括关联数据
func GetOrderDetailById(req request.GetById) (response.OrderDetail, error) {
	var order response.OrderDetail
	db := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,general_infos.nick_name as user_nick_name").
		Joins("left join sys_users on orders.sys_user_id = sys_users.id").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id").
		Joins(fmt.Sprintf("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id and sys_authorities.authority_type = %d", multi.GeneralAuthority))

	db = CheckTenancyIdAndUserId(db, req, "orders.")

	err := db.Where("orders.id = ?", req.Id).First(&order).Error
	if err != nil {
		return order, err
	}

	orderProducts, err := GetOrderProductsByOrderIds([]uint{order.ID})
	if err != nil {
		return order, err
	}

	order.OrderProduct = []response.OrderProduct{}
	for _, orderProduct := range orderProducts {
		if order.ID == orderProduct.OrderID {
			order.OrderProduct = append(order.OrderProduct, orderProduct)
		}
	}

	return order, nil
}

// GetOrderRecord 订单操作记录
func GetOrderRecord(id uint, info request.PageInfo) ([]model.OrderStatus, int64, error) {
	var orderRecord []model.OrderStatus
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.OrderStatus{}).Where("order_id = ?", id)
	err := db.Count(&total).Error
	if err != nil {
		return orderRecord, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&orderRecord).Error
	if err != nil {
		return orderRecord, total, err
	}
	return orderRecord, total, nil
}

// RemarkOrder  备注订单
func RemarkOrder(id uint, remark map[string]interface{}, ctx *gin.Context) error {
	return UpdateOrderById(g.TENANCY_DB, id, remark)
}

// DeliveryOrder 发货
func DeliveryOrder(id uint, delivery request.DeliveryOrder, ctx *gin.Context) error {
	var changeMessage string
	var deliveryName string
	switch delivery.DeliveryType {
	case model.DeliverTypeFH:
		express, err := GetExpressByCode(delivery.DeliveryName)
		if err != nil {
			return fmt.Errorf("get express err %w", err)
		}
		deliveryName = express.Name
		changeMessage = fmt.Sprintf("订单已配送【快递名称】:%s; 【快递单号】：%s", deliveryName, delivery.DeliveryId)
	case model.DeliverTypeSH:
		deliveryName = delivery.DeliveryName
		regexp := regexp.MustCompile(`^1[3456789]{1}\d{9}$`)
		if !regexp.MatchString(delivery.DeliveryId) {
			return fmt.Errorf("手机号格式错误")
		}
		changeMessage = fmt.Sprintf("订单已配送【送货人姓名】:%s; 【手机号】：%s", deliveryName, delivery.DeliveryId)
	case model.DeliverTypeXN:
		changeMessage = "订单已配送【虚拟发货】"
	default:
		return fmt.Errorf("error deliver type %d", delivery.DeliveryType)
	}

	orderDelivery := map[string]interface{}{
		"delivery_id":   delivery.DeliveryId,
		"delivery_name": deliveryName,
		"delivery_type": delivery.DeliveryType,
		"status":        model.OrderStatusNoReceive,
	}
	err := UpdateOrderById(g.TENANCY_DB, id, orderDelivery)
	if err != nil {
		return fmt.Errorf("update order info %w", err)
	}

	orderStatus := model.OrderStatus{
		ChangeType:    fmt.Sprintf("delivery_%d", delivery.DeliveryType),
		ChangeMessage: changeMessage,
		ChangeTime:    time.Now(),
		OrderID:       id,
	}
	err = g.TENANCY_DB.Model(&model.OrderStatus{}).Create(&orderStatus).Error
	if err != nil {
		return fmt.Errorf("update order status info %d", err)
	}
	return nil
}

// GetOrderInfoList 订单列表
func GetOrderInfoList(info request.OrderPageInfo, ctx *gin.Context) (gin.H, error) {

	stat := []map[string]interface{}{
		{"className": "el-icon-s-goods", "count": 0, "field": "件", "name": "已支付订单数量"},
		{"className": "el-icon-s-order", "count": 0, "field": "元", "name": "实际支付金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "已退款金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "微信支付金额"},
		{"className": "el-icon-s-finance", "count": 0, "field": "元", "name": "余额支付金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "支付宝支付金额"},
	}
	orderList := []response.OrderList{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,sys_tenancies.name as tenancy_name,sys_tenancies.is_trader as is_trader,group_orders.group_order_sn as group_order_sn").
		Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")

	db, err := getOrderSearch(info, ctx, db)
	if err != nil {
		return nil, err
	}

	stat, err = getStat(info, ctx, stat)
	if err != nil {
		return nil, err
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = db.Limit(limit).Offset(offset).Find(&orderList).Error
	if err != nil {
		return nil, err
	}

	if len(orderList) > 0 {
		orderIds := []uint{}
		for _, order := range orderList {
			orderIds = append(orderIds, order.ID)
		}

		orderProducts, err := GetOrderProductsByOrderIds(orderIds)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(orderList); i++ {
			orderList[i].OrderProduct = []response.OrderProduct{}
			for _, orderProduct := range orderProducts {
				if orderList[i].ID == orderProduct.OrderID {
					orderList[i].OrderProduct = append(orderList[i].OrderProduct, orderProduct)
				}
			}
		}
	}
	ginH := gin.H{
		"stat":     stat,
		"list":     orderList,
		"total":    total,
		"page":     info.Page,
		"pageSize": info.PageSize,
	}
	return ginH, nil
}

// GetOrdersProductByProductIds 订单产品
func GetOrdersProductByProductIds(orderProductIds []uint, req request.GetById) ([]response.OrderProduct, error) {
	orderProducts := []response.OrderProduct{}
	db := g.TENANCY_DB.Model(&model.OrderProduct{}).
		Where("order_id = ?", req.Id).
		Where("refund_num > 0").
		Where("id in ?", orderProductIds)
	err := db.Find(&orderProducts).Error
	if err != nil {
		return orderProducts, err
	}
	return orderProducts, nil
}

// GetTotalRefundPrice 退款总价
func GetTotalRefundPrice(products []response.OrderProduct) float64 {
	totalRefundPrice := decimal.NewFromInt(0)
	for _, product := range products {
		productPrice := decimal.NewFromFloat(product.ProductPrice).Mul(decimal.NewFromInt(product.ProductNum))
		totalRefundPrice = totalRefundPrice.Add(productPrice)
	}

	price, _ := totalRefundPrice.Round(2).Float64()
	return price
}

// GetOrderProductsByOrderIds 获取多个订单的产品
func GetOrderProductsByOrderIds(orderIds []uint) ([]response.OrderProduct, error) {
	orderProducts := []response.OrderProduct{}
	err := g.TENANCY_DB.Model(&model.OrderProduct{}).Where("order_id in ?", orderIds).Find(&orderProducts).Error
	if err != nil {
		return orderProducts, err
	}
	return orderProducts, nil
}

// getOrderSearch
func getOrderSearch(info request.OrderPageInfo, ctx *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	if info.Status != "" {
		cond := getOrderConditionByStatus(info.Status)
		if cond.IsDeleted {
			db = db.Unscoped()
		}

		for key, cn := range cond.Conditions {
			if cn == nil {
				db = db.Where(fmt.Sprintf("%s%s", "orders.", key))
			} else {
				db = db.Where(fmt.Sprintf("%s%s = ?", "orders.", key), cn)
			}
		}
	}

	if multi.IsTenancy(ctx) {
		db = db.Where("orders.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	} else {
		if info.SysTenancyId > 0 {
			db = db.Where("orders.sys_tenancy_id = ?", info.SysTenancyId)
		}
	}

	if info.SysUserId > 0 {
		db.Where("orders.sys_user_id = ?", info.SysUserId)
	}

	if info.PatientId > 0 {
		db.Where("orders.patient_id = ?", info.PatientId)
	}

	if info.Date != "" {
		db = filterDate(db, info.Date, "orders")
	}

	if info.OrderType != "" && info.OrderType != "0" {
		db = db.Where("orders.order_type = ?", info.OrderType)
	}

	if info.ActivityType != "" {
		db = db.Where("orders.activity_type = ?", info.ActivityType)
	}

	if info.IsTrader != "" {
		db = db.Where("sys_tenancies.is_trader = ?", info.IsTrader)
	}
	if info.OrderSn != "" {
		db = db.Where("orders.order_sn like ?", info.OrderSn+"%")
	}

	if info.Keywords != "" {
		db = db.Where(g.TENANCY_DB.Where("orders.order_sn like ?", info.Keywords+"%").Or("orders.real_name like ?", info.Keywords+"%").Or("orders.user_phone like ?", info.Keywords+"%"))
	}

	if info.Username != "" {
		db = db.Where(g.TENANCY_DB.Where("orders.order_sn like ?", info.Keywords+"%").Or("orders.real_name like ?", info.Keywords+"%").Or("orders.user_phone like ?", info.Keywords+"%"))
	}
	return db, nil
}

func getStat(info request.OrderPageInfo, ctx *gin.Context, stat []map[string]interface{}) ([]map[string]interface{}, error) {
	// 已支付订单数量
	var all int64
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Where("orders.paid =?", 1).Count(&all).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[0]["count"] = all
	}

	//实际支付金额
	var payPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {

		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).First(&payPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[1]["count"] = payPrice.Count
	}

	//已退款金额
	var orderIds []uint
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {

		err = db.Select("orders.id").Find(&orderIds).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if len(orderIds) > 0 {
			stat[2]["count"], _ = GetRefundOrder(orderIds, ctx)
		}
	}

	//微信支付金额
	var wxPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type in ?", []int{model.PayTypeWx, model.PayTypeRoutine, model.PayTypeH5}).First(&wxPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[3]["count"] = wxPayPrice.Count
	}

	//余额支付金额
	var blanPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type = ?", model.PayTypeBalance).First(&blanPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[4]["count"] = blanPayPrice.Count
	}

	//支付宝支付金额
	var aliPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type = ?", model.PayTypeAlipay).First(&aliPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[5]["count"] = aliPayPrice.Count
	}

	return stat, nil
}

// GetOrdersByGroupOrderId 根据订单组获取订单
func GetOrdersByGroupOrderId(groupOrderId uint) ([]model.Order, error) {
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).Where("group_order_id = ?", groupOrderId).Find(&orders).Error
	if err != nil {
		return orders, fmt.Errorf("获取订单错误 %w	", err)
	}
	return orders, nil
}

// ChangeOrder 修改订单备注和
func ChangeOrder(id uint, order request.OrderRemarkAndUpdate, ctx *gin.Context) error {
	data := map[string]interface{}{
		"pay_price":     order.PayPrice,
		"total_price":   order.TotalPrice,
		"total_postage": order.TotalPostage,
	}
	return UpdateOrderById(g.TENANCY_DB, id, data)
}

// UpdateOrderById 更新订单
func UpdateOrderById(db *gorm.DB, id uint, data map[string]interface{}) error {
	err := db.Model(&model.Order{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return fmt.Errorf("更新订单错误 %w", err)
	}
	return nil
}

// UpdateOrderByIds 批量更新订单
func UpdateOrderByIds(db *gorm.DB, ids []uint, data map[string]interface{}) error {
	err := db.Model(&model.Order{}).Where("id in ?", ids).Updates(data).Error
	if err != nil {
		return fmt.Errorf("批量更新订单 %w", err)
	}
	return nil
}

// GetNoPayOrdersByOrderSn 根据订单号获取未支付订单
func GetNoPayOrdersByOrderSn(orderSn string) ([]model.Order, error) {
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).Where("order_sn = ?", orderSn).
		Where("is_system_del = ?", g.StatusFalse).
		Where("is_cancel = ?", g.StatusFalse).
		Where("status = ?", model.OrderStatusNoPay).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// CheckOrder 结算订单
func CheckOrder(cartIds []uint, tenancyId, userId, patientId uint) (response.CheckOrder, error) {
	return GetOrderInfoByCartId(tenancyId, userId, patientId, cartIds)
}

// GetOrderInfoByCartId 根据购物车获取订单信息
func GetOrderInfoByCartId(tenancyId, userId, patientId uint, cartIds []uint) (response.CheckOrder, error) {
	var res response.CheckOrder
	list, fails, _, err := GetCartList(tenancyId, userId, patientId, cartIds)
	if err != nil {
		return res, fmt.Errorf("获取购物车信息 %w", err)
	}

	if len(fails) > 0 {
		return res, fmt.Errorf("购物车商品已失效")
	}

	if len(list) > 1 || len(list) == 0 {
		return res, fmt.Errorf("购物车数据异常:%d", len(list))
	}

	res.OrderType = model.OrderTypeSelf // TODO:所有订单默认自提
	res.PostagePrice = decimal.NewFromInt(0)
	res.DownPrice = decimal.NewFromInt(0)
	res.CartList = list[0]
	res.ProductPrices = map[uint]map[string]decimal.Decimal{}
	if len(res.CartList.Products) > 0 {
		for _, product := range res.CartList.Products {
			productPrice := decimal.NewFromFloat(product.AttrValue.Price).Mul(decimal.NewFromInt(product.CartNum))
			productOtPrice := decimal.NewFromFloat(product.AttrValue.OtPrice).Mul(decimal.NewFromInt(product.CartNum))
			res.ProductPrices[product.ProductID] = map[string]decimal.Decimal{
				"price":   productPrice,
				"otPrice": productOtPrice,
			}
			res.TotalPrice = res.TotalPrice.Add(productPrice)
			res.TotalOtPrice = res.TotalOtPrice.Add(productOtPrice)
			res.TotalNum += product.CartNum
		}
	}
	res.OrderPrice = res.TotalPrice.Add(res.PostagePrice).Sub(res.DownPrice)
	res.OrderOtPrice = res.TotalOtPrice.Add(res.PostagePrice).Sub(res.DownPrice)
	return res, nil
}

// CreateOrder 新建订单
// - 生成订单组
// - 生成订单
// - 二维码需要 data:image/png;base64,
func CreateOrder(req request.CreateOrder, tenancyId, userId, patientId uint, tenancyName string) ([]byte, uint, error) {
	var png []byte

	var order model.Order
	// 床旁用户登录，userId 为患者id
	patient, err := GetPatientById(patientId, tenancyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, fmt.Errorf("患者不存在")
		} else {
			return nil, 0, fmt.Errorf("患者数据错误")
		}
	}
	userAddress := fmt.Sprintf("%s-%s-%s床", tenancyName, patient.LocName, patient.BedNum)
	orderInfo, err := GetOrderInfoByCartId(tenancyId, userId, patientId, req.CartIds)
	if err != nil {
		return nil, 0, err
	}

	// 获取成本价
	var cost decimal.Decimal
	for _, product := range orderInfo.Products {
		costPrice := decimal.NewFromFloat(product.AttrValue.Cost).Mul(decimal.NewFromInt(product.CartNum))
		cost = cost.Add(costPrice)
	}

	totalPrice, _ := orderInfo.TotalPrice.Round(2).Float64()
	postagePrice, _ := orderInfo.PostagePrice.Round(2).Float64()
	orderPrice, _ := orderInfo.OrderPrice.Round(2).Float64()
	orderCost, _ := cost.Round(2).Float64()

	groupOrder := model.GroupOrder{
		GroupOrderSn: g.CreateOrderSn("G"),
		RealName:     patient.Name,
		UserPhone:    patient.Phone,
		UserAddress:  userAddress,
		TotalNum:     orderInfo.TotalNum,
		TotalPrice:   totalPrice,
		PayPrice:     orderPrice,
		TotalPostage: postagePrice,
		PayPostage:   postagePrice,
		Paid:         g.StatusFalse,
		PatientID:    patientId,
		SysUserID:    userId,
		Cost:         orderCost,
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 订单组
		err := CreateGroupOrder(tx, groupOrder)
		if err != nil {
			return err
		}
		// 订单
		order = model.Order{
			BaseOrder: model.BaseOrder{
				OrderSn:      g.CreateOrderSn(req.OrderType),
				RealName:     patient.Name,
				UserPhone:    patient.Phone,
				UserAddress:  userAddress,
				OrderType:    req.OrderType,
				Remark:       req.Remark,
				TotalNum:     orderInfo.TotalNum,
				TotalPrice:   totalPrice,
				PayPrice:     orderPrice,
				TotalPostage: postagePrice,
				PayPostage:   postagePrice,
				Paid:         g.StatusFalse,
				Cost:         orderCost,
			},
			PatientID:    patientId,
			SysUserID:    userId,
			SysTenancyID: tenancyId,
			GroupOrderID: groupOrder.ID,
		}

		err = tx.Model(&model.Order{}).Create(&order).Error
		if err != nil {
			return err
		}

		// 订单状态
		orderStatus := model.OrderStatus{ChangeType: "create", ChangeMessage: "生成订单", ChangeTime: time.Now(), OrderID: order.ID}
		err = CreateOrderStatus(tx, orderStatus)
		if err != nil {
			return err
		}

		// 生成订单商品
		orderProducts := []model.OrderProduct{}
		for _, cartProduct := range orderInfo.Products {
			cartInfo := request.CartInfo{
				Product: request.CartInfoProduct{
					Image:     cartProduct.AttrValue.Image,
					StoreName: cartProduct.StoreName,
				},
				ProductAttr: request.CartInfoProductAttr{
					Price: cartProduct.AttrValue.Price,
					Sku:   cartProduct.AttrValue.Sku,
				},
			}
			ci, _ := json.Marshal(&cartInfo)
			orderProduct := model.OrderProduct{
				OrderID:   order.ID,
				SysUserID: userId,
				CartID:    cartProduct.Id,
				ProductID: cartProduct.ProductID,
				CartInfo:  string(ci),
				BaseOrderProduct: model.BaseOrderProduct{
					ProductSku:   cartProduct.AttrValue.Sku,
					Unique:       cartProduct.AttrValue.Unique,
					IsRefund:     g.StatusFalse,
					ProductNum:   cartProduct.CartNum,
					ProductType:  model.GeneralSale,
					RefundNum:    cartProduct.CartNum,
					IsReply:      g.StatusFalse,
					ProductPrice: cartProduct.AttrValue.Price,
				},
			}
			orderProducts = append(orderProducts, orderProduct)
		}
		err = tx.Create(&orderProducts).Error
		if err != nil {
			return err
		}

		// 减库存
		for _, cartProduct := range orderInfo.Products {
			if err = DecStock(tx, cartProduct.ProductID, cartProduct.CartNum); err != nil {
				return err
			}
			if err = DecSkuStock(tx, cartProduct.ProductID, cartProduct.AttrValue.Unique, cartProduct.CartNum); err != nil {
				return err
			}
		}

		err = ChangeIsPayByIds(tx, req.CartIds)
		if err != nil {
			return fmt.Errorf("生成订单-修改购物车属性错误 %w", err)
		}
		// 生成二维码
		png, err = GetQrCode(order.ID, tenancyId, userId, order.OrderType)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return png, order.ID, nil
}

// CheckOrderStatusBeforePay 订单状态检测
func CheckOrderStatusBeforePay(req request.GetById) (model.Order, error) {
	order, err := GetOrderByOrderId(req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, fmt.Errorf("订单不存在或者已被删除")
	}
	if err != nil {
		return order, err
	}
	if order.Status != model.OrderStatusNoPay {
		return order, fmt.Errorf("订单已付款,请勿重复操作")
	}
	return order, nil
}

// GetQrCode 生成支付二维码
func GetQrCode(orderId, tenancyId, userId uint, orderType int) ([]byte, error) {
	seitURL, err := param.GetSeitURL()
	if err != nil {
		return nil, err
	}
	// 生成支付地址二维码
	payUrl := fmt.Sprintf("%s/v1/pay/payOrder?orderId=%d&tenancyId=%d&userId=%d&orderType=%d", seitURL, orderId, tenancyId, userId, orderType)
	q, err := qrcode.New(payUrl, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	png, err := q.PNG(256)
	if err != nil {
		return nil, err
	}
	return png, nil
}

// GetThisMonthOrdersByUserId 获取用户当月订单
func GetThisMonthOrdersByUserId(userId uint) ([]model.Order, error) {
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Where("sys_user_id = ?", userId).
		Where("DATE_FORMAT(`created_at`,'%Y%m')=DATE_FORMAT(CURDATE(),'%Y%m')").
		Where("paid = ?", g.StatusTrue).
		Not("status = ?", model.OrderStatusRefund).
		Where("is_cancel = ?", g.StatusFalse).
		Where("is_system_del = ?", g.StatusFalse).
		Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

// GetThisMonthOrderPriceByUserId 获取用户当月支付金额
func GetThisMonthOrderPriceByUserId(userId uint) (response.GeneralUserDetail, error) {
	var user response.GeneralUserDetail
	err := g.TENANCY_DB.Model(&model.Order{}).
		Select("sum(orders.pay_price) as total_pay_price, count(orders.id) as total_pay_count").
		Where("sys_user_id = ?", userId).
		Where("DATE_FORMAT(`created_at`,'%Y%m')=DATE_FORMAT(CURDATE(),'%Y%m')").
		Where("paid = ?", g.StatusTrue).
		Not("status = ?", model.OrderStatusRefund).
		Where("is_cancel = ?", g.StatusFalse).
		Where("is_system_del = ?", g.StatusFalse).
		Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetNoPayOrders 获取未支付订单
func GetNoPayOrders() ([]model.Order, error) {
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Where("paid = ?", g.StatusFalse).
		Where("status = ?", model.OrderStatusNoPay).
		Where("is_cancel = ?", g.StatusFalse).
		Where("is_system_del = ?", g.StatusFalse).
		Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

// CreateOrderStatus  新建订单状态
func CreateOrderStatus(db *gorm.DB, orderStatus model.OrderStatus) error {
	err := db.Create(&orderStatus).Error
	if err != nil {
		return err
	}
	return nil
}

// ChangeOrderStatusByOrderId 修改订单状态
func ChangeOrderStatusByOrderId(orderId uint, changeData map[string]interface{}, changeType, changeMessage string) error {
	orderProducts, err := GetOrderProductsByOrderIds([]uint{orderId})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := UpdateOrderById(tx, orderId, changeData)
		if err != nil {
			return err
		}
		orderStatus := model.OrderStatus{ChangeType: changeType, ChangeMessage: changeMessage, ChangeTime: time.Now(), OrderID: orderId}
		err = CreateOrderStatus(tx, orderStatus)
		if err != nil {
			return err
		}

		if changeType == "cancel" {
			// 退回库存
			for _, cartProduct := range orderProducts {
				if err = IncStock(tx, cartProduct.ProductID, cartProduct.ProductNum); err != nil {
					return err
				}
				if err = IncSkuStock(tx, cartProduct.ProductID, cartProduct.Unique, cartProduct.ProductNum); err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// CancelOrder 用户取消订单
func CancelOrder(req request.GetById) error {
	_, err := CheckOrderStatusBeforePay(req)
	if err != nil {
		return err
	}

	changeData := map[string]interface{}{"is_cancel": g.StatusTrue}
	err = ChangeOrderStatusByOrderId(req.Id, changeData, "cancel", "取消订单")
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrder 商户取消订单
func DeleteOrder(req request.GetById) error {
	_, err := CheckOrderStatusBeforePay(req)
	if err != nil {
		return err
	}
	changeData := map[string]interface{}{"is_system_del": g.StatusTrue}
	err = UpdateOrderById(g.TENANCY_DB, req.Id, changeData)
	if err != nil {
		return err
	}
	return nil
}

// ChangeOrderPayNotifyByOrderSn 修改支付状态
func ChangeOrderPayNotifyByOrderSn(changeData map[string]interface{}, orderSn, changeType, changeMessage string) (model.Payload, error) {
	var palyload model.Payload
	orders, err := GetNoPayOrdersByOrderSn(orderSn)
	if err != nil {
		return palyload, err
	}
	if len(orders) != 1 {
		return palyload, fmt.Errorf("%s 订单号重复生产 %d 个订单", orderSn, len(orders))
	}
	err = ChangeOrderStatusByOrderId(orders[0].ID, changeData, changeType, changeMessage)
	if err != nil {
		return palyload, err
	}

	palyload = model.Payload{
		OrderId:   orders[0].ID,
		TenancyId: orders[0].SysTenancyID,
		UserId:    orders[0].SysUserID,
		OrderType: orders[0].OrderType,
		PayType:   orders[0].PayType,
		CreatedAt: time.Now(),
	}
	return palyload, nil
}

func UpdateOrderProduct(db *gorm.DB, orderProductId uint, data map[string]interface{}) error {
	err := db.Model(&model.OrderProduct{}).Where("id = ?", orderProductId).Updates(data).Error
	if err != nil {
		return fmt.Errorf("update order product %w", err)
	}
	return nil
}

func GetOrderAutoAgree() ([]uint, error) {
	whereCreatedAt := fmt.Sprintf("now() > SUBDATE(pay_time,interval -%s DAY)", param.GetOrderAutoTakeOrderTime())
	orderIds := []uint{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Where("paid = ?", g.StatusTrue).
		Where("status = ?", model.OrderStatusNoReceive).
		Where(whereCreatedAt).
		Find(&orderIds).Error
	if err != nil {
		return orderIds, err
	}
	return orderIds, nil
}

// AutoTakeOrders 自动收货
func AutoTakeOrders(orderIds []uint) error {
	var orderStatues []model.OrderStatus
	for _, orderId := range orderIds {
		orderStatus := model.OrderStatus{ChangeType: "take", ChangeMessage: "已收货", ChangeTime: time.Now(), OrderID: orderId}
		orderStatues = append(orderStatues, orderStatus)
	}
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := UpdateOrderByIds(tx, orderIds, map[string]interface{}{"status": model.OrderStatusNoComment})
		if err != nil {
			return err
		}

		err = tx.Model(&model.OrderStatus{}).Create(&orderStatues).Error
		if err != nil {
			return fmt.Errorf("生成订单操作记录 %w", err)
		}
		return nil
	})
}
