package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

// 连接mysql数据库
func ConnectMysql() {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接数据库失败")
	} else {
		fmt.Println("-------------------------------------")
		fmt.Println("------------数据库连接成功-----------")
		fmt.Println("-------------------------------------")
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
