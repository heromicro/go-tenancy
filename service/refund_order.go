package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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

// GetRefundOrder 退款单金额
func GetRefundOrder(orderIds []uint, ctx *gin.Context) (float64, error) {
	var refundPayPrice request.Result
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Select("sum(refund_price) as count").
		Where("order_id in ?", orderIds).
		Where("status = ?", model.RefundStatusEnd)
	err := db.First(&refundPayPrice).Error
	if err != nil {
		return 0, err
	}
	return refundPayPrice.Count, nil
}

// getRefundOrderSearch 退款单筛选
func getRefundOrderSearch(info request.RefundOrderPageInfo, ctx *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	if multi.IsTenancy(ctx) {
		db = db.Where("refund_orders.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}

	if info.Date != "" {
		db = db.Scopes(scope.FilterDate(info.Date,"created_at", "refund_orders"))
	}

	if info.IsTrader != "" {
		db = db.Where("sys_tenancies.is_trader = ?", info.IsTrader)
	}

	// TODO:: 开启后床旁用户退款订单列表显示错误
	if info.OrderSn != "" {
		db = db.Where("orders.order_sn like ?", info.OrderSn+"%")
	}

	if info.SysUserId > 0 {
		db.Where("orders.sys_user_id = ?", info.SysUserId)
	}

	if info.PatientId > 0 {
		db.Where("orders.patient_id = ?", info.PatientId)
	}

	if info.RefundOrderSn != "" {
		db = db.Where("refund_orders.refund_order_sn like ?", info.RefundOrderSn+"%")
	}

	return db, nil
}

// GetRefundOrderInfoList 退款单列表
func GetRefundOrderInfoList(info request.RefundOrderPageInfo, ctx *gin.Context) ([]response.RefundOrderList, map[string]int64, int64, error) {
	stat := map[string]int64{
		"agree":    0,
		"all":      0,
		"audit":    0,
		"backgood": 0,
		"end":      0,
		"refuse":   0,
	}
	refundOrderList := []response.RefundOrderList{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Select("refund_orders.*,sys_tenancies.name as tenancy_name,sys_tenancies.is_trader as is_trader,c_users.nick_name as user_nick_name,orders.order_sn as order_sn,orders.activity_type as activity_type").
		Joins("left join orders on refund_orders.order_id = orders.id").
		Joins("left join sys_tenancies on refund_orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join c_users on refund_orders.sys_user_id = c_users.id")

	if info.Status != "" {
		db = db.Where("refund_orders.status = ?", info.Status)
	}

	db, err := getRefundOrderSearch(info, ctx, db)
	if err != nil {
		return refundOrderList, stat, total, err
	}

	stat, err = getRefundStat(stat, info, ctx)
	if err != nil {
		return refundOrderList, stat, total, err
	}

	err = db.Count(&total).Error
	if err != nil {
		return refundOrderList, stat, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy, "refund_orders.")
	err = db.Limit(limit).Offset(offset).Find(&refundOrderList).Error
	if err != nil {
		return refundOrderList, stat, total, err
	}

	if len(refundOrderList) > 0 {
		var refundOrderIds []uint
		for _, refundOrder := range refundOrderList {
			refundOrderIds = append(refundOrderIds, refundOrder.ID)
		}

		refundProducts, err := getRefundProducts(refundOrderIds)
		if err != nil {
			return refundOrderList, stat, total, err
		}

		for i := 0; i < len(refundOrderList); i++ {
			for _, refundProduct := range refundProducts {
				if refundOrderList[i].ID == refundProduct.RefundOrderID {
					refundOrderList[i].RefundProduct = append(refundOrderList[i].RefundProduct, refundProduct)
				}
			}
		}
	}

	return refundOrderList, stat, total, nil
}

// getRefundProducts 退款商品
func getRefundProducts(refundOrderIds []uint) ([]response.RefundProduct, error) {
	refundProducts := []response.RefundProduct{}
	err := g.TENANCY_DB.Model(&model.RefundProduct{}).
		Select("refund_products.*,order_products.*").
		Joins("left join order_products on refund_products.order_product_id = order_products.id").
		Where("refund_products.refund_order_id in ?", refundOrderIds).Find(&refundProducts).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return refundProducts, fmt.Errorf("get refund products %w", err)
	}
	return refundProducts, nil
}

// getRefundStat 退款单按状态统计数量
func getRefundStat(stat map[string]int64, info request.RefundOrderPageInfo, ctx *gin.Context) (map[string]int64, error) {
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Joins("left join orders on refund_orders.order_id = orders.id")
	{
		var all int64
		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Count(&all).Error
		if err != nil {
			return nil, err
		}
		stat["all"] = all
	}

	{
		var agree int64

		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Where("orders.status = ?", model.RefundStatusAgree).Count(&agree).Error
		if err != nil {
			return nil, err
		}
		stat["agree"] = agree
	}

	{
		var audit int64
		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Where("orders.status = ?", model.RefundStatusAudit).Count(&audit).Error
		if err != nil {
			return nil, err
		}
		stat["audit"] = audit
	}

	{
		var backgood int64
		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Where("orders.status = ?", model.RefundStatusBackgood).Count(&backgood).Error
		if err != nil {
			return nil, err
		}
		stat["backgood"] = backgood
	}

	{
		var end int64
		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Where("orders.status = ?", model.RefundStatusEnd).Count(&end).Error
		if err != nil {
			return nil, err
		}
		stat["end"] = end
	}

	{
		var refuse int64
		db, err := getRefundOrderSearch(info, ctx, db)
		if err != nil {
			return nil, err
		}
		err = db.Where("orders.status = ?", model.RefundStatusRefuse).Count(&refuse).Error
		if err != nil {
			return nil, err
		}
		stat["refuse"] = refuse
	}

	return stat, nil
}

// GetRefundOrderRecord 退款单操作记录
func GetRefundOrderRecord(id uint, info request.PageInfo) ([]model.RefundStatus, int64, error) {
	returnRecord := []model.RefundStatus{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.RefundStatus{}).Where("refund_order_id = ?", id)
	err := db.Count(&total).Error
	if err != nil {
		return returnRecord, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&returnRecord).Error
	if err != nil {
		return returnRecord, total, err
	}
	return returnRecord, total, nil
}

// GetRefundOrderRemarkMap 退款单备注表单
func GetRefundOrderRemarkMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	refundOrder, err := GetRefundOrderById(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"mer_mark","value":"%s","title":"备注","props":{"type":"text","placeholder":"请输入备注"}}],"action":"","method":"POST","title":"备注信息","config":{}}`, refundOrder.MerMark)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/refundOrder/remarkRefundOrder", id), ctx)
	return form, err
}

// GetRefundOrderMap 退款审核表单
func GetRefundOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := `{"rule":[{"type":"radio","field":"status","value":5,"title":"审核","props":{},"control":[{"value":5,"rule":[{"type":"input","field":"failMessage","value":"","title":"拒绝原因","props":{"type":"text","placeholder":"请输入拒绝原因"},"validate":[{"message":"请输入拒绝原因","required":true,"type":"string","trigger":"change"}]}]}],"options":[{"value":2,"label":"同意"},{"value":5,"label":"拒绝"}]}],"action":"","method":"POST","title":"退款审核","config":{}}`

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/refundOrder/auditRefundOrder", id), ctx)
	return form, err
}

// GetRefundOrderById 退款单详情
func GetRefundOrderById(id uint) (model.RefundOrder, error) {
	var refundOrder model.RefundOrder
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", id)
	err := db.First(&refundOrder).Error
	return refundOrder, err
}

// GetRefundOrderByIds 根据ids查询订单
func GetRefundOrderByIds(ids []uint) (model.RefundOrder, error) {
	var refundOrder model.RefundOrder
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id in ?", ids)
	err := db.First(&refundOrder).Error
	return refundOrder, err
}

// RemarkRefundOrder 备注退款单
func RemarkRefundOrder(id uint, merMark map[string]interface{}) error {
	db := g.TENANCY_DB.Model(&model.RefundOrder{})

	return db.Where("id = ?", id).Updates(merMark).Error
}

// GetRefundPriceByOrderIds 获取已退金额
func GetRefundPriceByOrderIds(ids []uint) (float64, error) {
	var wxPayPrice request.Result
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("sum(refund_price) as count")
	err := db.Where("status = ?", model.RefundStatusEnd).Where("order_id in ?", ids).First(&wxPayPrice).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return wxPayPrice.Count, err
	}
	return wxPayPrice.Count, nil
}

// checkRefundPrice 检查退款金额
func checkRefundPrice(refundOrder model.RefundOrder) (float64, error) {
	order, err := GetOrderById(refundOrder.OrderID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("当前订单不存在")
	} else if err != nil {
		return 0, err
	}
	// 已退款金额
	refundPrice, err := GetRefundPriceByOrderIds([]uint{refundOrder.OrderID})
	if err != nil {
		g.TENANCY_LOG.Error("检查退款金额", zap.String("GetRefundPriceByOrderIds()", err.Error()))
		return 0, fmt.Errorf("get refund order price %w", err)
	}
	payPrice := decimal.NewFromFloat(order.PayPrice)
	refundPriceD := decimal.NewFromFloat(refundPrice)
	pefundPrice := decimal.NewFromFloat(refundOrder.RefundPrice)
	if payPrice.Sub(refundPriceD).LessThan(pefundPrice) {
		g.TENANCY_LOG.Error("退款金额超出订单可退金额", zap.String("订单支付金额", payPrice.String()), zap.String("已退款金额", refundPriceD.String()), zap.String("当前退款金额", pefundPrice.String()))
		return 0, fmt.Errorf("退款金额超出订单可退金额")
	}
	return refundPrice, nil
}

// GetOtherRefundOrderIds 获取其他退款单ID集合
func GetOtherRefundOrderIds(orderId, refundOrderId uint) ([]uint, error) {
	ids := []uint{}
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("id").Where("order_id = ?", orderId).
		Where("status in ?", []int{model.RefundStatusAudit, model.RefundStatusAgree, model.RefundStatusBackgood})
	if refundOrderId > 0 {
		db = db.Where("id != ?", refundOrderId)
	}
	err := db.Find(&ids).Error
	if err != nil {
		return ids, fmt.Errorf("get other refund order ids %w", err)
	}
	return ids, nil
}

// AuditRefundOrder 审核退款单
func AuditRefundOrder(id uint, audit request.OrderAudit, refundAgreeMsg string) error {
	refundOrder, err := GetRefundOrderById(id)
	if err != nil {
		return fmt.Errorf("get refund order %w", err)
	}

	if audit.Status == model.RefundStatusAgree {
		err := agreeRefundOrder(refundOrder, refundAgreeMsg)
		if err != nil {
			return err
		}
	} else if audit.Status == model.RefundStatusRefuse {
		err := refuseRefundOrder(refundOrder, audit.FailMessage)
		if err != nil {
			return err
		}
	}

	return nil
}

// refuseRefundOrder 拒绝退款
//    2.1 如果退款数量 等于 购买数量 返还可退款数 is_refund = 0
//    2.2 商品总数小于可退数量 返还可退数 以商品数为准
//    2.3 是否存在其他图款单,是 ,退款中 ,否, 部分退款
func refuseRefundOrder(refundOrder model.RefundOrder, failMessage string) error {
	status := refundOrder.Status

	// 其他退款订单
	refundOrderIds, err := GetOtherRefundOrderIds(refundOrder.OrderID, refundOrder.ID)
	if err != nil {
		return err
	}
	refundProducts, err := getRefundProducts([]uint{refundOrder.ID})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单商品状态
		for _, refundProduct := range refundProducts {
			var isRefund int

			refundNum := refundProduct.RefundNum + refundProduct.OrderProduct.RefundNum //返还可退款数 当前退款数+可退款数量
			// 如果商品数量等于退款数量，商品变成初始未退款状态
			if refundProduct.OrderProduct.ProductNum == refundNum {
				isRefund = 0
			}

			// 否则，可退款数量退货数量等于订单商品数量
			if refundProduct.OrderProduct.ProductNum < refundNum {
				refundNum = refundProduct.OrderProduct.ProductNum
			}

			// 判断如果还有其他退款单，商品状态变为退款中
			if len(refundOrderIds) > 0 {
				var count int64
				err := g.TENANCY_DB.Model(&model.RefundProduct{}).Where("refund_order_id in ?", refundOrderIds).Where("order_product_id = ?", refundProduct.ProductID).Count(&count).Error
				if err != nil {
					return fmt.Errorf("get check refund product %w", err)
				}
				if count > 0 {
					isRefund = 1
				}
			}
			refundProduct.OrderProduct.IsRefund = isRefund
			err := UpdateOrderProduct(tx, refundProduct.OrderProduct.ID, map[string]interface{}{"is_refund": isRefund, "refund_num": refundNum})
			if err != nil {
				return fmt.Errorf("update refund product is_refund %w", err)
			}
		}
		status = model.RefundStatusRefuse
		err := UpdateRefundOrderById(tx, refundOrder.ID, map[string]interface{}{"status": status, "status_time": time.Now(), "fail_message": failMessage})
		if err != nil {
			return fmt.Errorf("update refund order status %w", err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func GetRefundProductCountByRefundOrderIds(refundOrderIds []uint, refundProductId uint) (int64, error) {
	var count int64
	err := g.TENANCY_DB.Model(&model.RefundProduct{}).Where("refund_order_id in ?", refundOrderIds).Where("order_product_id = ?", refundProductId).Count(&count).Error
	if err != nil {
		return count, fmt.Errorf("get check refund product %w", err)
	}
	return count, nil
}

// agreeRefundOrder 同意退款
//    1.1 仅退款
//       1.1.1 是 , 如果退款数量 等于 购买数量 is_refund = 3 全退退款 不等于 is_refund = 2 部分退款
//       1.1.2 否, is_refund = 1 退款中
//    1.2 退款退货 is_refund = 1
//    修改商品库存和销量
//    修改商品规格库存和销量
func agreeRefundOrder(refundOrder model.RefundOrder, refundAgreeMsg string) error {
	status := refundOrder.Status
	refundPrice, err := checkRefundPrice(refundOrder) // 可退款金额
	if err != nil {
		return err
	}
	refundProducts, err := getRefundProducts([]uint{refundOrder.ID})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单商品状态
		// isRefund  0:未退款 1:退款中 2:部分退款 3=全退
		for _, refundProduct := range refundProducts {
			var isRefund int

			if refundOrder.RefundType == model.RefundTypeTK { // 退款
				if refundProduct.RefundNum == refundProduct.BaseOrderProduct.ProductNum { //全退
					isRefund = 3
				} else { //部分退款
					isRefund = 2
				}
			}

			refundProduct.OrderProduct.IsRefund = isRefund
			err := UpdateOrderProduct(tx, refundProduct.OrderProduct.ID, map[string]interface{}{"is_refund": isRefund})
			if err != nil {
				return err
			}
			// 修改商品库存和销量
			if err = IncStock(tx, refundProduct.ProductID, refundProduct.ProductNum); err != nil {
				return err
			}
			// 修改商品规格库存和销量
			if err = IncSkuStock(tx, refundProduct.ProductID, refundProduct.Unique, refundProduct.ProductNum); err != nil {
				return err
			}
		}

		// 更新退款单状态
		status = model.RefundStatusAgree
		if refundOrder.RefundType == model.RefundTypeTK { // 只退款，直接原路退款给用户
			// 获取订单支付类型，获取订单号
			// 生成支付平台退款单
			order, err := GetOrderById(refundOrder.OrderID)
			if err != nil {
				return err
			}

			// 申请退款
			err = RefundOrder(order, refundOrder, refundPrice)
			if err != nil {
				return err
			}
		}
		err := addRefundOrderStatus(tx, refundOrder.ID, "refund_agree", refundAgreeMsg)
		if err != nil {
			return fmt.Errorf("add refund order status %w", err)
		}

		err = UpdateRefundOrderById(tx, refundOrder.ID, map[string]interface{}{"status": status, "status_time": time.Now()})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateRefundOrderById 更新退款单
func UpdateRefundOrderById(db *gorm.DB, refundOrderId uint, data map[string]interface{}) error {
	err := db.Model(&model.RefundOrder{}).Where("id = ?", refundOrderId).Updates(data).Error
	if err != nil {
		return fmt.Errorf("update refund order status %w", err)
	}
	return nil
}

// CreateRefundOrder 添加退款单
func CreateRefundOrder(reqId request.GetById, req request.CreateRefundOrder, orderProducts []response.OrderProduct) (uint, error) {
	var returnOrderId uint
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		refundOrder := model.RefundOrder{
			OrderID:      reqId.Id,
			PatientID:    reqId.PatientId,
			SysUserID:    reqId.UserId,
			SysTenancyID: reqId.TenancyId,
			BaseRefundOrder: model.BaseRefundOrder{
				RefundOrderSn: g.CreateOrderSn("R"),
				Mark:          req.Mark,
				MerMark:       "",
				AdminMark:     "",
				RefundType:    req.RefundType,
				RefundMessage: req.RefundMessage,
				RefundPrice:   req.RefundPrice,
				RefundNum:     req.Num,
				FailMessage:   "",
				Status:        model.RefundStatusAudit,
				StatusTime:    time.Now(),
			},
		}
		err := tx.Model(&model.RefundOrder{}).Create(&refundOrder).Error
		if err != nil {
			g.TENANCY_LOG.Error("添加退款单错误", zap.String("CreateRefundOrder()", err.Error()))
			return fmt.Errorf("添加退款单错误 %w", err)
		}

		err = addRefundOrderStatus(tx, refundOrder.ID, "create", "创建退款单")
		if err != nil {
			return fmt.Errorf("add refund order status %w", err)
		}

		var refundProducts []model.RefundProduct
		for _, orderProduct := range orderProducts {
			refundProduct := model.RefundProduct{
				RefundOrderID:  refundOrder.ID,
				OrderProductID: orderProduct.ProductID,
				RefundNum:      orderProduct.ProductNum,
			}
			refundProducts = append(refundProducts, refundProduct)
		}
		err = CreateRefundProduct(tx, refundProducts)
		if err != nil {
			return err
		}

		returnOrderId = refundOrder.ID
		return nil
	})
	if err != nil {
		return returnOrderId, err
	}
	return returnOrderId, nil
}

// CreateRefundProduct 新建退款单商品
func CreateRefundProduct(db *gorm.DB, refundProducts []model.RefundProduct) error {
	err := db.Model(&model.RefundProduct{}).Create(&refundProducts).Error
	if err != nil {
		g.TENANCY_LOG.Error("添加退款商品错误", zap.String("CreateRefundProduct()", err.Error()))
		return fmt.Errorf("添加退款商品错误: %v", err)
	}
	return nil
}

// DeleteRefundOrder 删除退款单
func DeleteRefundOrder(id uint) error {
	return g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", id).Update("is_system_del", g.StatusTrue).Error
}

// GetRefundOrderAutoAgree 获取超时退款单
func GetRefundOrderAutoAgree() ([]model.RefundOrder, error) {
	whereCreatedAt := fmt.Sprintf("now() > SUBDATE(created_at,interval -%s DAY)", param.GetRefundOrderAutoAgreeTime())
	refundOrder := []model.RefundOrder{}
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Where("status = ?", model.RefundStatusAudit).
		Where(whereCreatedAt).
		Find(&refundOrder).Error
	if err != nil {
		g.TENANCY_LOG.Error("获取超时退款单错误", zap.String("GetRefundOrderAutoAgree()", err.Error()))
		return refundOrder, fmt.Errorf("获取超时退款单错误: %v", err)
	}
	return refundOrder, nil
}

// AutoAgreeRefundOrders 退款单自动审核
func AutoAgreeRefundOrders(refundOrders []model.RefundOrder) {
	for _, refundOrder := range refundOrders {
		agreeRefundOrder(refundOrder, "审核通过[自动]")
	}
}

// addRefundOrderStatus 添加退款单记录
func addRefundOrderStatus(db *gorm.DB, id uint, cahngeType, changeMessage string) error {
	status := model.RefundStatus{
		RefundOrderID: id,
		ChangeType:    cahngeType,
		ChangeMessage: changeMessage,
		ChangeTime:    time.Now(),
	}
	err := db.Model(&model.RefundStatus{}).Create(&status).Error
	if err != nil {
		g.TENANCY_LOG.Error("添加退款单记录错误", zap.String("addRefundOrderStatus()", err.Error()))
		return fmt.Errorf("添加退款单记录错误 %w", err)
	}
	return nil
}

// GetStatusAgreeRefundOrdersByOrderSn 获取审核通过订单失败
func GetStatusAgreeRefundOrdersByOrderSn(orderSn string, payType int) ([]model.RefundOrder, error) {
	var refundOrders []model.RefundOrder
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Joins("left join orders on orders.id = refund_orders.order_id").
		Where("refund_orders.status = ?", model.RefundStatusAgree).
		Where("orders.order_sn = ?", orderSn).
		Where("orders.pay_type = ?", payType).
		Find(&refundOrders).Error
	if err != nil {
		g.TENANCY_LOG.Error("获取审核通过订单失败", zap.String("GetStatusAgreeRefundOrdersByOrderSn()", err.Error()))
		return refundOrders, fmt.Errorf("获取审核通过订单失败 %w", err)
	}
	return refundOrders, nil

}

// GetStatusAgreeRefundOrdersByReturnOrderSn 获取审核通过订单失败
func GetStatusAgreeRefundOrdersByReturnOrderSn(refundOrderSn string, payType int) ([]model.RefundOrder, error) {
	var refundOrders []model.RefundOrder
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Joins("left join orders on orders.id = refund_orders.order_id").
		Where("refund_orders.status = ?", model.RefundStatusAgree).
		Where("refund_orders.refund_order_sn = ?", refundOrderSn).
		Where("orders.pay_type = ?", payType).
		Find(&refundOrders).Error
	if err != nil {
		g.TENANCY_LOG.Error("获取审核通过订单失败", zap.String("GetStatusAgreeRefundOrdersByReturnOrderSn()", err.Error()))
		return refundOrders, fmt.Errorf("获取审核通过订单失败 %w", err)
	}
	return refundOrders, nil
}

// ChangeReturnOrderStatusByOrderSn 修改退款订单状态 by orderSn
func ChangeReturnOrderStatusByOrderSn(payType, status int, orderSn, changeType, changeMessage string) error {
	refundOrders, err := GetStatusAgreeRefundOrdersByOrderSn(orderSn, payType)
	if err != nil {
		return err
	}
	if len(refundOrders) != 1 {
		return fmt.Errorf("%s 订单号重复生产 %d 个订单", orderSn, len(refundOrders))
	}
	err = ChangeRefundStatusById(refundOrders[0].ID, map[string]interface{}{"status": status, "status_time": time.Now()}, changeType, changeMessage)
	if err != nil {
		return err
	}

	return nil
}

// ChangeReturnOrderStatusByReturnOrderSn 修改退款订单状态 by refundOrderSn
func ChangeReturnOrderStatusByReturnOrderSn(payType, status int, refundOrderSn, changeType, changeMessage string) error {
	refundOrders, err := GetStatusAgreeRefundOrdersByReturnOrderSn(refundOrderSn, payType)
	if err != nil {
		return err
	}
	if len(refundOrders) != 1 {
		return fmt.Errorf("%s 退款订单号重复生产 %d 个订单", refundOrderSn, len(refundOrders))
	}
	err = ChangeRefundStatusById(refundOrders[0].ID, map[string]interface{}{"status": status, "status_time": time.Now()}, changeType, changeMessage)
	if err != nil {
		return err
	}
	return nil
}

// ChangeRefundStatusById 修改退款订单状态 by refundOrderId
func ChangeRefundStatusById(refundOrderId uint, changeData map[string]interface{}, changeType, changeMessage string) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := UpdateRefundOrderById(tx, refundOrderId, changeData)
		if err != nil {
			return err
		}
		err = addRefundOrderStatus(tx, refundOrderId, changeType, changeMessage)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
