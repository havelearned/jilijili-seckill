package service

import (
	"_11160_jilijili_seckill/models/seckill"
	"_11160_jilijili_seckill/utils"
	"github.com/gin-gonic/gin"
)

// GetSeckillList
// @Summary 获取本场秒杀列表
// @Param page query  string true "当前页"
// @Param size query  string true "页面大小"
// @Param time query  string false "结束时间"
// @Success 200 (string}json{"flag","data","message"}
// @Router /seckill/kill/list [get]
func GetSeckillList(c *gin.Context) {

	page, size := utils.GetPageAndSize(c)
	endTimeStr := c.Query("time")

	seckillList, _ := seckill.GetSeckillList(size, page, endTimeStr)

	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    seckillList,
		Message: "",
	})
}

// GetProductListBySeckillId 通过秒杀场次_查询商品列表
func GetProductListBySeckillId(c *gin.Context) {
	seckillId := c.Param("seckillId")
	pageInt, pageSize := utils.GetPageAndSize(c)
	productList, _ := seckill.GetProductListBySeckillId(seckillId, pageInt, pageSize)

	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    productList,
		Message: "",
	})

}

// AddSeckill 添加秒杀场次
func AddSeckill(c *gin.Context) {
	var sk seckill.Seckill
	err := c.ShouldBindJSON(&sk)

	if err != nil {
		c.JSON(500, utils.Response{
			Code:    500,
			Flag:    false,
			Data:    err,
			Message: "操作失败",
		})
		return
	}
	utils.DB.Create(&sk)
	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    sk,
		Message: "操作成功",
	})
}

// PutSeckill 修改秒杀场次
func PutSeckill(c *gin.Context) {
	var sk seckill.Seckill
	err := c.ShouldBindJSON(&sk)

	if err != nil {
		c.JSON(500, utils.Response{
			Code:    500,
			Flag:    false,
			Data:    err,
			Message: "操作失败",
		})
		return
	}
	// 通过id修改
	utils.DB.Where("seckill_id = ?", sk.SeckillId).Updates(sk)
	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    sk,
		Message: "操作成功",
	})

}

// AddProductListBySeckillId 通过秒杀场次ID_添加商品列表
func AddProductListBySeckillId(c *gin.Context) {
	var seckillProductsArray []seckill.SeckillProducts
	if err := c.ShouldBindJSON(&seckillProductsArray); err != nil {
		c.JSON(200, utils.Response{
			Code:    405,
			Flag:    false,
			Data:    err.Error(),
			Message: "操作失败",
		})
		return
	}

	utils.DB.Model(seckillProductsArray).CreateInBatches(seckillProductsArray, 100)
	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    seckillProductsArray,
		Message: "操作成功",
	})

}

// PutProductListBySeckillId  通过秒杀场次ID_修改商品列表
func PutProductListBySeckillId(c *gin.Context) {
	var sp seckill.SeckillProducts
	if err := c.ShouldBindJSON(&sp); err != nil {
		c.JSON(200, utils.Response{
			Code:    405,
			Flag:    false,
			Data:    err.Error(),
			Message: "参数错误",
		})
		return
	}
	if err := utils.DB.Model(&sp).
		Where("seckill_id = ? and products_id =?", sp.SeckillId, sp.ProductsId).
		Update("stock_quantity", sp.StockQuantity).Error; err != nil {
		c.JSON(200, utils.Response{
			Code:    500,
			Flag:    false,
			Data:    err.Error(),
			Message: "操作失败",
		})
		return
	}

	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    "",
		Message: "操作成功",
	})

}

// DelProductListBySeckillId  通过秒杀场次ID_删除商品列表
func DelProductListBySeckillId(c *gin.Context) {
	var sp seckill.SeckillProducts
	if err := c.ShouldBindJSON(&sp); err != nil {
		c.JSON(200, utils.Response{
			Code:    405,
			Flag:    false,
			Data:    err.Error(),
			Message: "参数错误",
		})
		return
	}
	utils.DB.Where("seckill_id=? and products_id=?", sp.SeckillId, sp.ProductsId).Delete(&sp)
	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    "",
		Message: "操作成功",
	})

}

// GetSaleList  获取秒杀统计数据
func GetSaleList(c *gin.Context) {

}

// GetSaleListBySeckillId  通过秒杀场次id查询统计数据
func GetSaleListBySeckillId(c *gin.Context) {

}

// KillProduct 秒杀商品
func KillProduct(c *gin.Context) {

}
