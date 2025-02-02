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
	"github.com/snowlyg/go-tenancy/service/scope"
	"github.com/snowlyg/go-tenancy/utils/param"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
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
func getOrderCount(info request.OrderPageInfo, name string, ctx *gin.Context) (int64, error) {
	var count int64
	wheres := getOrderConditions()
	for _, where := range wheres {
		if where.Name == name {
			db := g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id")
			if where.Conditions != nil && len(where.Conditions) > 0 {
				for key, cn := range where.Conditions {
					if cn == nil {
						db = db.Where("orders." + key)
					} else {
						db = db.Where(fmt.Sprintf("%s = ?", "orders."+key), cn)
					}
				}
			}
			if multi.IsTenancy(ctx) {
				db = CheckTenancyId(db, multi.GetTenancyId(ctx), "orders.")
			}
			db, err := getOrderSearch(info, ctx, db)
			if err != nil {
				return count, err
			}

			err = db.Count(&count).Error
			if err != nil {
				return count, err
			}
		}
	}

	return count, nil
}

// GetFilter 订单过滤数量
func GetFilter(info request.OrderPageInfo, ctx *gin.Context) ([]map[string]interface{}, error) {
	charts := []map[string]interface{}{
		{"count": 0, "orderType": "", "title": "全部"},
		{"count": 0, "orderType": "1", "title": "普通订单"},
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
		db, err := getOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}

		err = db.Where("orders.order_type = ?", chart["orderType"]).Count(&count).Error
		if err != nil {
			return nil, err
		} else {
			chart["count"] = count
		}
	}

	return charts, nil
}

