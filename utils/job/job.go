package job

import (
	"fmt"
	"path/filepath"

	"github.com/chindeo/pkg/file"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
)

const EveryMinute = "0 */1 * * * *"    //每分钟
const EveryTeenHour = "0 0 */10 * * *" //每个11小时

func Timer() {
	if g.TENANCY_CONFIG.Timer.Start {
		for _, detail := range g.TENANCY_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				g.TENANCY_Timer.AddTaskByFunc("ClearDB", g.TENANCY_CONFIG.Timer.Spec, "清除操作记录日志", func() {
					err := utils.ClearTable(g.TENANCY_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						g.TENANCY_LOG.Info("清除操作记录日志", zap.String("错误", err.Error()))
					}
				})
			}(detail)
		}

		// 订单过期自动取消
		g.TENANCY_Timer.AddTaskByFunc("NoPayOrderAutoClose", EveryMinute, "订单过期自动取消", func() {
			groupOrderIds, err := service.GetNoPayGroupOrderAutoClose(false)
			if err != nil {
				g.TENANCY_LOG.Info("订单过期自动取消", zap.String("获取订单错误", err.Error()))
				return
			}
			if len(groupOrderIds) == 0 {
				return
			}
			for _, groupOrderId := range groupOrderIds {
				err := service.CancelNoPayGroupOrders(groupOrderId)
				if err != nil {
					g.TENANCY_LOG.Info("订单过期自动取消", zap.String("订单状态更新错误", err.Error()))
				}
			}
		})

		// 商户自动处理退款订单，自动通过退款审核
		g.TENANCY_Timer.AddTaskByFunc("RefundOrderAutoAgree", EveryMinute, "商户自动处理退款订单", func() {
			refundOrders, err := service.GetRefundOrderAutoAgree()
			if err != nil {
				g.TENANCY_LOG.Info("商户自动处理退款订单", zap.String("获取退款订单错误", err.Error()))
				return
			}
			if len(refundOrders) == 0 {
				return
			}
			service.AutoAgreeRefundOrders(refundOrders)
		})

		// 用户订单自动收货，自动确认收货
		// - 发货前：如果没有发生退款售后，按设定日期自动确认收货（正常流程），发货后：申请售后，不影响订单流程订单，正常收货完成。
		// - 发货前：如果有退款售后，并且已经全额退款完成，才会按设定日期自动确认收货。
		g.TENANCY_Timer.AddTaskByFunc("OrderAutoAgree", EveryMinute, "用户订单自动收货", func() {
			orderIds, err := service.GetOrderAutoAgree()
			if err != nil {
				g.TENANCY_LOG.Info("用户订单自动收货", zap.String("获取订单错误", err.Error()))
				return
			}
			if len(orderIds) == 0 {
				return
			}
			err = service.AutoTakeOrders(orderIds)
			if err != nil {
				g.TENANCY_LOG.Info("用户订单自动收货", zap.String("订单状态更新错误", err.Error()))
			}
		})

		// 定时获取微信平台证书
		g.TENANCY_Timer.AddTaskByFunc("CheckOrdersPayStatus", EveryTeenHour, "定时获取微信平台证书", func() {
			wechatConf, err := param.GetWechatPayConfig()
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
