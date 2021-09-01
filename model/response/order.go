package response

import (
	"github.com/shopspring/decimal"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/datatypes"
)

type OrderList struct {
	TenancyResponse
	model.BaseOrder
	GroupOrderSn     string         `json:"groupOrderSn" form:"groupOrderSn"`
	TenancyName      string         `json:"tenancyName" form:"tenancyName"`
	IsTrader         int            `json:"isTrader" form:"isTrader"`
	SysUserID        uint           `json:"sysUserId" form:"sysUserId"`
	SysTenancyID     uint           `json:"sysTenancyId"`
	GroupOrderID     int            `json:"groupOrderId"`
	ReconciliationID uint8          `json:"reconciliationId"`
	OrderProduct     []OrderProduct `gorm:"-" json:"orderProduct"`
}

type OrderProduct struct {
	ID       uint           `json:"id"`
	CartInfo datatypes.JSON `json:"cartInfo"`
	model.BaseOrderProduct
	OrderID   uint `json:"orderID"`
	ProductID uint `json:"productId"` // 商品ID
}

type OrderCondition struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Conditions map[string]interface{} `json:"conditions"`
	IsDeleted  bool                   `json:"is_deleted"`
}

type OrderDetail struct {
	TenancyResponse
	model.BaseOrder
	SysUserID        uint           `json:"sysUserId" form:"sysUserId"`
	SysTenancyID     uint           `json:"sysTenancyId"`
	GroupOrderID     int            `json:"groupOrderId"`
	ReconciliationID uint8          `json:"reconciliationId"`
	UserNickName     string         `json:"userNickName" form:"userNickName"`
	OrderProduct     []OrderProduct `gorm:"-" json:"orderProduct"`
}

type CheckOrder struct {
	CartList
	CheckOrderDetail
}

type CheckOrderDetail struct {
	TotalPrice    decimal.Decimal                     `json:"totalPrice" form:"totalPrice"`       // 总计价格
	TotalOtPrice  decimal.Decimal                     `json:"totalOtPrice" form:"totalOtPrice"`   // 总计原价
	OrderPrice    decimal.Decimal                     `json:"orderPrice" form:"orderPrice"`       // 订单价格
	OrderOtPrice  decimal.Decimal                     `json:"orderOtPrice" form:"orderOtPrice"`   // 订单原价
	PostagePrice  decimal.Decimal                     `json:"postagePrice" form:"postagePrice"`   // 油费
	DownPrice     decimal.Decimal                     `json:"downPrice" form:"downPrice"`         // 优惠价格
	ProductPrices map[uint]map[string]decimal.Decimal `json:"productPrices" form:"productPrices"` // 商品价格
	TotalNum      int64                               `json:"totalNum" form:"totalNum"`           // 商品总数
	OrderType     int                                 `json:"orderType" form:"orderType"`         // 订单类型
}
