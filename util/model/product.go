package model

// Product 商品表
type Product struct {
	ID    int `gorm:"primary_key"`
	SkuId int `gorm:"column:sku_id"` // 商品标识
	Num   int `gorm:"column:num"`    // 库存数量
}

func (Product) TableName() string {
	return "product"
}
