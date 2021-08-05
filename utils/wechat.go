package utils

import (
	"net/url"
)

// GetAutoCode 微信网页授权获取code
//	appId：应用唯一标识，在微信开放平台提交应用审核通过后获得
//	redirectUri：授权后重定向的回调链接地址， 请使用 urlEncode 对链接进行处理
//	scope：应用授权作用域，snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid），snsapi_userinfo （弹出授权页面，可通过openid拿到昵称、性别、所在地。并且， 即使在未关注的情况下，只要用户授权，也能获取其信息 ）
//	state：重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
//	文档：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func GetAutoCode(appId, redirectUri, scope, state string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appId + "&redirect_uri=" + url.QueryEscape(redirectUri) + "&response_type=code&scope=" + scope + "&state=" + state + "#wechat_redirect"
}
