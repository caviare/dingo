package model

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB
var RDB *redis.Client

// 连接mysql数据库
func ConnectMysql() {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接MYSQL数据库失败")
	} else {
		fmt.Println("")
		fmt.Println("\033[0;32m------------MYSQL连接成功------------\033[0m")
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("数据库丢失: ", err)
			panic(err)
		}
		//设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		//设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(20)

		// 全局变量外部调用
		DB = db

		// 自动迁移
		migration()
	}
}

// 连接redis数据库
func ConnctRedis() {
	dsn := os.Getenv("REDIS_DSN")
	// 连接redis
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatal("连接REDIS数据库失败")
	} else {
		fmt.Println("")
		fmt.Println("\033[0;32m------------REDIS连接成功------------\033[0m")
	}
	RDB = redis.NewClient(opt)
}
