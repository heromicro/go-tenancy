package job

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// 添加订单定时任务
type SendMqtt struct {
	OrderId   uint `json:"orderId" form:"orderId"`
	TenancyId uint `json:"tenancyId" form:"tenancyId"`
	UserId    uint `json:"userId" form:"userId"`
	OrderType int  `json:"orderType" form:"orderType"`
	CreatedAt time.Time
}

func (d SendMqtt) Run() {
	mqtts, err := service.GetStatusMqtts()
	if err != nil {
		g.TENANCY_LOG.Error("mqtt消息发送队列任务", zap.String("获取mqtt列表", err.Error()))
	}
	if len(mqtts) > 0 {
		for _, mqtt := range mqtts {
			fmt.Println(mqtt)
		}
	}
}
