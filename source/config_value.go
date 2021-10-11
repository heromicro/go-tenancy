package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var SysConfigValue = new(configValue)

type configValue struct{}

var configValues = []model.SysConfigValue{
	{ConfigKey: "site_url", Value: "http://tenancy.t.chindeo.com", SysTenancyId: 0},
	{ConfigKey: "site_name", Value: "GOTENANCY\u591a\u5546\u6237\u5546\u57ce", SysTenancyId: 0},
	{ConfigKey: "site_open", Value: "1", SysTenancyId: 0},
	{ConfigKey: "wechat_name", Value: "GOTENANCY \u591a\u5546\u6237", SysTenancyId: 0},
	{ConfigKey: "set_phone", Value: "18741523695", SysTenancyId: 55},
	{ConfigKey: "set_email", Value: "mkpmkmp", SysTenancyId: 55},
	{ConfigKey: "tenancy_admin_password", Value: "123456", SysTenancyId: 0},
	{ConfigKey: "wechat_qrcode", Value: "", SysTenancyId: 0},
	{ConfigKey: "wechat_avatar", Value: "", SysTenancyId: 0},
	{ConfigKey: "wechat_share_img", Value: "", SysTenancyId: 0},
	{ConfigKey: "wechat_share_title", Value: "", SysTenancyId: 0},
	{ConfigKey: "wechat_share_synopsis", Value: "", SysTenancyId: 0},
	{ConfigKey: "wechat_encode", Value: "0", SysTenancyId: 0},
	{ConfigKey: "upload_type", Value: "1", SysTenancyId: 0},
	{ConfigKey: "mer_store_stock", Value: "10", SysTenancyId: 55},
	{ConfigKey: "sms_user_pay_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_user_postage_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_user_take_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_order_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_pay_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_refund_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_take_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_user_change_order_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "recharge_attention", Value: "\u5145\u503c\u540e\u5e10\u6237\u7684\u91d1\u989d\u4e0d\u80fd\u63d0\u73b0\uff0c\u53ef\u7528\u4e8e\u5546\u57ce\u6d88\u8d39\u4f7f\u7528\n\u4f63\u91d1\u5bfc\u5165\u8d26\u6237\u4e4b\u540e\u4e0d\u80fd\u518d\u6b21\u5bfc\u51fa\u3001\u4e0d\u53ef\u63d0\u73b0\n\u8d26\u6237\u5145\u503c\u51fa\u73b0\u95ee\u9898\u53ef\u8054\u7cfb\u5546\u57ce\u5ba2\u670d\uff0c\u4e5f\u53ef\u62e8\u6253\u5546\u57ce\u5ba2\u670d\u70ed\u7ebf\uff1a4008888888", SysTenancyId: 0},
	{ConfigKey: "auto_close_order_timer", Value: "15", SysTenancyId: 0},
	{ConfigKey: "auto_take_order_timer", Value: "15", SysTenancyId: 0},
	{ConfigKey: "refund_message", Value: "\u6536\u8d27\u5730\u5740\u586b\u9519\u4e86;\u4e0e\u63cf\u8ff0\u4e0d\u7b26;\u4fe1\u606f\u586b\u9519\u4e86;\u91cd\u65b0\u62cd;\u6536\u5230\u5546\u54c1\u635f\u574f\u4e86;\u672a\u6309\u9884\u5b9a\u65f6\u95f4\u53d1\u8d27;\u5176\u5b83\u539f\u56e0", SysTenancyId: 0},
	{ConfigKey: "mer_refund_order_agree", Value: "7", SysTenancyId: 0},
	{ConfigKey: "mer_refund_address", Value: "", SysTenancyId: 56},
	{ConfigKey: "mer_refund_user", Value: "", SysTenancyId: 55},
	{ConfigKey: "bank", Value: "", SysTenancyId: 55},
	{ConfigKey: "bank_name", Value: "", SysTenancyId: 55},
	{ConfigKey: "bank_number", Value: "32342354353", SysTenancyId: 55},
	{ConfigKey: "bank_address", Value: "", SysTenancyId: 55},
	{ConfigKey: "user_extract_min", Value: "100", SysTenancyId: 0},
	{ConfigKey: "lock_brokerage_timer", Value: "0", SysTenancyId: 0},
	{ConfigKey: "recharge_switch", Value: "0", SysTenancyId: 0},
	{ConfigKey: "store_user_min_recharge", Value: "100", SysTenancyId: 0},
	{ConfigKey: "balance_func_status", Value: "0", SysTenancyId: 0},
	{ConfigKey: "yue_pay_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "home_ad_pic", Value: "", SysTenancyId: 0},
	{ConfigKey: "home_ad_url", Value: "/pages/users/user_coupon/index", SysTenancyId: 0},
	{ConfigKey: "promoter_ explain", Value: "\u963f\u8d3e\u514b\u65af", SysTenancyId: 0},
	{ConfigKey: "promoter_bag_number", Value: "2", SysTenancyId: 0},
	{ConfigKey: "promoter_explain", Value: "\u5145\u503c\u540e\u5e10\u6237\u7684\u91d1\u989d\u4e0d\u80fd\u63d0\u73b0\uff0c\u53ef\u7528\u4e8e\u5546\u57ce\u6d88\u8d39\u4f7f\u7528\n\u4f63\u91d1\u5bfc\u5165\u8d26\u6237\u4e4b\u540e\u4e0d\u80fd\u518d\u6b21\u5bfc\u51fa\u3001\u4e0d\u53ef\u63d0\u73b0\n\u8d26\u6237\u5145\u503c\u51fa\u73b0\u95ee\u9898\u53ef\u8054\u7cfb\u5546\u57ce\u5ba2\u670d\uff0c\u4e5f\u53ef\u62e8\u6253\u5546\u57ce\u5ba2\u670d\u70ed\u7ebf\uff1a4008888888", SysTenancyId: 0},
	{ConfigKey: "max_bag_number", Value: "20", SysTenancyId: 0},
	{ConfigKey: "site_logo", Value: "", SysTenancyId: 0},
	{ConfigKey: "share_info", Value: "", SysTenancyId: 0},
	{ConfigKey: "share_pic", Value: "", SysTenancyId: 0},
	{ConfigKey: "share_title", Value: "", SysTenancyId: 0},
	{ConfigKey: "sms_fahuo_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_take_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_pay_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_revision_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_pay_false_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_refund_fail_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_refund_success_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_refund_confirm_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_return_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_admin_postage_status", Value: "1", SysTenancyId: 0},
	{ConfigKey: "sms_account", Value: "", SysTenancyId: 0},
	{ConfigKey: "sms_token", Value: "", SysTenancyId: 0},
	{ConfigKey: "sys_login_logo", Value: "", SysTenancyId: 0},
	{ConfigKey: "set_phone", Value: "15109234132", SysTenancyId: 65},
	{ConfigKey: "set_email", Value: "78532941@qq.com", SysTenancyId: 65},
	{ConfigKey: "mer_store_stock", Value: "5", SysTenancyId: 65},
	{ConfigKey: "mer_refund_address", Value: "\u9655\u897f\u7701\u897f\u5b89\u5e02\u5317\u5927\u885775\u53f7", SysTenancyId: 65},
	{ConfigKey: "mer_refund_user", Value: "\u90d1\u6b63", SysTenancyId: 65},
	{ConfigKey: "bank", Value: "\u4e2d\u56fd\u519c\u4e1a\u94f6\u884c", SysTenancyId: 65},
	{ConfigKey: "bank_number", Value: "4214512365015841214", SysTenancyId: 65},
	{ConfigKey: "bank_name", Value: "\u90d1\u8def", SysTenancyId: 65},
	{ConfigKey: "bank_address", Value: "\u5317\u5927\u8857\u652f\u884c", SysTenancyId: 65},
	{ConfigKey: "sys_menu_logo", Value: "", SysTenancyId: 0},
	{ConfigKey: "sys_menu_slogo", Value: "", SysTenancyId: 0},
	{ConfigKey: "sys_login_title", Value: "", SysTenancyId: 0},
	{ConfigKey: "express_app_code", Value: "", SysTenancyId: 0},
	{ConfigKey: "sys_intention_agree", Value: "", SysTenancyId: 0},
	{ConfigKey: "mer_intention_open", Value: "", SysTenancyId: 0},
	{ConfigKey: "sms_time", Value: "30", SysTenancyId: 0},
}

func (m *configValue) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysConfigValue{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> config_values 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configValues).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> config_values 表初始数据成功!")
		return nil
	})
}
