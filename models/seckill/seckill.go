package seckill

import (
	"_11160_jilijili_seckill/models"
	"_11160_jilijili_seckill/utils"
	"log"
)

type Seckill struct {
	// id
	SeckillId int64 `json:"seckillId" form:"seckillId" gorm:"primaryKey;column:seckill_id" `
	// 秒杀开场描述
	Desc string `json:"desc" form:"desc" `
	// 秒杀开始时间
	CreatedTime models.MyTime `json:"createdTime" form:"createdTime" gorm:"column:created_time"`
	// 结束时间
	EndTime models.MyTime `json:"endTime" form:"endTime" gorm:"column:end_time"`
	// 多个商品关联
	SeckillProducts []SeckillProducts
}

func (table *Seckill) TableName() string {
	return "shop_seckill"
}

// GetSeckillList 获取场次秒杀信息
func GetSeckillList(size, page int, endTime string) (utils.PageDto, error) {
	skills := make([]*Seckill, 10)
	offset := size * (page - 1)

	var total int64
	utils.DB.Model(&skills).Count(&total)
	query := utils.DB.Limit(size).Offset(offset)
	if endTime != "" {
		query = query.Where("end_time = ?", endTime)
		utils.DB.Where("end_time = ?", endTime).Find(&skills).Count(&total)
	}

	err := query.Find(&skills).Error

	result := utils.PageDto{
		Current: page,
		Size:    size,
		Records: skills,
		Total:   total,
	}

	log.Println(len(skills))
	return result, err
}

func GetProductListBySeckillId(seckillId string, page, size int) (utils.PageDto, error) {
	var seckillProducts []SeckillProducts
	var seckillProduct SeckillProducts
	offset := (page - 1) * size // 计算偏移量
	var total int64
	utils.DB.Table("shop_seckill_products").
		Select("shop_products.*, shop_seckill_products.*").
		Joins("left join shop_products on shop_products.product_id = shop_seckill_products.products_id").
		Where("shop_seckill_products.seckill_id = ?", seckillId).Count(&total)
	var err error
	// 执行关联查询
	err = utils.DB.Preload("Products").Take(&seckillProduct).
		Where("shop_seckill_products.seckill_id = ?", seckillId).
		Offset(offset).Limit(size).
		Find(&seckillProducts).Error

	return utils.PageDto{
		Current: page,
		Size:    size,
		Records: seckillProducts,
		Total:   total,
	}, err

}
