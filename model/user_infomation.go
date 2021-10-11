package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// UserVisit 商品浏览分析表
type UserVisit struct {
	g.TENANCY_MODEL

	Type    string `gorm:"index:type;column:type;type:varchar(32);not null" json:"type"`        // 记录类型 page,product
	TypeID  int    `gorm:"index:type;column:type_id;type:int;not null;default:0" json:"typeId"` // 商品ID
	Content string `gorm:"column:content;type:varchar(255)" json:"content"`                     // 备注描述

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
}

// UserBill 用户账单表
type UserBill struct {
	g.TENANCY_MODEL

	Pm       uint8   `gorm:"column:pm;type:tinyint unsigned;not null;default:0" json:"pm"`                   // 0 = 支出 1 = 获得
	Title    string  `gorm:"column:title;type:varchar(64);not null" json:"title"`                            // 账单标题
	Category string  `gorm:"index:type;column:category;type:varchar(64);not null" json:"category"`           // 明细种类
	Type     string  `gorm:"index:type;column:type;type:varchar(64);not null;default:''" json:"type"`        // 明细类型
	Number   float64 `gorm:"column:number;type:decimal(8,2) unsigned;not null;default:0.00" json:"number"`   // 明细数字
	Balance  float64 `gorm:"column:balance;type:decimal(8,2) unsigned;not null;default:0.00" json:"balance"` // 剩余
	Mark     string  `gorm:"column:mark;type:varchar(512);not null" json:"mark"`                             // 备注
	Status   int     `gorm:"column:status;type:tinyint(1);not null;default:2" json:"status"`                 // 1 = 待确定 2 = 有效 3 = 无效

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
	LinkID  uint `gorm:"index:type;column:link_id;type:varchar(32);not null;default:0" json:"linkId"` // 关联id
}

// UserExtract 用户提现表
type UserExtract struct {
	g.TENANCY_MODEL

	RealName     string    `gorm:"column:real_name;type:varchar(64)" json:"realName"`                                      // 姓名
	ExtractType  int       `gorm:"column:extract_type;type:tinyint(1);default:0" json:"extractType"`                       // 0 银行卡 1 支付宝 2微信
	BankCode     string    `gorm:"column:bank_code;type:varchar(32);default:0" json:"bankCode"`                            // 银行卡
	BankAddress  string    `gorm:"column:bank_address;type:varchar(256);default:''" json:"bankAddress"`                    // 开户地址
	AlipayCode   string    `gorm:"column:alipay_code;type:varchar(64);default:''" json:"alipayCode"`                       // 支付宝账号
	Wechat       string    `gorm:"column:wechat;type:varchar(15)" json:"wechat"`                                           // 微信号
	ExtractPic   string    `gorm:"column:extract_pic;type:varchar(128)" json:"extractPic"`                                 // 收款码
	ExtractPrice float64   `gorm:"column:extract_price;type:decimal(8,2) unsigned;default:0.00" json:"extractPrice"`       // 提现金额
	Balance      float64   `gorm:"column:balance;type:decimal(8,2) unsigned;default:0.00" json:"balance"`                  // 余额
	Mark         string    `gorm:"column:mark;type:varchar(512)" json:"mark"`                                              // 管理员备注
	AdminID      int       `gorm:"column:admin_id;type:int;default:0" json:"adminId"`                                      // 审核管理员
	FailMsg      string    `gorm:"column:fail_msg;type:varchar(128)" json:"failMsg"`                                       // 无效原因
	StatusTime   time.Time `gorm:"column:status_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"statusTime"` // 无效时间
	Status       int       `gorm:"column:status;type:tinyint;default:2" json:"status"`                                     // 1 审核中 2 已提现 3 未通过

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
}

