package public

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

func PayOrder(ctx *gin.Context) {
	var req request.PayOrder
	if err := ctx.ShouldBind(&req); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}

	req.UserAgent = ctx.Request.UserAgent()
	if strings.Contains(req.UserAgent, "MicroMessenger") {
		if req.Code == "" || req.State == "" {
			url, err := service.GetAutoCode(ctx.Request.RequestURI)
			if err != nil {
				g.TENANCY_LOG.Error("获取微信autoCode错误!", zap.Any("获取微信autoCode错误", err))
				response.FailWithMessage("操作失败:"+err.Error(), ctx)
				return
			} else {
				ctx.Redirect(http.StatusFound, url)
				return
			}
		} else {
			if req.State != g.TENANCY_CONFIG.WechatPay.State {
				response.FailWithMessage("微信网页授权验证错误", ctx)
				return
			}
			openid, err := service.GetOpenId(req.Code)
			if err != nil {
				g.TENANCY_LOG.Error("获取微信openid错误!", zap.Any("获取微信openid错误", err))
				response.FailWithMessage("操作失败:"+err.Error(), ctx)
				return
			}
			req.OpenId = openid
		}
	}

	if res, err := service.PayOrder(req); err != nil {
		g.TENANCY_LOG.Error("支付订单错误!", zap.Any("支付订单错误", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		if res.AliPayUrl != "" {
			ctx.Redirect(http.StatusFound, res.AliPayUrl)
		} else {
			jsapi := map[string]interface{}{
				"appId":     res.JSAPIPayParams.AppId,     //公众号名称，由商户传入
				"timeStamp": res.JSAPIPayParams.TimeStamp, //时间戳，自1970年以来的秒数
				"nonceStr":  res.JSAPIPayParams.NonceStr,  //随机串
				"package":   res.JSAPIPayParams.Package,
				"signType":  res.JSAPIPayParams.SignType, //微信签名方式：
				"paySign":   res.JSAPIPayParams.PaySign,  //微信签名
			}
			ctx.HTML(http.StatusOK, "wechat_pay.tmpl", jsapi)
		}
	}
}

func NotifyAliPay(ctx *gin.Context) {
	if err := service.NotifyAliPay(ctx); err != nil {
		g.TENANCY_LOG.Error("支付宝支付异步通知失败!", zap.Any("支付宝支付异步通知失败", err))
		ctx.String(http.StatusOK, "%s", err.Error())
	} else {
		ctx.String(http.StatusOK, "%s", "success")
	}
}

func NotifyWechatPay(ctx *gin.Context) {
	if err := service.NotifyWechatPay(ctx); err != nil {
		g.TENANCY_LOG.Error("微信支付异步通知失败!", zap.Any("微信支付异步通知失败", err))
		ctx.String(http.StatusOK, "%s", err.Error())
	} else {
		ctx.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
	}
}
