package router

import (
	"_11160_jilijili_seckill/handle"
	"_11160_jilijili_seckill/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	baseGroup := r.Group("/")
	baseGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	baseGroup.GET("/index", service.GetIndex)
	baseGroup.GET("/test", service.TestApi)

	seckillGroup := r.Group("/seckill")
	seckillGroup.Use(handle.AuthUser(), handle.InterceptParams())
	// 获取秒杀场次列表
	seckillGroup.GET("/field/list", service.GetSeckillList)

	//// 通过秒杀场次_查询商品列表
	seckillGroup.GET("/field/:seckillId", service.GetProductListBySeckillId)
	//// 添加秒杀场次
	seckillGroup.POST("/field", service.AddSeckill)
	//// 通过秒杀场次ID_添加商品`列表
	seckillGroup.POST("/field/addList", service.AddProductListBySeckillId)
	//// 通过秒杀场次ID_修改商品列表
	seckillGroup.PUT("/field", service.PutProductListBySeckillId)
	//// 通过秒杀场次ID_删除商品列表
	seckillGroup.DELETE("/field", service.DelProductListBySeckillId)

	//saleGroup := r.Group("/sale")
	//// 获取秒杀统计数据
	//saleGroup.GET("/list", t.list) //不设置限制条件的画默认查询所有
	//saleGroup.GET("/one", t.one)
	//saleGroup.POST("/insert", t.insert)
	//saleGroup.POST("/update", t.update)
	//saleGroup.POST("/delete", t.delete)

	//// 秒杀商品
	//seckillGroup.GET("/kill/:productId/:userId", service.KillProduct)

	return r
}
