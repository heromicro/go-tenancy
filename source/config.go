package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Config = new(config)

type config struct{}

var configs = []model.SysConfig{
	{SysConfigCategoryID: 2, ConfigName: "网站域名", ConfigKey: "site_url", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 2, ConfigName: "网站名称", ConfigKey: "site_name", ConfigType: "input", ConfigRule: "", Required: 1, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 2, ConfigName: "网站模式", ConfigKey: "site_open", ConfigType: "radio", ConfigRule: "0:测试;1:生产", Required: 1, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号名称", ConfigKey: "wechat_name", ConfigType: "input", ConfigRule: "", Required: 2, Info: "设置公众号名称", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "联系电话", ConfigKey: "set_phone", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "联系邮箱", ConfigKey: "set_email", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信号", ConfigKey: "wechat_id", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号原始id", ConfigKey: "wechat_sourceid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号原始id", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号AppID", ConfigKey: "wechat_appid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号AppID", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号AppSecret", ConfigKey: "wechat_appsecret", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号AppSecret", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信验证TOKEN", ConfigKey: "wechat_token", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信验证TOKEN", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信EncodingAESKey", ConfigKey: "wechat_encodingaeskey", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号消息加解密Key,在使用安全模式情况下要填写该值，请先在管理中心修改，然后填写该值，仅限服务号和认证订阅号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号二维码", ConfigKey: "wechat_qrcode", ConfigType: "image", ConfigRule: "", Required: 2, Info: "公众号二维码", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号logo", ConfigKey: "wechat_avatar", ConfigType: "image", ConfigRule: "", Required: 2, Info: "公众号logo", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享图片", ConfigKey: "wechat_share_img", ConfigType: "image", ConfigRule: "", Required: 2, Info: "若填写此图片地址，则分享网页出去时会分享此图片。可有效防止分享图片变形", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享标题", ConfigKey: "wechat_share_title", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信分享标题", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享简介", ConfigKey: "wechat_share_synopsis", ConfigType: "textarea", ConfigRule: "", Required: 2, Info: "微信分享简介", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "消息加解密方式", ConfigKey: "wechat_encode", ConfigType: "radio", ConfigRule: "0:明文模式;1:兼容模式;2:安全模式", Required: 1, Info: "如需使用安全模式请在管理中心修改，仅限服务号和认证订阅号", Sort: 1, UserType: 2, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "警戒库存", ConfigKey: "mer_store_stock", ConfigType: "input", ConfigRule: "", Required: 0, Info: "设置商品的警戒库存", Sort: 0, UserType: 1, Status: 1},
	// 198	0	短信平台账号	sms_account	input		0	设置短信平台账号	0	0	0	2020-05-18 15:33:58
	// 199	0	短信平台密码	sms_token	input		0	设置短信平台密码	0	0	0	2020-05-18 15:34:22
	{SysConfigCategoryID: 6, ConfigName: "Appid", ConfigKey: "pay_weixin_appid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信公众号身份的唯一标识。审核通过后，在微信发送的邮件中查看。", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "Appsecret", ConfigKey: "pay_weixin_appsecret", ConfigType: "input", ConfigRule: "", Required: 2, Info: "JSAPI接口中获取openid，审核后在公众平台开启开发模式后可查看。", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "Mchid", ConfigKey: "pay_weixin_mchid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "受理商ID，身份标识", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "微信支付证书", ConfigKey: "pay_weixin_client_cert", ConfigType: "file", ConfigRule: "", Required: 2, Info: "微信支付证书，在微信商家平台中可以下载！文件名一般为apiclient_cert.pem", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "微信支付证书密钥", ConfigKey: "pay_weixin_client_key", ConfigType: "file", ConfigRule: "", Required: 2, Info: "微信支付证书密钥，在微信商家平台中可以下载！文件名一般为apiclient_key.pem", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "V3Key", ConfigKey: "pay_weixin_key", ConfigType: "input", ConfigRule: "", Required: 2, Info: "商户支付密钥v3Key", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "SerialNo", ConfigKey: "pay_serial_no", ConfigType: "input", ConfigRule: "", Required: 2, Info: "商户证书序列号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 6, ConfigName: "开启", ConfigKey: "pay_weixin_open", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否启用微信支付", Sort: 0, UserType: 2, Status: 1},
	// 215	12	充值注意事项	recharge_attention	textarea		0	充值注意事项	0	0	1	2020-06-03 09:49:01
	{SysConfigCategoryID: 1, ConfigName: "订单自动关闭时间", ConfigKey: "auto_close_order_timer", ConfigType: "number", ConfigRule: "", Required: 2, Info: "订单自动关闭时间(单位:分钟)", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "默认退货收货地址", ConfigKey: "mer_refund_address", ConfigType: "input", ConfigRule: "", Required: 2, Info: "设置默认退货收货地址", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "默认退货收货人", ConfigKey: "mer_refund_user", ConfigType: "input", ConfigRule: "", Required: 2, Info: "设置默认退货收货人", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "退款理由", ConfigKey: "refund_message", ConfigType: "textarea", ConfigRule: "", Required: 2, Info: "设置退款理由", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "商户自动处理退款订单期限（天）", ConfigKey: "mer_refund_order_agree", ConfigType: "number", ConfigRule: "", Required: 1, Info: "申请退款的订单超过期限，将自动退款处理。", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "银行卡开户行", ConfigKey: "bank", ConfigType: "input", ConfigRule: "", Required: 1, Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "银行卡卡号", ConfigKey: "bank_number", ConfigType: "input", ConfigRule: "", Required: 1, Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "银行卡持卡人姓名", ConfigKey: "bank_name", ConfigType: "input", ConfigRule: "", Required: 1, Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "银行卡开户行地址", ConfigKey: "bank_address", ConfigType: "input", ConfigRule: "", Required: 1, Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "快递查询密钥", ConfigKey: "express_app_code", ConfigType: "input", ConfigRule: "", Required: 2, Info: "阿里云快递查询接口密钥购买地址：https://market.aliyun.com/products/56928004/cmapi021863.html", Sort: 0, UserType: 2, Status: 1},
	// 228	7	空间域名 Domain	uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 10:21:19
	// 229	7	accessKey	accessKey	input		0	accessKey	0	0	1	2020-06-18 10:21:37
	// 230	7	secretKey	secretKey	input		0	secretKey	0	0	1	2020-06-18 10:22:40
	// 231	7	存储空间名称	storage_name	input		0	存储空间名称	0	0	1	2020-06-18 10:22:57
	// 232	7	所属地域	storage_region	input		0	所属地域	0	0	1	2020-06-18 10:23:21
	// 233	0	上传类型	upload_type	radio	1:本地存储
	// 2:七牛云存储
	// 3:阿里云OSS
	// 4:腾讯COS	0	文件上传的类型	0	0	1	2020-06-18 14:46:29
	// 237	8	空间域名 Domain	qiniu_uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 15:14:29
	// 238	8	accessKey	qiniu_accessKey	input		0	accessKey	0	0	1	2020-06-18 15:14:51
	// 239	8	secretKey	qiniu_secretKey	input		0	secretKey	0	0	1	2020-06-18 15:15:05
	// 240	8	存储空间名称	qiniu_storage_name	input		0	存储空间名称	0	0	1	2020-06-18 15:15:22
	// 241	8	所属地域	qiniu_storage_region	input		0	所属地域	0	0	1	2020-06-18 15:15:40
	// 242	9	空间域名 Domain	tengxun_uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 15:40:03
	// 243	9	accessKey	tengxun_accessKey	input		0	accessKey	0	0	1	2020-06-18 15:40:32
	// 244	9	secretKey	tengxun_secretKey	input		0	secretKey	0	0	1	2020-06-18 15:40:48
	// 245	9	存储空间名称	tengxun_storage_name	input		0	存储空间名称	0	0	1	2020-06-18 15:41:06
	// 246	9	所属地域	tengxun_storage_region	input		0	所属地域	0	0	1	2020-06-18 15:41:39
	{SysConfigCategoryID: 10, ConfigName: "appId", ConfigKey: "routine_appId", ConfigType: "input", ConfigRule: "", Required: 2, Info: "appId", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 10, ConfigName: "小程序AppSecret", ConfigKey: "routine_appsecret", ConfigType: "input", ConfigRule: "", Required: 2, Info: "小程序AppSecret", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 10, ConfigName: "小程序授权logo", ConfigKey: "routine_logo", ConfigType: "image", ConfigRule: "", Required: 2, Info: "小程序授权logo", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 10, ConfigName: "小程序名称", ConfigKey: "routine_name", ConfigType: "input", ConfigRule: "", Required: 2, Info: "小程序名称", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "Appid", ConfigKey: "pay_routine_appid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "小程序Appid", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "Appsecret", ConfigKey: "pay_routine_appsecret", ConfigType: "input", ConfigRule: "", Required: 2, Info: "小程序Appsecret", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "Mchid", ConfigKey: "pay_routine_mchid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "商户号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "Key", ConfigKey: "pay_routine_key", ConfigType: "input", ConfigRule: "", Required: 2, Info: "商户key", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "小程序支付证书", ConfigKey: "pay_routine_client_cert", ConfigType: "file", ConfigRule: "", Required: 2, Info: "小程序支付证书", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 11, ConfigName: "小程序支付证书密钥", ConfigKey: "pay_routine_client_key", ConfigType: "file", ConfigRule: "", Required: 2, Info: "小程序支付证书密钥", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 12, ConfigName: "余额充值开关", ConfigKey: "recharge_switch", ConfigType: "radio", ConfigRule: "1:开启;0:关闭", Required: 2, Info: "余额充值开关", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 12, ConfigName: "用户最低充值金额", ConfigKey: "store_user_min_recharge", ConfigType: "number", ConfigRule: "", Required: 2, Info: "用户最低充值金额", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 12, ConfigName: "余额功能启用", ConfigKey: "balance_func_status", ConfigType: "radio", ConfigRule: "1:开启;0:关闭", Required: 2, Info: "商城余额功能启用或者关闭", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 12, ConfigName: "余额支付状态", ConfigKey: "yue_pay_status", ConfigType: "radio", ConfigRule: "1:开启;0:关闭", Required: 2, Info: "余额支付状态", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "首页广告图", ConfigKey: "home_ad_pic", ConfigType: "image", ConfigRule: "", Required: 2, Info: "设置首页广告图(750*164px)", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "首页广告链接", ConfigKey: "home_ad_url", ConfigType: "input", ConfigRule: "", Required: 2, Info: "设置首页广告链接", Sort: 0, UserType: 2, Status: 1},
	// 264	13	分销说明	promoter_explain	textarea		0		0	0	1	2020-06-25 15:32:21
	// 265	13	商户设置礼包最大数量	max_bag_number	number		0		0	0	1	2020-06-25 15:48:27
	// 266	2	商城 logo	site_logo	image		0	设置商城logo(254*90px)	0	0	1	2020-06-26 11:41:23
	// 268	2	商城分享标题	share_title	input		0	商城分享标题	0	0	1	2020-06-30 16:16:30
	// 269	2	商城分享简介	share_info	input		0	商城分享简介	0	0	1	2020-06-30 16:16:57
	// 270	2	商城分享图片	share_pic	image		0	商城分享图片	0	0	1	2020-06-30 16:17:23
	// 271	4	发货提醒	sms_fahuo_status	radio	0:关闭
	// 1:开启	0	发货提醒	0	0	1	2020-07-01 11:00:13
	// 272	4	确认收货短信提醒	sms_take_status	radio	0:关闭
	// 1:开启	0	确认收货短信提醒	0	0	1	2020-07-01 11:20:51
	// 273	4	用户下单通知提醒	sms_pay_status	radio	0:关闭
	// 1:开启	0	用户下单通知提醒	0	0	1	2020-07-01 11:25:32
	// 274	4	改价提醒	sms_revision_status	radio	0:关闭
	// 1:开启	0	改价提醒	0	0	1	2020-07-01 11:30:04
	// 275	4	提醒付款通知	sms_pay_false_status	radio	0:关闭
	// 1:开启	0	提醒付款通知	0	0	1	2020-07-01 11:42:41
	// 276	4	商家拒绝退款提醒	sms_refund_fail_status	radio	0:关闭
	// 1:开启	0	商家拒绝退款提醒	0	0	1	2020-07-01 11:55:05
	// 277	4	商家同意退款提醒	sms_refund_success_status	radio	0:关闭
	// 1:开启	0	商家同意退款提醒	0	0	1	2020-07-01 12:01:32
	// 278	4	退款确认提醒	sms_refund_confirm_status	radio	0:关闭
	// 1:开启	0	退款确认提醒	0	0	1	2020-07-01 12:06:14
	// 279	4	管理员下单提醒	sms_admin_pay_status	radio	0:关闭
	// 1:开启	0	管理员下单提醒	0	0	1	2020-07-01 12:17:42
	// 280	4	管理员退货提醒	sms_admin_return_status	radio	0:关闭
	// 1:开启	0	管理员退货提醒	0	0	1	2020-07-01 12:31:13
	// 281	4	管理员确认收货提醒	sms_admin_take_status	radio	0:关闭
	// 1:开启	0	管理员确认收货提醒	0	0	1	2020-07-01 12:32:01
	// 282	4	退货信息提醒	sms_admin_postage_status	radio	0:关闭
	// 1:开启	0	退货信息提醒	0	0	1	2020-07-01 14:04:27
	// 283	2	后台登录页logo	sys_login_logo	image		0	后台登录页logo	0	0	1	2020-07-06 15:20:18
	// 284	2	后台登录页标题	sys_login_title	input		0	后台登录页标题	0	0	1	2020-07-07 16:39:25
	{SysConfigCategoryID: 1, ConfigName: "菜单logo", ConfigKey: "sys_menu_logo", ConfigType: "image", ConfigRule: "", Required: 2, Info: "设置菜单顶部logo", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "菜单小logo", ConfigKey: "sys_menu_slogo", ConfigType: "image", ConfigRule: "", Required: 2, Info: "设置菜单顶部小logo", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "商户入驻协议", ConfigKey: "sys_intention_agree", ConfigType: "textarea", ConfigRule: "", Required: 2, Info: "商户入驻协议", Sort: 0, UserType: 2, Status: 0},
	{SysConfigCategoryID: 1, ConfigName: "开启商户入驻", ConfigKey: "mer_intention_open", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否开启商户入驻功能", Sort: 0, UserType: 2, Status: 1},
	// 289	4	预售尾款支付通知	sms_pay_presell_status	radio	0:关闭
	// 1:开启	0		1	0	1	2020-11-30 17:46:45
	{SysConfigCategoryID: 5, ConfigName: "打印机终端号", ConfigKey: "terminal_number", ConfigType: "input", Required: 2, Info: "打印机终端号", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "打印机应用ID", ConfigKey: "printing_client_id", ConfigType: "input", Required: 2, Info: "打印机开发者用户ID", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "打印机用户ID", ConfigKey: "develop_id", ConfigType: "input", Required: 2, Info: "打印机的应用ID", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "打印机密匙", ConfigKey: "printing_api_key", ConfigType: "input", Required: 2, Info: "打印机应用密匙", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "开启直播免审核", ConfigKey: "broadcast_room_type", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否开启直播免审核", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "开启复制第三方平台商品", ConfigKey: "copy_product_status", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否开启复制商品功能", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "复制商品接口KEY", ConfigKey: "copy_product_apikey", ConfigType: "input", ConfigRule: "", Required: 2, Info: "接口key", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "开启直播商品免审核", ConfigKey: "broadcast_goods_type", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否开启直播商品免审核", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "腾讯地图KEY", ConfigKey: "tx_map_key", ConfigType: "input", ConfigRule: "", Required: 2, Info: "腾讯地图KEY", Sort: 0, UserType: 2, Status: 1},
	// 299	14	开启门店自提	mer_take_status	radio	0:关闭
	// 1:开启	0	是否开启门店自提	0	1	1	2020-08-01 11:59:14
	// 300	14	自提点名称	mer_take_name	input		0	设置自提点名称	0	1	1	2020-08-01 12:01:08
	// 301	14	自提点手机号	mer_take_phone	input		0	设置自提点手机号	0	1	1	2020-08-01 12:02:04
	// 302	14	自提点地址	mer_take_address	input		0	设置自提点地址	0	1	1	2020-08-01 12:03:57
	// 303	14	店铺经纬度	mer_take_location	input		0	设置店铺经纬度	0	1	1	2020-08-01 12:16:18
	// 304	14	自提点营业日期	mer_take_day	input		0	设置自提点营业日期	0	1	1	2020-08-01 12:17:06
	// 305	14	自提点营业时间	mer_take_time	input		0	设置自提点营业时间	0	1	1	2020-08-01 12:17:39
	{SysConfigCategoryID: 1, ConfigName: "	订单自动收货时间(天)", ConfigKey: "auto_take_order_timer", ConfigType: "number", ConfigRule: "", Required: 2, Info: "设置订单自动收货时间(天)", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "默认赠送复制次数", ConfigKey: "copy_product_defaul", ConfigType: "number", ConfigRule: "", Required: 2, Info: "默认给商户赠送可用次数", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 1, ConfigName: "是否展示店铺", ConfigKey: "hide_mer_status", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否展示店铺", Sort: 0, UserType: 2, Status: 1},
	// 309	4	直播审核通过主播通知	sms_broadcast_room_status	radio	0:关闭
	// 1:开启	0		0	0	1	2020-09-08 15:53:42
	// 310	4	验证码时效配置(分钟)	sms_time	number		0		0	0	1	2020-09-08 15:53:42
	{SysConfigCategoryID: 5, ConfigName: "打印机自动打印", ConfigKey: "printing_auto_status", ConfigType: "radio", ConfigRule: "	0:关闭;1:开启", Required: 2, Info: "开启后订单支付成功后自动打印", Sort: 0, UserType: 1, Status: 1},

	{SysConfigCategoryID: 15, ConfigName: "支付宝支付状态", ConfigKey: "alipay_open", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "是否开启支付宝支付", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 15, ConfigName: "支付宝支付环境", ConfigKey: "alipay_env", ConfigType: "radio", ConfigRule: "0:沙箱;1:正式", Required: 2, Info: "支付宝支付环境，是否使用沙箱环境", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 15, ConfigName: "支付宝app_id", ConfigKey: "alipay_app_id", ConfigType: "input", ConfigRule: "", Required: 2, Info: "支付宝支付appid", Sort: 0, UserType: 2, Status: 1},
	// {SysConfigCategoryID: 15, ConfigName: "支付宝公钥", ConfigKey: "alipay_public_key", ConfigType: "input", ConfigRule: "", Required: 2, Info: "支付宝支付应用公钥", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 15, ConfigName: "支付密钥", ConfigKey: "alipay_private_key", ConfigType: "input", ConfigRule: "", Required: 2, Info: "支付宝支付应用密钥", Sort: 0, UserType: 2, Status: 1},

	{SysConfigCategoryID: 5, ConfigName: "打印机开启", ConfigKey: "printing_status", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "开启发票", ConfigKey: "mer_open_receipt", ConfigType: "radio", ConfigRule: "0:关闭;1:开启", Required: 2, Info: "设置是否开启发票", Sort: 0, UserType: 1, Status: 1},
}

func (m *config) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysConfig{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_configs 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configs).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_configs 表初始数据成功!")
		return nil
	})
}
