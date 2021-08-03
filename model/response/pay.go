package response

import "github.com/go-pay/gopay/wechat/v3"

type PayOrder struct {
	AliPayUrl      string
	JSAPIPayParams *wechat.JSAPIPayParams
}
