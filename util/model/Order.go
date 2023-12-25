package model

// Order 订单表
type Order struct {
	ID        int `gorm:"primary_key"`
	ProductID int `gorm:"column:product_id"` // 商品ID
	UserID    int `gorm:"column:user_id"`    // 用户ID
}

func (Order) TableName() string {
	return "orders"
}
