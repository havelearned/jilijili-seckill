package main

//
//import (
//	"_11160_jilijili_seckill/models/seckill"
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//	"time"
//)
//
//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	dsn := "root:123456.ljj@tcp(127.0.0.1:13306)/jilijili_system?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("数据库连接错误!!!", err)
//	}
//
//	// 自动迁移表结构
//	//db.AutoMigrate(&seckill.Seckill{})
//
//	var seckills []seckill.Seckill
//	// 查询前 10 条记录
//	db.Limit(1000).Find(&seckills)
//
//	for _, seckill := range seckills {
//		fmt.Printf("Seckill ID: %d\n", seckill.SeckillId)
//		fmt.Printf("Product ID: %d\n", seckill.SeckillProductsId)
//		fmt.Printf("Description: %s\n", seckill.Desc)
//		fmt.Printf("Created Time: %s\n", seckill.CreatedTime.Format(time.RFC3339))
//		fmt.Printf("End Time: %s\n", seckill.EndTime.Format(time.RFC3339))
//		fmt.Println("----------")
//	}
//}
//
//func installDemo(db *gorm.DB) {
//	// 创建一个新 Seckill 记录
//	newSeckill := seckill.Seckill{
//		SeckillId:         1,    // 自己设定的假数据
//		SeckillProductsId: 1001, // 自己设定的假数据
//		Desc:              "假数据描述",
//		EndTime:           time.Now().Add(24 * time.Hour), // 结束时间设为当前时间加 24 小时
//	}
//
//	// 插入记录
//	result := db.Create(&newSeckill)
//	if result.Error != nil {
//		log.Fatal("插入记录失败!!!", result.Error)
//	}
//
//	//fmt.Printf("插入记录成功，插入的记录ID为: %d\n", newSeckill.ID)
//}
