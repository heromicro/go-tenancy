package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// UserMerchant 商户用户表
type UserMerchant struct {
	g.TENANCY_MODEL

	FirstPayTime time.Time `gorm:"column:first_pay_time;type:timestamp" json:"firstPayTime"`                           // 首次消费时间
	LastPayTime  time.Time `gorm:"column:last_pay_time;type:timestamp" json:"lastPayTime"`                             // 最后一次消费时间
	PayCount     int       `gorm:"column:pay_count;type:int unsigned;not null;default:0" json:"payCount"`              // 消费次数
	PayPrice     float64   `gorm:"column:pay_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"payPrice"` // 消费金额
	LastTime     time.Time `gorm:"column:last_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"lastTime"` // 最后一次访问时间
	Status       int       `gorm:"column:status;type:tinyint unsigned;default:1" json:"status"`                        // 状态

	CUserId      uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
	SysTenancyId uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
