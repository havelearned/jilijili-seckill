package seckill

type SeckillSaleStatistics struct {
	// 统计id
	SaleId int64 `json:"saleId" form:"saleId" gorm:"primaryKey" `
	// 本场秒杀,秒杀表id,
	StatisticId int64 `json:"statisticId" form:"statisticId" `
	// 销售数量
	SaleCount int64 `json:"saleCount" form:"saleCount" `
	// 销售额=销售收入=售出数量*单价
	TotalSales int64 `json:"totalSales" form:"totalSales" `
	// 总销售额
	TotalRevenue interface{} `json:"totalRevenue" form:"totalRevenue" `
}

func (table *SeckillSaleStatistics) TableName() string {
	return "shop_seckill_sale_statistics"
}
