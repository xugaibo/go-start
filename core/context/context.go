package context

import (
	"fmt"
	logger2 "go-start/core/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB = nil
var Log *logger2.Log

func InitDb() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	Log = logger2.NewLog(newLogger)

	dsn := "root:123456@tcp(127.0.0.1:3306)/go_start?charset=utf8mb4&parseTime=True&loc=Local"
	dbGet, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger, NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Println("connect mysql bizerror")
		panic(err)
	}
	Db = dbGet
}
