package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 32}, Pid: 38, Path: "/110/38/", Icon: "", MenuName: "api管理", Route: "/admin/setting/api", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 33}, Pid: 0, Path: "/", Icon: "s-home", MenuName: "仪表盘", Route: "/admin/dashboard", Params: "", Sort: 100, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 34}, Pid: 110, Path: "/110/", Icon: "", MenuName: "系统配置", Route: "/admin/config", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 35}, Pid: 34, Path: "/110/34/", Icon: "", MenuName: "配置分类", Route: "/admin/config/classify", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 36}, Pid: 34, Path: "/110/34/", Icon: "", MenuName: "配置管理", Route: "/admin/config/setting", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 38}, Pid: 110, Path: "/110/", Icon: "", MenuName: "权限管理", Route: "/admin/setting", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 39}, Pid: 38, Path: "/110/38/", Icon: "", MenuName: "角色管理", Route: "/admin/setting/systemRole", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 40}, Pid: 38, Path: "/110/38/", Icon: "", MenuName: "员工管理", Route: "/admin/setting/systemAdmin", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 41}, Pid: 34, Path: "/110/34/", Icon: "", MenuName: "素材管理", Route: "/admin/config/picture", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 42}, Pid: 0, Path: "/", Icon: "s-shop", MenuName: "商户", Route: "/admin/merchant", Params: "", Sort: 94, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 43}, Pid: 42, Path: "/42/", Icon: "", MenuName: "商户菜单管理", Route: "/admin/merchant/system", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 44}, Pid: 42, Path: "/42/", Icon: "", MenuName: "商户列表", Route: "/admin/merchant/list", Params: "", Sort: 9, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 47}, Pid: 38, Path: "/110/38/", Icon: "", MenuName: "操作日志", Route: "/admin/setting/systemLog", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 48}, Pid: 38, Path: "/110/38/", Icon: "", MenuName: "菜单管理", Route: "/admin/setting/menu", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 49}, Pid: 526, Path: "/526/", Icon: "", MenuName: "权限管理", Route: "/merchant/setting", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 50}, Pid: 49, Path: "/526/49/", Icon: "", MenuName: "角色管理", Route: "/merchant/setting/systemRole", Params: "", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 51}, Pid: 49, Path: "/526/49/", Icon: "", MenuName: "员工管理", Route: "/merchant/setting/systemAdmin", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 52}, Pid: 49, Path: "/526/49/", Icon: "", MenuName: "操作日志", Route: "/merchant/setting/systemLog", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 54}, Pid: 526, Path: "/526/", Icon: "", MenuName: "素材管理", Route: "/merchant/config/picture", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 55}, Pid: 0, Path: "/", Icon: "s-home", MenuName: "首页", Route: "/merchant/dashboard", Params: "", Sort: 100, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 57}, Pid: 521, Path: "/520/521/", Icon: "", MenuName: "组合数据", Route: "/admin/group/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 58}, Pid: 519, Path: "/519/", Icon: "", MenuName: "公众号", Route: "/admin/app/wechat", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 59}, Pid: 58, Path: "/519/58/", Icon: "", MenuName: "微信菜单", Route: "/admin/app/wechat/menus", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 60}, Pid: 0, Path: "/", Icon: "s-management", MenuName: "文章", Route: "/cms", Params: "", Sort: 96, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 61}, Pid: 60, Path: "/60/", Icon: "", MenuName: "文章管理", Route: "/admin/cms/article", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 62}, Pid: 60, Path: "/60/", Icon: "", MenuName: "文章分类", Route: "/admin/cms/articleCategory", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 63}, Pid: 110, Path: "/110/", Icon: "", MenuName: "商城配置", Route: "/admin/system", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 64}, Pid: 63, Path: "/110/63/", Icon: "", MenuName: "基础配置", Route: "/admin/systemForm/Basics/base", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 68}, Pid: 63, Path: "/110/63/", Icon: "", MenuName: "商城配置", Route: "/admin/systemForm/Basics/shop", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 72}, Pid: 111, Path: "/110/111/", Icon: "", MenuName: "短信配置", Route: "/admin/systemForm/Basics/message", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 74}, Pid: 526, Path: "/526/", Icon: "", MenuName: "店铺配置", Route: "/merchant/systemForm/Basics/mer_base", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 77}, Pid: 58, Path: "/519/58/", Icon: "", MenuName: "自动回复", Route: "/admin/app/wechat/reply", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 79}, Pid: 77, Path: "/519/58/77/", Icon: "", MenuName: "微信关注回复", Route: "/admin/app/wechat/reply/follow/subscribe", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 80}, Pid: 77, Path: "/519/58/77/", Icon: "", MenuName: "关键字回复", Route: "/admin/app/wechat/reply/keyword", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 81}, Pid: 77, Path: "/519/58/77/", Icon: "", MenuName: "无效关键词回复", Route: "/admin/app/wechat/reply/index/default", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 82}, Pid: 58, Path: "/519/58/", Icon: "", MenuName: "图文管理", Route: "/admin/app/wechat/newsCategory", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 87}, Pid: 0, Path: "/", Icon: "s-goods", MenuName: "商品", Route: "/product", Params: "", Sort: 99, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 88}, Pid: 87, Path: "/87/", Icon: "", MenuName: "商品分类", Route: "/admin/product/classify", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 92}, Pid: 87, Path: "/87/", Icon: "", MenuName: "品牌管理", Route: "/admin/product/brand", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 93}, Pid: 92, Path: "/87/92/", Icon: "", MenuName: "品牌分类", Route: "/admin/product/band/brandClassify", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 94}, Pid: 92, Path: "/87/92/", Icon: "", MenuName: "品牌列表", Route: "/admin/product/band/brandList", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 95}, Pid: 0, Path: "/", Icon: "s-goods", MenuName: "商品", Route: "/merchant/product", Params: "", Sort: 99, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 96}, Pid: 95, Path: "/95/", Icon: "", MenuName: "商品分类", Route: "/merchant/product/classify", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 99}, Pid: 95, Path: "/95/", Icon: "", MenuName: "商品规格", Route: "/merchant/product/attr", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 100}, Pid: 42, Path: "/42/", Icon: "", MenuName: "商户分类", Route: "/merchant/classify", Params: "", Sort: 10, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 101}, Pid: 0, Path: "/", Icon: "user-solid", MenuName: "用户", Route: "/admin/user", Params: "", Sort: 96, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 102}, Pid: 101, Path: "/101/", Icon: "", MenuName: "用户分组", Route: "/admin/user/group", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 103}, Pid: 101, Path: "/101/", Icon: "", MenuName: "用户列表", Route: "/admin/user/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 104}, Pid: 101, Path: "/101/", Icon: "", MenuName: "用户标签", Route: "/admin/user/label", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 105}, Pid: 95, Path: "/95/", Icon: "", MenuName: "商品列表", Route: "/merchant/product/list", Params: "", Sort: 2, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 106}, Pid: 0, Path: "/", Icon: "orange", MenuName: "营销", Route: "/merchant/marketing", Params: "", Sort: 97, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 107}, Pid: 106, Path: "/106/", Icon: "", MenuName: "优惠券", Route: "/merchant/marketing/coupon", Params: "", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 110}, Pid: 0, Path: "/", Icon: "s-tools", MenuName: "设置", Route: "/merchant/settings", Params: "", Sort: 92, Hidden: 2, IsTenancy: 2, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 111}, Pid: 110, Path: "/110/", Icon: "", MenuName: "短信设置", Route: "/sms", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 112}, Pid: 111, Path: "/110/111/", Icon: "", MenuName: "短信账户", Route: "/admin/sms/config", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 113}, Pid: 111, Path: "/110/111/", Icon: "", MenuName: "短信模板", Route: "/admin/sms/template", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 114}, Pid: 111, Path: "/110/111/", Icon: "", MenuName: "短信购买", Route: "/admin/sms/pay", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 115}, Pid: 107, Path: "/106/107/", Icon: "", MenuName: "优惠券列表", Route: "/merchant/marketing/coupon/list", Params: "", Sort: 1, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 116}, Pid: 520, Path: "/520/", Icon: "", MenuName: "安全维护", Route: "/admin/maintain", Params: "", Sort: 9, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 117}, Pid: 116, Path: "/520/116/", Icon: "", MenuName: "数据备份", Route: "/maintain/dataBackup", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 118}, Pid: 110, Path: "/110/", Icon: "", MenuName: "物流管理", Route: "/admin/freight", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 119}, Pid: 118, Path: "/110/118/", Icon: "", MenuName: "物流公司", Route: "/admin/freight/express", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 120}, Pid: 526, Path: "/526/", Icon: "", MenuName: "客服管理", Route: "/merchant/config/service", Params: "", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 121}, Pid: 87, Path: "/87/", Icon: "", MenuName: "评论管理", Route: "/admin/product/comment", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 122}, Pid: 107, Path: "/106/107/", Icon: "", MenuName: "会员领取记录", Route: "/merchant/marketing/coupon/user", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 466}, Pid: 101, Path: "/101/", Icon: "", MenuName: "用户反馈", Route: "/admin/feedback", Params: "", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 467}, Pid: 466, Path: "/101/466/", Icon: "", MenuName: "反馈分类", Route: "/admin/feedback/classify", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 468}, Pid: 466, Path: "/101/466/", Icon: "", MenuName: "反馈列表", Route: "/admin/feedback/list", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 512}, Pid: 0, Path: "/", Icon: "s-order", MenuName: "订单", Route: "/merchant/order", Params: "", Sort: 99, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 513}, Pid: 512, Path: "/512/", Icon: "", MenuName: "订单管理", Route: "/merchant/order/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 514}, Pid: 0, Path: "/", Icon: "s-flag", MenuName: "分销", Route: "/admin/promoter", Params: "", Sort: 97, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 515}, Pid: 0, Path: "/", Icon: "s-data", MenuName: "财务", Route: "/admin/accounts", Params: "", Sort: 95, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 516}, Pid: 515, Path: "/515/", Icon: "", MenuName: "提现管理", Route: "/admin/accounts/extract", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 517}, Pid: 537, Path: "/515/537/", Icon: "", MenuName: "充值记录", Route: "/admin/accounts/bill", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 518}, Pid: 515, Path: "/515/", Icon: "", MenuName: "财务对账", Route: "/admin/accounts/reconciliation", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 519}, Pid: 0, Path: "/", Icon: "s-grid", MenuName: "应用", Route: "/admin/apploction", Params: "", Sort: 93, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 520}, Pid: 0, Path: "/", Icon: "s-help", MenuName: "维护", Route: "/admin/safe", Params: "", Sort: 91, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 521}, Pid: 520, Path: "/520/", Icon: "", MenuName: "开发配置", Route: "/admin/safe/exploit", Params: "[]", Sort: 10, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 522}, Pid: 514, Path: "/514/", Icon: "", MenuName: "分销员列表", Route: "/admin/promoter/user", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 523}, Pid: 514, Path: "/514/", Icon: "", MenuName: "分销配置", Route: "/admin/promoter/config", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 524}, Pid: 519, Path: "/519/", Icon: "", MenuName: "小程序", Route: "/admin/app/routine", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 525}, Pid: 0, Path: "/", Icon: "s-data", MenuName: "财务", Route: "/merchant/accounts", Params: "", Sort: 97, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 526}, Pid: 0, Path: "/", Icon: "s-tools", MenuName: "设置", Route: "/merchant/config", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 528}, Pid: 512, Path: "/512/", Icon: "", MenuName: "退款单", Route: "/merchant/order/refund", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 532}, Pid: 58, Path: "/519/58/", Icon: "", MenuName: "微信模板消息", Route: "/admin/app/wechat/template", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 536}, Pid: 42, Path: "/42/", Icon: "", MenuName: "商户对账", Route: "/merchant/list/record/:id?", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 537}, Pid: 515, Path: "/515/", Icon: "", MenuName: "财务记录", Route: "/admin/accounts/record", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 538}, Pid: 537, Path: "/515/537/", Icon: "", MenuName: "资金记录", Route: "/admin/accounts/capital", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 539}, Pid: 87, Path: "/87/", Icon: "", MenuName: "商品管理", Route: "/admin/product/examine", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 540}, Pid: 0, Path: "/", Icon: "s-cooperation", MenuName: "订单管理", Route: "/admin/order", Params: "", Sort: 98, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 541}, Pid: 540, Path: "/540/", Icon: "", MenuName: "订单列表", Route: "/admin/order/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 542}, Pid: 540, Path: "/540/", Icon: "", MenuName: "退款单", Route: "/admin/order/refund", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 543}, Pid: 525, Path: "/525/", Icon: "", MenuName: "财务对账", Route: "/merchant/accounts/reconciliation", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 544}, Pid: 95, Path: "/95/", Icon: "", MenuName: "商品评价", Route: "/merchant/product/reviews", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 545}, Pid: 524, Path: "/519/524/", Icon: "", MenuName: "小程序订阅消息", Route: "/app/routine/template", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 546}, Pid: 526, Path: "/526/", Icon: "", MenuName: " 基础配置", Route: "/merchant/systemForm/modifyStoreInfo", Params: "", Sort: 1, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 666}, Pid: 116, Path: "/520/116/", Icon: "", MenuName: "商业授权", Route: "/admin/maintain/auth", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 667}, Pid: 63, Path: "/110/63/", Icon: "", MenuName: "余额/充值设置", Route: "/admin/systemForm/Basics/balance", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 668}, Pid: 63, Path: "/110/63/", Icon: "", MenuName: "文件上传", Route: "/admin/systemForm/Basics/upload", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 669}, Pid: 110, Path: "/110/", Icon: "", MenuName: "支付配置", Route: "/admin/pay_config", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 670}, Pid: 110, Path: "/110/", Icon: "", MenuName: "应用配置", Route: "/admin/app_config", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 671}, Pid: 669, Path: "/110/669/", Icon: "", MenuName: "公众号支付配置", Route: "/admin/systemForm/Basics/wechat_payment", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 672}, Pid: 669, Path: "/110/669/", Icon: "", MenuName: "小程序支付配置", Route: "/admin/systemForm/Basics/routine_pay", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 673}, Pid: 670, Path: "/110/670/", Icon: "", MenuName: "公众号配置", Route: "/admin/systemForm/Basics/wechat", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 674}, Pid: 670, Path: "/110/670/", Icon: "", MenuName: "小程序配置", Route: "/admin/systemForm/Basics/smallapp", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 676}, Pid: 514, Path: "/514/", Icon: "", MenuName: "相关配置", Route: "/admin/systemForm/Basics/brokerage", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 677}, Pid: 514, Path: "/514/", Icon: "", MenuName: "提现银行管理", Route: "/admin/promoter/bank/76", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 678}, Pid: 110, Path: "/110/", Icon: "", MenuName: "页面设置", Route: "/admin/page_setting", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 679}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "首页幻灯片", Route: "/admin/promoter/bank/72", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 680}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "首页导航按钮", Route: "/admin/promoter/bank/73", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 681}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "首页推荐区", Route: "/admin/promoter/bank/74", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 682}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "个人中心幻灯片", Route: "/admin/promoter/bank/71", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 683}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "个人中心菜单", Route: "/admin/promoter/bank/70", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 684}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "热门搜索", Route: "/admin/promoter/bank/67", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 685}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "分销特权", Route: "/admin/promoter/bank/75", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 686}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "分销海报", Route: "/admin/promoter/bank/68", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 687}, Pid: 678, Path: "/110/678/", Icon: "", MenuName: "充值金额配置", Route: "//adminpromoter/bank/69", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 688}, Pid: 63, Path: "/110/63/", Icon: "", MenuName: "登录页幻灯片", Route: "/admin/promoter/bank/77", Params: "", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 689}, Pid: 110, Path: "/110/", Icon: "", MenuName: "测试页面", Route: "/test", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 690}, Pid: 689, Path: "/110/689/", Icon: "", MenuName: "功能测试", Route: "/admin/test/pay", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 691}, Pid: 110, Path: "/110/", Icon: "", MenuName: "MQTT管理", Route: "/mqtt", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 692}, Pid: 691, Path: "/110/691/", Icon: "", MenuName: "MQTT客户端", Route: "/admin/mqtt/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 693}, Pid: 691, Path: "/110/691/", Icon: "", MenuName: "MQTT消息日志", Route: "/admin/mqtt/record", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 694}, Pid: 110, Path: "/110/", Icon: "", MenuName: "定时任务管理", Route: "/job", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 695}, Pid: 694, Path: "/110/694/", Icon: "", MenuName: "定时任务", Route: "/admin/job/list", Params: "", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 700}, Pid: 526, Path: "/526/", Icon: "", MenuName: "运费模板", Route: "/merchant/config/freight/shippingTemplates", Params: "", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 719}, Pid: 0, Path: "/", Icon: "help", MenuName: "营销", Route: "/admin/marketing", Params: "[]", Sort: 97, Hidden: 1, IsTenancy: 0, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 720}, Pid: 719, Path: "/719/", Icon: "", MenuName: "优惠券", Route: "/admin/marketing/coupon", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 721}, Pid: 720, Path: "/719/720/", Icon: "", MenuName: "优惠券列表", Route: "/admin/marketing/coupon/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 731}, Pid: 514, Path: "/514/", Icon: "", MenuName: "分销礼包", Route: "/admin/promoter/gift", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 734}, Pid: 720, Path: "/719/720/", Icon: "", MenuName: "会员领取记录", Route: "/admin/marketing/coupon/user", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 778}, Pid: 42, Path: "/42/", Icon: "", MenuName: "商户入驻申请", Route: "/merchant/application", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 779}, Pid: 780, Path: "/719/780/", Icon: "", MenuName: "秒杀配置", Route: "/admin/marketing/seckill/seckillConfig", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 780}, Pid: 719, Path: "/719/", Icon: "", MenuName: "秒杀", Route: "/admin/marketing/seckill", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 781}, Pid: 782, Path: "/719/782/", Icon: "", MenuName: "直播间管理", Route: "/admin/marketing/studio/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 782}, Pid: 719, Path: "/719/", Icon: "", MenuName: "直播", Route: "/admin/marketing2", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 783}, Pid: 782, Path: "/719/782/", Icon: "", MenuName: "直播商品管理", Route: "/admin/marketing/broadcast/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 784}, Pid: 540, Path: "/540/", Icon: "", MenuName: "核销订单", Route: "/admin/order/cancellation", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 785}, Pid: 106, Path: "/106/", Icon: "", MenuName: "直播", Route: "/", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 786}, Pid: 785, Path: "/106/785/", Icon: "", MenuName: "直播间管理", Route: "/merchant/marketing/studio/list", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 787}, Pid: 785, Path: "/106/785/", Icon: "", MenuName: "直播商品管理", Route: "/merchant/marketing/broadcast/list", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 788}, Pid: 106, Path: "/106/", Icon: "", MenuName: "秒杀", Route: "/merchant/marketing/seckill/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 789}, Pid: 512, Path: "/512/", Icon: "", MenuName: "核销订单", Route: "/merchant/order/cancellation", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 790}, Pid: 525, Path: "/525/", Icon: "", MenuName: "资金流水", Route: "/merchant/accounts/capitalFlow", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 791}, Pid: 515, Path: "/515/", Icon: "", MenuName: "资金流水", Route: "/admin/accounts/capitalFlow", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 793}, Pid: 526, Path: "/526/", Icon: "", MenuName: "自提点设置", Route: "/merchant/systemForm/systemStore", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 794}, Pid: 780, Path: "/719/780/", Icon: "", MenuName: "秒杀管理", Route: "/admin/marketing/seckill/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1022}, Pid: 719, Path: "/719/", Icon: "", MenuName: "预售", Route: "/admin/marketing/presell", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1023}, Pid: 1022, Path: "/719/1022/", Icon: "", MenuName: "预售商品", Route: "/admin/marketing/presell/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1024}, Pid: 1022, Path: "/719/1022/", Icon: "", MenuName: "预售协议", Route: "/admin/marketing/presell/agreement", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1025}, Pid: 106, Path: "106/", Icon: "", MenuName: "预售", Route: "/merchant/marketing/presell/list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1026}, Pid: 669, Path: "/110/669/", Icon: "", MenuName: "支付宝支付配置", Route: "/admin/systemForm/Basics/alipay", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1027}, Pid: 0, Path: "/", Icon: "user-solid", MenuName: "用户", Route: "/merchant/user", Params: "[]", Sort: 1, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1028}, Pid: 1027, Path: "1027/", Icon: "", MenuName: "标签管理", Route: "/merchant/user/_label", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1029}, Pid: 1028, Path: "1028/", Icon: "", MenuName: "手动标签", Route: "/merchant/user/label", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1030}, Pid: 1028, Path: "1028/", Icon: "", MenuName: "自动标签", Route: "/merchant/user/maticlabel", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1095}, Pid: 1051, Path: "/719/1051/", Icon: "", MenuName: "活动商品", Route: "/marketing/assist/goods_list", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 2, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1096}, Pid: 1051, Path: "/719/1051/", Icon: "", MenuName: "助力活动", Route: "/marketing/assist/list", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1099}, Pid: 106, Path: "106/", Icon: "", MenuName: "助力", Route: "/merchant/assist", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1100}, Pid: 1099, Path: "1099/", Icon: "", MenuName: "助力商品", Route: "/merchant/marketing/assist/list", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1101}, Pid: 1099, Path: "1099/", Icon: "", MenuName: "助力活动", Route: "/merchant/marketing/assist/assist_set", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1102}, Pid: 525, Path: "525/", Icon: "", MenuName: "发票管理", Route: "/merchant/order/invoice", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1103}, Pid: 1027, Path: "1027/", Icon: "", MenuName: "用户管理", Route: "/merchant/user/list", Params: "[]", Sort: 0, Hidden: 2, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1119}, Pid: 0, Path: "/", Icon: "message", MenuName: "公告列表", Route: "/merchant/station/notice", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1120}, Pid: 110, Path: "/110/", Icon: "", MenuName: "公告管理", Route: "/merchant/station/notice", Params: "[]", Sort: 0, Hidden: 1, IsTenancy: 1, IsMenu: 1},
}

//Init sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
