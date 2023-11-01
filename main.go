package main

import (
	"_11160_jilijili_seckill/router"
	"_11160_jilijili_seckill/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	r.Run(":11160")
}
