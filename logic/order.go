package logic

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
	"gorm.io/gorm"
)

// PayOrder 结算订单，生成支付二维码
// 逻辑： 支付未付款，未取消的，本人当前商户的订单，
func PayOrder(req request.GetById) (string, error) {
	order, err := service.GetOrderById(req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("订单不存在")
	} else if err != nil {
		return "", err
	}

	if order.IsCancel == g.StatusTrue {
		return "", fmt.Errorf("订单已取消")
	}

	if order.Status > model.OrderStatusNoPay {
		return "", fmt.Errorf("订单已经支付，请勿重复支付")
	}
	qrcode, err := service.GetQrCode(req.Id, req.TenancyId, order.OrderType)
	if err != nil {
		return "", err
	}
	return string(qrcode), nil
}

// CancelOrder 取消订单，取消未支付订单
// 逻辑： 取消未支付，未取消订单，本人当前商户的订单，
func CancelOrder(req request.GetById) error {
	noPayScope := scope.SimpleScope("is_cancel", g.StatusFalse)
	order, err := service.GetOrderById(req.Id, noPayScope)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("订单不存在")
	} else if err != nil {
		return err
	}

	if order.Status > model.OrderStatusNoPay {
		return fmt.Errorf("订单已经支付，无法取消")
	}
	if err := service.CancelOrder(req); err != nil {
		return err
	}
	return nil
}
