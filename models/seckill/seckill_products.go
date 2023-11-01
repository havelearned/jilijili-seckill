package seckill

import "_11160_jilijili_seckill/models/products"

type SeckillProducts struct {
	// id
	SeckillProductsId int64 `json:"seckillProductsId" form:"seckillProductsId" gorm:"primaryKey" `
	// 秒杀表id
	SeckillId int64 `json:"seckillId" form:"seckillId" `
	// 商品id
	ProductsId int64 `json:"productsId" form:"products_Id" `
	// 库存数量
	StockQuantity int64 `json:"stockQuantity" form:"stockQuantity" `
	// 一对一关联商品
	Products products.Products
	// 添加专用
}

func (table *SeckillProducts) TableName() string {
	return "shop_seckill_products"
}
