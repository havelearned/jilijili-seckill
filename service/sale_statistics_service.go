package service

import (
	"_11160_jilijili_seckill/models/sale"
	"_11160_jilijili_seckill/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var SaleStatisticsService = saleStatisticsService{}

// saleStatisticsService 业务层
type saleStatisticsService struct {
}

func (t saleStatisticsService) db() *gorm.DB {
	return utils.DB
}

// List 分页列表
func (t saleStatisticsService) List(page, size int, v *sale.SaleStatistics) gin.H {
	lists := make([]sale.SaleStatistics, 0) // 结果
	t.db().Model(&sale.SaleStatistics{}).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	var total int64 // 统计
	t.db().Model(&v).Where(&v).Count(&total)
	h := gin.H{"list": lists, "total": total}
	return h
}

// One 根据主键查询
func (t saleStatisticsService) One(id interface{}) (v sale.SaleStatistics) {
	t.db().Find(&v, id)
	return
}

// Update 修改记录
func (t saleStatisticsService) Update(v sale.SaleStatistics) {
	t.db().Model(&v).Updates(v)
}

// Insert 插入记录
func (t saleStatisticsService) Insert(v sale.SaleStatistics) {
	t.db().Save(&v)
}

// Delete 根据主键删除
func (t saleStatisticsService) Delete(ids []int) {
	t.db().Delete(sale.SaleStatistics{}, ids)
}
