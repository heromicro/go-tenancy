package model

// FinancialRecord 商户财务流水
type FinancialRecord struct {
	RecordSn string `gorm:"column:record_sn;type:varchar(32);not null" json:"recordSn"` // 流水号
	OrderSn  string `gorm:"column:order_sn;type:varchar(32);not null" json:"orderSn"`   // 订单编号
	UserInfo string `gorm:"column:user_info;type:varchar(32);not null" json:"userInfo"` // 用户名

	FinancialType string  `gorm:"index:financial_type;column:financial_type;type:varchar(32);not null" json:"financialType"` // 流水类型
	FinancialPm   uint8   `gorm:"column:financial_pm;type:tinyint unsigned;not null;default:0" json:"financialPm"`           // 0 = 支出 1 = 获得
	Number        float64 `gorm:"column:number;type:decimal(8,2);not null;default:0.00" json:"number"`                       // 金额

	SysTenancyId uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	CUserId      uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
	OrderId      uint `gorm:"index:order_id;column:order_id;type:int unsigned;not null" json:"orderId"` // 订单id
}