// GetChart 订单分类抬头
func GetChart(info request.OrderPageInfo, ctx *gin.Context) (map[string]interface{}, error) {
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
		if cc, err := getOrderCount(info, name, ctx); err != nil {
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

// GetOrderById 获取订单
func GetOrderById(id uint, funcs ...func(*gorm.DB) *gorm.DB) (model.Order, error) {
	var order model.Order
	db := g.TENANCY_DB.Model(&model.Order{}).Where("id = ?", id)
	// db = CheckTenancyIdAndUserId(db, req, "")
	db = db.Scopes(funcs...)
	err := db.First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

// GetOrderDetailById 订单详情，包括关联数据
func GetOrderDetailById(id uint, funcs ...func(*gorm.DB) *gorm.DB) (response.OrderDetail, error) {
	var order response.OrderDetail
	db := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,c_users.nick_name as user_nick_name").
		Joins("left join c_users on orders.c_user_id = c_users.id").
		Joins(fmt.Sprintf("left join sys_authorities on sys_authorities.authority_id = c_users.authority_id and sys_authorities.authority_type = %d", multi.GeneralAuthority))
	db = db.Scopes(funcs...)

	err := db.Where("orders.id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}

	if order.CUserId > 0 {
		cuser, err := GetGeneralDetail(order.CUserId)
		if err != nil {
			return order, err
		}
		order.UserNickName = cuser.NickName
	}

	orderProducts, err := GetOrderProductsByOrderIds([]uint{order.ID})
	if err != nil {
		return order, err
	}

	order.OrderProduct = []response.OrderProduct{}
	for _, orderProduct := range orderProducts {
		if order.ID == orderProduct.OrderId {
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
	db = OrderBy(db, info.OrderBy, info.SortBy)
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
		ChangeType:    model.OrderChangeType(fmt.Sprintf("delivery_%d", delivery.DeliveryType)),
		ChangeMessage: changeMessage,
		ChangeTime:    time.Now(),
		OrderId:       id,
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

	db = OrderBy(db, info.OrderBy, info.SortBy, "orders.")
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
				if orderList[i].ID == orderProduct.OrderId {
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
func GetOrdersProductByProductIds(orderProductIds []uint, orderId uint) ([]response.OrderProduct, error) {
	orderProducts := []response.OrderProduct{}
	db := g.TENANCY_DB.Model(&model.OrderProduct{}).
		Where("order_id = ?", orderId).
		Where("refund_num > 0").
		Where("is_refund < 3"). // 未退款
		Where("id in ?", orderProductIds)
	err := db.Find(&orderProducts).Error
	if err != nil {
		g.TENANCY_LOG.Error("获取订单产品错误", zap.String("GetOrdersProductByProductIds()", err.Error()))
		return orderProducts, fmt.Errorf("获取订单产品错误 %w", err)
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

	price, b := totalRefundPrice.Round(2).Float64()
	if !b {
		g.TENANCY_LOG.Error("退款总价计算错误", zap.String("GetTotalRefundPrice()", "totalRefundPrice.Round(2).Float64()"))
	}
	return price
}

// GetOrderProductsByOrderIds 获取多个订单的产品
func GetOrderProductsByOrderIds(orderIds []uint) ([]response.OrderProduct, error) {
	orderProducts := []response.OrderProduct{}
	err := g.TENANCY_DB.Model(&model.OrderProduct{}).Where("order_id in ?", orderIds).Find(&orderProducts).Error
	if err != nil {
		g.TENANCY_LOG.Error("获取订单商品错误", zap.String("GetOrderProductsByOrderIds()", err.Error()))
		return orderProducts, fmt.Errorf("获取订单商品错误 %w", err)
	}
	return orderProducts, nil
}

// getOrderSearch
func getOrderSearch(info request.OrderPageInfo, ctx *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	if multi.IsTenancy(ctx) {
		db = db.Where("orders.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	} else {
		if info.SysTenancyId > 0 {
			db = db.Where("orders.sys_tenancy_id = ?", info.SysTenancyId)
		}
	}

	if info.CUserId > 0 {
		db.Where("orders.c_user_id = ?", info.CUserId)
	}

	if info.Date != "" {
		db = db.Scopes(scope.FilterDate(info.Date, "created_at", "orders"))
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
	db := g.TENANCY_DB.Model(&model.Order{}).
		Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")
	if db, err := getOrderSearch(info, ctx, db); err != nil {
		return nil, err
	} else {
		err = db.Where("orders.paid =?", g.StatusTrue).Count(&all).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[0]["count"] = all
	}

	//实际支付金额
	var payPrice request.Result
	if db, err := getOrderSearch(info, ctx, db); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", g.StatusTrue).First(&payPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[1]["count"] = payPrice.Count
	}

	//已退款金额
	var orderIds []uint
	if db, err := getOrderSearch(info, ctx, db); err != nil {
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
	if db, err := getOrderSearch(info, ctx, db); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", g.StatusTrue).Where("orders.pay_type in ?", []int{model.PayTypeWx, model.PayTypeRoutine, model.PayTypeH5}).First(&wxPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[3]["count"] = wxPayPrice.Count
	}

	//余额支付金额
	var blanPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, db); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", g.StatusTrue).Where("orders.pay_type = ?", model.PayTypeBalance).First(&blanPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[4]["count"] = blanPayPrice.Count
	}

	//支付宝支付金额
	var aliPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, db); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", g.StatusTrue).Where("orders.pay_type = ?", model.PayTypeAlipay).First(&aliPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[5]["count"] = aliPayPrice.Count
	}

	return stat, nil
}

// GetOrdersByGroupOrderId 根据订单组获取未取消订单订单
func GetOrdersByGroupOrderId(groupOrderId uint) ([]model.Order, error) {
	whereCreatedAt := fmt.Sprintf("now() > SUBDATE(created_at,interval -%d minute)", param.GetOrderAutoCloseTime())
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Where("group_order_id = ? and is_cancel = ? and paid =?", groupOrderId, g.StatusFalse, g.StatusFalse).
		Where(whereCreatedAt).
		Find(&orders).Error
	if err != nil {
		return orders, fmt.Errorf("获取订单错误 %w", err)
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
		g.TENANCY_LOG.Error("更新订单错误", zap.String("UpdateOrderById()", err.Error()))
		return fmt.Errorf("更新订单错误 %w", err)
	}
	return nil
}

// UpdateOrderByIds 批量更新订单
func UpdateOrderByIds(db *gorm.DB, ids []uint, data map[string]interface{}) error {
	err := db.Model(&model.Order{}).Where("id in ?", ids).Updates(data).Error
	if err != nil {
		g.TENANCY_LOG.Error("批量更新订单错误", zap.String("UpdateOrderByIds()", err.Error()))
		return fmt.Errorf("批量更新订单 %w", err)
	}
	return nil
}

// GetNoPayOrdersByOrderSn 根据订单号获取未支付订单
func GetNoPayOrdersByOrderSn(orderSn string) ([]response.OrderDetail, error) {
	orders := []response.OrderDetail{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,c_users.nick_name as user_nick_name").
		Joins("left join c_users on orders.c_user_id = c_users.id").
		Where("orders.order_sn = ?", orderSn).
		Where("orders.is_system_del = ?", g.StatusFalse).
		Where("orders.is_cancel = ?", g.StatusFalse).
		Where("orders.status = ?", model.OrderStatusNoPay).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// CheckOrder 结算订单
func CheckOrder(cartIds []uint, tenancyId, userId uint) (response.CheckOrder, error) {
	return GetOrderInfoByCartId(tenancyId, userId, cartIds)
}

// GetOrderInfoByCartId 根据购物车获取订单信息
func GetOrderInfoByCartId(tenancyId, userId uint, cartIds []uint) (response.CheckOrder, error) {
	var res response.CheckOrder
	list, fails, _, err := GetCartList(tenancyId, userId, cartIds)
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
			res.ProductPrices[product.ProductId] = map[string]decimal.Decimal{
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
func CreateOrder(req request.CreateOrder, tenancyId, userId uint, tenancyName string) ([]byte, uint, error) {
	var png []byte

	var order model.Order

	// userAddress := fmt.Sprintf("%s-%s-%s床", tenancyName, patient.LocName, patient.BedNum)
	orderInfo, err := GetOrderInfoByCartId(tenancyId, userId, req.CartIds)
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
		// RealName:     patient.Name,
		// UserPhone:    patient.Phone,
		// UserAddress:  userAddress,
		TotalNum:     orderInfo.TotalNum,
		TotalPrice:   totalPrice,
		PayPrice:     orderPrice,
		TotalPostage: postagePrice,
		PayPostage:   postagePrice,
		Paid:         g.StatusFalse,
		CUserId:      userId,
		Cost:         orderCost,
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 订单组
		err := CreateGroupOrder(tx, &groupOrder)
		if err != nil {
			return err
		}
		// 订单
		order = model.Order{
			BaseOrder: model.BaseOrder{
				OrderSn: g.CreateOrderSn(req.OrderType),
				// RealName:     patient.Name,
				// UserPhone:    patient.Phone,
				// UserAddress:  userAddress,
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

			CUserId:      userId,
			SysTenancyId: tenancyId,
			GroupOrderId: groupOrder.ID,
		}

		err = tx.Model(&model.Order{}).Create(&order).Error
		if err != nil {
			return err
		}

		// 订单状态
		orderStatus := model.OrderStatus{ChangeType: "create", ChangeMessage: "生成订单", ChangeTime: time.Now(), OrderId: order.ID}
		err = CreateOrderStatus(tx, &orderStatus)
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
				OrderId:   order.ID,
				CUserId:   userId,
				CartId:    cartProduct.Id,
				ProductId: cartProduct.ProductId,
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
			if err = DecStock(tx, cartProduct.ProductId, cartProduct.CartNum); err != nil {
				return err
			}
			if err = DecSkuStock(tx, cartProduct.ProductId, cartProduct.AttrValue.Unique, cartProduct.CartNum); err != nil {
				return err
			}
		}

		err = ChangeIsPayByIds(tx, req.CartIds)
		if err != nil {
			return fmt.Errorf("生成订单-修改购物车属性错误 %w", err)
		}
		// 生成二维码
		png, err = GetQrCode(order.ID, tenancyId, order.OrderType)
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

// GetQrCode 生成支付二维码
func GetQrCode(orderId, tenancyId uint, orderType int) ([]byte, error) {
	seitURL, err := param.GetSeitURL()
	if err != nil {
		return nil, err
	}

	// 生成支付地址二维码
	// 订单自动过期时间
	autoCloseTime := param.GetOrderAutoCloseTime()
	payUrl := fmt.Sprintf("%s/v1/pay/payOrder?orderId=%d&tenancyId=%d&orderType=%d&expire=%d", seitURL, orderId, tenancyId, orderType, time.Now().Add(time.Duration(autoCloseTime)*time.Minute).Unix())
	if g.TENANCY_CONFIG.System.Level == "debug" {
		g.TENANCY_LOG.Info("支付二维码", zap.String("url", payUrl))
	}
	q, err := qrcode.New(payUrl, qrcode.Medium)
	if err != nil {
		g.TENANCY_LOG.Error("生成二维码错误", zap.String("qrcode.New()", err.Error()))
		return nil, err
	}
	png, err := q.PNG(256)
	if err != nil {
		g.TENANCY_LOG.Error("生成二维码错误", zap.String("q.PNG(256)", err.Error()))
		return nil, err
	}
	return png, nil
}

// GetThisMonthOrdersByUserId 获取用户当月订单
func GetThisMonthOrdersByUserId(userId uint) ([]model.Order, error) {
	orders := []model.Order{}
	err := g.TENANCY_DB.Model(&model.Order{}).
		Where("c_user_id = ?", userId).
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
		Where("c_user_id = ?", userId).
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

// CreateOrderStatus  添加订单处理记录
func CreateOrderStatus(db *gorm.DB, orderStatus *model.OrderStatus) error {
	err := db.Create(orderStatus).Error
	if err != nil {
		g.TENANCY_LOG.Error("添加订单处理记录错误", zap.String("CreateOrderStatus", err.Error()))
		return fmt.Errorf("添加订单处理记录错误 %w", err)
	}
	return nil
}

// ChangeOrderStatusByOrderId 修改订单状态
func ChangeOrderStatusByOrderId(orderId uint, changeData map[string]interface{}, changeType model.OrderChangeType, changeMessage string, financials ...*model.FinancialRecord) error {
	if changeMessage == "" {
		changeMessage = changeType.ToMessage()
	}
	orderProducts, err := GetOrderProductsByOrderIds([]uint{orderId})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := UpdateOrderById(tx, orderId, changeData)
		if err != nil {
			return err
		}
		orderStatus := model.OrderStatus{ChangeType: changeType, ChangeMessage: changeMessage, ChangeTime: time.Now(), OrderId: orderId}
		err = CreateOrderStatus(tx, &orderStatus)
		if err != nil {
			return err
		}

		// 用户没有付款的商品才回退库存，已经付款的商品部回退库存
		if changeType == "cancel" {
			// 退回库存
			for _, cartProduct := range orderProducts {
				if err = IncStock(tx, cartProduct.ProductId, cartProduct.ProductNum); err != nil {
					return err
				}
				if err = IncSkuStock(tx, cartProduct.ProductId, cartProduct.Unique, cartProduct.ProductNum); err != nil {
					return err
				}
			}
		}

		// TODO:: 添加交易记录
		if len(financials) == 1 {
			if err = CreateFinancialRecord(tx, financials[0]); err != nil {
				return err
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
	changeData := map[string]interface{}{"is_cancel": g.StatusTrue}
	err := ChangeOrderStatusByOrderId(req.Id, changeData, model.ChangeTypeCancel, "")
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrder 商户取消订单
func DeleteOrder(req request.GetById) error {
	_, err := GetOrderById(req.Id)
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
func ChangeOrderPayNotifyByOrderSn(changeData map[string]interface{}, orderSn string, changeType model.OrderChangeType) (model.Payload, error) {
	var palyload model.Payload
	orders, err := GetNoPayOrdersByOrderSn(orderSn)
	if err != nil {
		return palyload, err
	}
	if len(orders) != 1 {
		return palyload, fmt.Errorf("%s 订单号重复生产 %d 个订单", orderSn, len(orders))
	}
	financialRecord := &model.FinancialRecord{
		RecordSn:      g.CreateOrderSn("RC"),
		OrderSn:       orders[0].OrderSn,
		UserInfo:      orders[0].UserNickName,
		FinancialType: "order",
		FinancialPm:   model.InFinancialPm,
		Number:        orders[0].PayPrice,
		SysTenancyId:  orders[0].SysTenancyId,
		CUserId:       orders[0].CUserId,
		OrderId:       orders[0].ID,
	}
	err = ChangeOrderStatusByOrderId(orders[0].ID, changeData, changeType, "", financialRecord)
	if err != nil {
		return palyload, err
	}

	palyload = model.Payload{
		OrderId:   orders[0].ID,
		TenancyId: orders[0].SysTenancyId,
		UserId:    orders[0].CUserId,
		OrderType: orders[0].OrderType,
		PayType:   orders[0].PayType,
		CreatedAt: time.Now(),
	}
	return palyload, nil
}

// UpdateOrderProduct 更新订单商品
func UpdateOrderProduct(db *gorm.DB, orderProductId uint, data map[string]interface{}) error {
	err := db.Model(&model.OrderProduct{}).Where("id = ?", orderProductId).Updates(data).Error
	if err != nil {
		return fmt.Errorf("update order product %w", err)
	}
	return nil
}

// GetOrderAutoAgree 获取需要自动收货订单
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
		orderStatus := model.OrderStatus{ChangeType: "take", ChangeMessage: "已收货", ChangeTime: time.Now(), OrderId: orderId}
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

// GetOrderPayPrice 获取订单支付价格
func GetOrderPayPrice(scopes ...func(*gorm.DB) *gorm.DB) (float64, error) {
	var payPrice request.Result
	db := g.TENANCY_DB.Model(&model.Order{}).Select("sum(pay_price) as count")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.First(&payPrice).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return payPrice.Count, nil
}

// GetOrderPayPriceGroup 获取订单支付价格按时间分组集合
func GetOrderPayPriceGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]request.Result, error) {
	var res []request.Result
	db := g.TENANCY_DB.Model(&model.Order{}).Select("sum(pay_price) as count , from_unixtime(unix_timestamp(pay_time),'%H:%i') as time").
		Where("paid =?", g.StatusTrue)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Group("time").Find(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, nil
}

// GetPayOrderNum 获取支付订单数量
func GetPayOrderNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := g.TENANCY_DB.Model(&model.Order{}).Where("paid =?", g.StatusTrue)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return count, nil
}

// GetOrderUserNum 获取订单用户数量
func GetOrderUserNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := g.TENANCY_DB.Model(&model.Order{})
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Where("c_user_id > ?", 0).Group("c_user_id").Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return count, nil
}

// GetOrderGroup 获取订单数量按时间分组集合
func GetOrderGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]response.ClientStaticOrder, error) {
	var res []response.ClientStaticOrder
	db := g.TENANCY_DB.Model(&model.Order{}).Select("sum(pay_price) as pay_price,count(*) as total,count(distinct c_user_id) as user,from_unixtime(unix_timestamp(pay_time),'%m-%d') as `day`").
		Where("paid =?", g.StatusTrue)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Order("day ASC").Group("day").Find(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, nil
}

// GetOrderNumGroup 获取订单数量按时间分组集合
func GetOrderNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]request.Result, error) {
	var res []request.Result
	db := g.TENANCY_DB.Model(&model.Order{}).Select("count(*) as total , from_unixtime(unix_timestamp(pay_time),'%H:%i') as time").
		Where("paid =?", g.StatusTrue)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Group("time").Find(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, nil
}

// GetOrderUserNumGroup 获取订单数量按时间分组集合
func GetOrderUserNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]request.Result, error) {
	var res []request.Result
	db := g.TENANCY_DB.Model(&model.Order{}).Select("count(DISTINCT c_user_id) as total , from_unixtime(unix_timestamp(pay_time),'%H:%i') as time").Where("paid =?", g.StatusTrue)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Where("c_user_id > ?", 0).Group("time").Find(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, nil
}

// GetPayOrderProductNum 获取支付订单商品数量
func GetPayOrderProductNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var res request.Result
	db := g.TENANCY_DB.Model(&model.OrderProduct{}).
		Select("sum(order_products.product_num) as count").
		Joins("left join orders on order_products.order_id = orders.id")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Where("orders.paid =?", g.StatusTrue).First(&res).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return int64(res.Count), nil
}

// GetPayOrderProductNumGroup 获取支付订单商品数量
// - product_id 分组
// - 根据 total desc 排序
// - limit 7 个
func GetPayOrderProductNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.MerchantStockData, error) {
	var stockData []*response.MerchantStockData
	db := g.TENANCY_DB.Model(&model.OrderProduct{}).
		Select("sum(order_products.product_num) as total,order_products.product_id,products.store_name,products.image").
		Joins("left join orders on order_products.order_id = orders.id").
		Joins("left join products on order_products.product_id = products.id")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Where("orders.paid =?", g.StatusTrue).Group("order_products.product_id").Order("total desc").Limit(7).Find(&stockData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return stockData, nil
}
