package g

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type TENANCY_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

var node *snowflake.Node

// CreateOrderSn 生成订单号 orderType+20060102150405+随机数
// 订单组 G+20060102150405+随机数
// 退款单 R+20060102150405+随机数
// 资金流水 RC+20060102150405+随机数
func CreateOrderSn(orderType interface{}) string {
	id := fmt.Sprintf("%d", getNodeId().Generate().Int64())
	now := time.Now().Format("20060102150405")
	return fmt.Sprintf("%v%s%s", orderType, now, id[2:])
}

func getNodeId() *snowflake.Node {
	if node != nil {
		return node
	}
	node, _ = snowflake.NewNode(1)
	return node
}

// IsInit 项目是否需要初始化
func IsInit() bool {
	if TENANCY_DB == nil {
		return false
	} else if TENANCY_CONFIG.System.CacheType == "redis" && TENANCY_CACHE == nil {
		return false
	}
	return true
}
