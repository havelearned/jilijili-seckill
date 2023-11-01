package sale

type SaleStatistics struct {
	// 统计id
	SaleId int64 `json:"saleId" form:"saleId" gorm:"primaryKey" `
	// 本场秒杀,秒杀表id,
	SekillId int64 `json:"sekillId" form:"sekillId" `
	// 销售数量
	SaleCount int64 `json:"saleCount" form:"saleCount" `
	// 销售额=销售收入=售出数量*单价
	TotalSales int64 `json:"totalSales" form:"totalSales" `
	// 总销售额
	TotalRevenue float32 `json:"totalRevenue" form:"totalRevenue" `
}
