package logic

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
	"gorm.io/gorm"
)

// CheckRefundOrder 申请退款结算页面
// 逻辑：付款，未取消，未完成订单支持退款申请
func CheckRefundOrder(req request.GetById, orderPorductIds []uint) (response.CheckRefundOrder, error) {
	var checkRefundOrder response.CheckRefundOrder
	order, orderProducts, err := checkOrder(req, orderPorductIds)
	if err != nil {
		return checkRefundOrder, err
	}

	checkRefundOrder.TotalRefundPrice = service.GetTotalRefundPrice(orderProducts)
	if order.Status >= model.OrderStatusNoReceive || order.OrderType == model.OrderTypeGeneral { // 发货,非自提订单需要减去邮费
		checkRefundOrder.PostagePrice = order.PayPostage
	}
	checkRefundOrder.Status = order.Status
	checkRefundOrder.Product = orderProducts

	return checkRefundOrder, nil
}

func checkOrder(req request.GetById, orderPorductIds []uint) (model.Order, []response.OrderProduct, error) {
	payScope := scope.SimpleScope("status", []int{model.OrderStatusNoDeliver, model.OrderStatusNoReceive, model.OrderStatusNoComment}, "in")
	order, err := service.GetOrderByOrderId(req, payScope)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, nil, fmt.Errorf("订单不存在")
	} else if err != nil {
		return order, nil, err
	}

	if order.IsCancel == g.StatusTrue {
		return order, nil, fmt.Errorf("订单已取消")
	}

	orderProducts, err := service.GetOrdersProductByProductIds(orderPorductIds, req)
	if err != nil {
		return order, nil, err
	}
	if len(orderProducts) == 0 {
		return order, nil, fmt.Errorf("商品不存在")
	}
	if len(orderProducts) != len(orderPorductIds) {
		return order, nil, fmt.Errorf("请选择正确的退款商品")
	}
	return order, orderProducts, nil
}

// CreateRefundOrder 提交退款申请
// 逻辑：付款，未取消，未完成订单支持退款申请
func CreateRefundOrder(reqId request.GetById, req request.CreateRefundOrder) (uint, error) {
	_, orderProducts, err := checkOrder(reqId, req.Ids)
	if err != nil {
		return 0, err
	}
	return service.CreateRefundOrder(reqId, req, orderProducts)
}
