package utils

import (
	// 解析yml配置信息
	"github.com/spf13/viper"
	// mysql驱动
	"gorm.io/driver/mysql"
	// gorm映射框架
	"gorm.io/gorm"
	// gorm日志
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	// 从顶级目录(我这里是:_11160_jilijili_seckill)开始查询config/app.yml文件
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("yml配置文件读取失败:", err)
		return
	}
	log.Println("config.app:", viper.Get("server"))
	log.Println("config.app:", viper.Get("gin"))

}

func InitMySql() *gorm.DB {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:        time.Second,
		Colorful:             true,
		ParameterizedQueries: false,
		LogLevel:             logger.Info,
	})
	DB, _ = gorm.Open(mysql.Open(viper.GetString("gin.datasource.url")), &gorm.Config{
		Logger: newLogger,
	})

	DBConnectionTest()

	return DB

}

func DBConnectionTest() {
	row := DB.Unscoped().Raw("SELECT 1").Row()
	var test int
	row.Scan(&test)
	log.Println("连接数据测试,查询sql: SELECT 1 result\n", test)
}
