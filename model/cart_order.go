package model

// CartOrder 用户标签关系表
type CartOrder struct {
	CartId  uint `json:"cartId" form:"cartId" gorm:"column:cart_id;comment:关联标记"`
	OrderId uint `gorm:"column:order_id;" json:"orderId"`
}
