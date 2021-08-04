package public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func PayOrder(ctx *gin.Context) {
	var req request.PayOrder
	if err := ctx.ShouldBind(&req); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	userAgent := ctx.Request.UserAgent()
	if res, err := service.PayOrder(req, userAgent, multi.GetTenancyName(ctx)); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
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
			fmt.Printf("wxRsp: %+v", jsapi)
			ctx.HTML(200, "wechat-pay.tmpl", jsapi)
		}
	}
}
