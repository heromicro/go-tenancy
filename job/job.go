package job

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/chindeo/pkg/file"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

const EveryMinute = "*/1 * * * *"    //每分钟
const EveryTeenHour = "0 */10 * * *" //每个11小时

func Timer() {
	if g.TENANCY_CONFIG.Timer.Start {
		for _, detail := range g.TENANCY_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				g.TENANCY_Timer.AddTaskByFunc("ClearDB", g.TENANCY_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(g.TENANCY_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						g.TENANCY_LOG.Info("清楚操作记录日志", zap.String("错误", err.Error()))
					}
				})
			}(detail)
		}

		// 订单过期自动取消
		g.TENANCY_Timer.AddTaskByFunc("CheckOrdersPayStatus", EveryMinute, func() {
			orders, err := service.GetNoPayOrders()
			if err != nil {
				g.TENANCY_LOG.Info("订单过期自动取消", zap.String("获取订单错误", err.Error()))
				return
			}
			if len(orders) == 0 {
				return
			}
			var orderIds []uint
			var orderStatues []model.OrderStatus
			for _, order := range orders {
				orderIds = append(orderIds, order.ID)
				orderStatus := model.OrderStatus{ChangeType: "cancel", ChangeMessage: "取消订单[自动]", ChangeTime: time.Now(), OrderID: order.ID}
				orderStatues = append(orderStatues, orderStatus)
			}
			err = service.CancelNoPayOrders(orderIds, orderStatues)
			if err != nil {
				g.TENANCY_LOG.Info("订单过期自动取消", zap.String("订单状态更新错误", err.Error()))
			}
		})

		// 定时获取微信平台证书
		g.TENANCY_Timer.AddTaskByFunc("CheckOrdersPayStatus", EveryTeenHour, func() {
			wechatConf, err := service.GetWechatPayConfig()
			if err != nil {
				g.TENANCY_LOG.Info("定时获取微信平台证书", zap.String("获取微信支付配置错误", err.Error()))
				return
			}

			pKContent, err := file.ReadString(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, wechatConf.PayWeixinClientKey))
			if err != nil {
				g.TENANCY_LOG.Info("定时获取微信平台证书", zap.String("获取微信证书内容错误", err.Error()))
				return
			}

			certs, err := wechat.GetPlatformCerts(wechatConf.PayWeixinMchid, wechatConf.PayWeixinKey, "", pKContent)
			if err != nil {
				g.TENANCY_LOG.Info("定时获取微信平台证书", zap.String("定时获取微信平台证书错误", err.Error()))
				return
			}
			for _, cert := range certs.Certs {
				fmt.Printf("%+v\n", cert)
			}
		})
	}
}
