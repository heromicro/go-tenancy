package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// ShippingTemplate 运费表
type ShippingTemplate struct {
	g.TENANCY_MODEL
	BaseShippingTemplate

	SysTenancyId uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`
}

type BaseShippingTemplate struct {
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`                           // 模板名称
	Type       int    `gorm:"column:type;type:tinyint unsigned;not null;default:2" json:"type"`             // 计费方式 1=数量 2=重量 3=体积
	Appoint    int    `gorm:"column:appoint;type:tinyint unsigned;not null;default:2" json:"appoint"`       // 开启指定包邮
	Undelivery int    `gorm:"column:undelivery;type:tinyint unsigned;not null;default:2" json:"undelivery"` // 开启指定区域不配送
	IsDefault  int    `gorm:"column:is_default;type:tinyint;default:2" json:"isDefault"`                    // 默认模板
	Sort       int    `gorm:"index:st_sort;column:sort;type:int;not null;default:0" json:"sort"`            // 排序
}

// ShippingTemplateUndelivery 指定不配送区域表
type ShippingTemplateUndelivery struct {
	g.TENANCY_MODEL
	Code int `json:"code" gorm:""`

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}

// ShippingTemplateRegion 配送区域表
type ShippingTemplateRegion struct {
	g.TENANCY_MODEL
	First         float64 `gorm:"column:first;type:decimal(10,2) unsigned;not null;default:0.00" json:"first"`                  // 首件
	FirstPrice    float64 `gorm:"column:first_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"firstPrice"`       // 首件运费
	Continue      float64 `gorm:"column:continue;type:decimal(10,2) unsigned;not null;default:0.00" json:"continue"`            // 续件
	ContinuePrice float64 `gorm:"column:continue_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"continuePrice"` // 续件运费
	Code          int     `json:"code" gorm:""`

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}

// ShippingTemplateFree 指定包邮信息表
type ShippingTemplateFree struct {
	g.TENANCY_MODEL
	Number uint    `gorm:"column:number;type:int unsigned;not null;default:0" json:"number"`            // 包邮件数
	Price  float64 `gorm:"column:price;type:decimal(10,2) unsigned;not null;default:0.00" json:"price"` // 包邮金额
	Code   int     `json:"code" gorm:""`                                                                // 城市ID

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}
