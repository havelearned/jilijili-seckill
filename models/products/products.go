package products

import (
	"time"
)

type Products struct {
	// 商品ID
	ProductId int64 `json:"productId" form:"productId" gorm:"primaryKey;column:product_id" `
	// 分类ID
	CategoryId int64 `json:"categoryId" form:"categoryId" gorm:"column:category_id" `
	// 商品名称
	ProductName string `json:"productName" form:"productName" gorm:"column:product_name" `
	// 商品描述
	Description string `json:"description" form:"description" gorm:"column:description" `
	// 价格
	Price float32 `json:"price" form:"price" gorm:"column:price" `
	// 库存数量
	StockQuantity int64 `json:"stockQuantity" form:"stockQuantity" gorm:"column:stock_quantity" `
	// 图片链接
	ImageUrl string `json:"imageUrl" form:"imageUrl" gorm:"column:image_url" `
	// 创建时间
	CreatedTime time.Time `json:"createdTime" form:"createdTime" gorm:"autoCreateTime;column:created_time" `
	// 更新时间
	UpdatedTime time.Time `json:"updatedTime" form:"updatedTime" gorm:"autoUpdateTime;column:updated_time" `
}

func (table *Products) TableName() string {
	return "shop_products"
}
