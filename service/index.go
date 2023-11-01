package service

import (
	"_11160_jilijili_seckill/utils"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(200, utils.Response{
		Code:    500,
		Flag:    false,
		Data:    nil,
		Message: "操作成功",
	})

}

func TestApi(c *gin.Context) {
	c.JSON(200, utils.Response{
		Code:    200,
		Flag:    true,
		Data:    "hello",
		Message: "操作成功",
	})

}