// UserReceipt 发票
type UserReceipt struct {
	g.TENANCY_MODEL

	ReceiptType      int    `json:"receiptType" form:"receiptType" gorm:"type:tinyint(1);column:receipt_type;default:0;comment:发票类型：1.普通，2.增值"`
	ReceiptTitle     string `json:"receiptTitle" form:"receiptTitle" gorm:"type:varchar(128);column:receipt_title;comment:发票抬头"`
	ReceiptTitleType int    `json:"receiptTitleType" form:"receiptTitleType" gorm:"type:tinyint(1);column:receipt_title_type;comment:发票抬头类型：1.个人，2.企业"`
	DutyGaragraph    string `json:"dutyGaragraph" form:"dutyGaragraph" gorm:"column:duty_garagraph;comment:税号"`
	Email            string `json:"email" form:"email" gorm:"column:email;comment:邮箱"`
	BankName         string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:开户行"`
	BankCode         string `json:"bankCode" form:"bankCode" gorm:"column:bank_code;comment:银行账号"`
	Address          string `json:"address" form:"address" gorm:"column:address;comment:企业地址"`
	Tel              string `json:"tel" form:"tel" gorm:"column:tel;comment:企业电话"`
	IsDefault        int    `json:"isDefault" form:"isDefault" gorm:"type:tinyint(1);column:is_default;comment:是否默认"`

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
}

// UserRecharge 用户充值表
type UserRecharge struct {
	g.TENANCY_MODEL

	Price        float64   `gorm:"column:price;type:decimal(8,2) unsigned;not null;default:0.00" json:"price"`      // 充值金额
	GivePrice    float64   `gorm:"column:give_price;type:decimal(8,2);not null;default:0.00" json:"givePrice"`      // 购买赠送金额
	RechargeType string    `gorm:"column:recharge_type;type:varchar(32);not null" json:"rechargeType"`              // 充值类型
	Paid         int       `gorm:"column:paid;type:tinyint unsigned;not null;default:0" json:"paid"`                // 是否充值
	PayTime      time.Time `gorm:"column:pay_time;type:timestamp" json:"payTime"`                                   // 充值支付时间
	RefundPrice  float64   `gorm:"column:refund_price;type:decimal(10,2) unsigned;default:0.00" json:"refundPrice"` // 退款金额

	CUserId uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	OrderId string `gorm:"unique;column:order_id;type:varchar(32);not null" json:"orderId"` // 订单号
}

// UserRelation 用户记录表,关注店铺和商品
type UserRelation struct {
	g.TENANCY_MODEL

	Type   int  `gorm:"uniqueIndex:type_index;column:type;type:tinyint;not null;default:0" json:"type"`    // 关联类型(1= 普通商品、10 = 店铺、12=购买过)
	TypeID uint `gorm:"uniqueIndex:type_id_index;column:type_id;type:int unsigned;not null" json:"typeId"` // 类型的 id

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
}

// UserAddress 用户收货地址
type UserAddress struct {
	g.TENANCY_MODEL
	Name      string `json:"name" gorm:"type:varchar(32);not null;comment:姓名"`
	Phone     string `json:"phone" gorm:"type:varchar(16);not null;comment:手机号"`
	Sex       int    `json:"sex" form:"sex" gorm:"not null;column:sex;comment:性别 0:女 1:男，2：未知"`
	Country   string `json:"country" form:"country" gorm:"type:varchar(64);not null;column:country;comment:国家"`
	Province  string `json:"province" form:"province" gorm:"type:varchar(64);not null;column:province;comment:省份"`
	City      string `json:"city" form:"city" gorm:"type:varchar(64);not null;column:city;comment:城市"`
	District  string `json:"district" form:"district" gorm:"type:varchar(64);not null;column:district;comment:地区"`
	IsDefault int    `json:"isDefault" form:"isDefault" gorm:"not null;type:tinyint(1);column:is_default;comment:是否默认"`
	Detail    string `json:"detail" form:"detail" gorm:"type:varchar(254);not null;column:detail;comment:详细地址"`
	Postcode  string `json:"postcode" form:"postcode" gorm:"type:varchar(20);not null;column:postcode;comment:邮政编码"`

	// 可选字段
	Age          int    `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	HospitalName string `json:"hospitalName" form:"hospitalName" gorm:"type:varchar(50);column:hospital_name;comment:医院"`
	LocName      string `json:"locName" form:"locName" gorm:"type:varchar(50);column:loc_name;comment:科室名称"`
	BedNum       string `json:"bedNum" form:"bedNum" gorm:"type:varchar(10);column:bed_num;comment:床号"`
	HospitalNO   string `json:"hospitalNo" form:"hospitalNo" gorm:"type:varchar(20);column:hospital_no;comment:住院号"`
	Disease      string `json:"disease" form:"disease" gorm:"type:varchar(150);column:disease;comment:病种"`

	CUserId uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
}
