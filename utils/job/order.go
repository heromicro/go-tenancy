package job

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// CheckOrderPayStatus 暂时未使用
type CheckOrderPayStatus struct {
	OrderId   uint `json:"orderId" form:"orderId"`
	TenancyId uint `json:"tenancyId" form:"tenancyId"`
	UserId    uint `json:"userId" form:"userId"`
	OrderType int  `json:"orderType" form:"orderType"`
	CreatedAt time.Time
}

func (d CheckOrderPayStatus) Run() {
	if time.Since(d.CreatedAt).Minutes() >= 15 {
		changeData := map[string]interface{}{"is_cancel": g.StatusTrue}
		err := service.ChangeOrderStatusByOrderId(d.OrderId, changeData, model.ChangeTypeCancel, "取消订单[自动]")
		if err != nil {
			g.TENANCY_LOG.Error("定时自动取消任务错误", zap.String("自动取消订单任务", err.Error()))
		}
	}
}
