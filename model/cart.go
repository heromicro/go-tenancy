package model

import "github.com/snowlyg/go-tenancy/g"

// Cart 购物车表
type Cart struct {
	g.TENANCY_MODEL
	BaseCart
	PatientId    uint `json:"patientId" form:"patientId" gorm:"column:patient_id;comment:患者"`
	ProductId    uint `gorm:"index:product_id;column:product_id;type:int unsigned;not null" json:"productId"` // 商品ID
	CUserId      uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
	SysTenancyId uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}

type BaseCart struct {
	ProductType       int32  `gorm:"column:product_type;type:tinyint;not null;default:1" json:"productType" binding:"required"`                   // 类型 1=普通产品，2.预售商品
	ProductAttrUnique string `gorm:"column:product_attr_unique;type:varchar(16);not null;default:''" json:"productAttrUnique" binding:"required"` // 商品属性
	CartNum           int64  `gorm:"column:cart_num;type:smallint unsigned;not null;default:0" json:"cartNum" binding:"required"`                 // 商品数量
	Source            uint8  `gorm:"column:source;type:tinyint unsigned;not null;default:0" json:"source"`                                        // 来源 1.直播间,2.预售商品,3.助力商品
	SourceID          uint   `gorm:"column:source_id;type:int unsigned;not null;default:0" json:"sourceId"`                                       // 来源关联 id
	IsPay             int    `gorm:"column:is_pay;type:tinyint(1);not null;default:2" json:"isPay"`                                               // 2 = 未购买 1 = 已购买                                              // 是否删除
	IsNew             int    `gorm:"column:is_new;type:tinyint(1);not null;default:2" json:"isNew" binding:"required"`                            // 是否为立即购买
	IsFail            int    `gorm:"column:is_fail;type:tinyint unsigned;not null;default:2" json:"isFail"`                                       // 是否失效
}
