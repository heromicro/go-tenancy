package initialize

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/chindeo/pkg/file"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/job"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

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
		g.TENANCY_Timer.AddTaskByFunc("CheckOrdersPayStatus", job.EveryMinute, func() {
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
			err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
				err := tx.Model(&model.Order{}).
					Where("id in ?", orderIds).
					Update("status", model.OrderStatusCancel).Error
				if err != nil {
					return err
				}
				err = tx.Create(&orderStatues).Error
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				g.TENANCY_LOG.Info("订单过期自动取消", zap.String("订单状态更新错误", err.Error()))
			}
		})

		// 定时获取微信平台证书
		g.TENANCY_Timer.AddTaskByFunc("CheckOrdersPayStatus", job.EveryTeenHour, func() {
			wechatConf, err := service.GetWechatPay()
			if err != nil {
				g.TENANCY_LOG.Info("定时获取微信平台证书", zap.String("获取支付宝配置错误", err.Error()))
				return
			}

			pKContent, err := file.ReadString(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, wechatConf["pay_weixin_client_key"]))
			certs, err := wechat.GetPlatformCerts(wechatConf["pay_weixin_mchid"], wechatConf["pay_weixin_key"], "", pKContent)
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
