package main

import (
	"dingo/cache"
	"dingo/model"
	"dingo/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("配置文件加载失败，请检查相关配置文件重新启动。")
	}
	var dsn string

	// 连接MYSQL数据库
	dsn = os.Getenv("MYSQL_DSN")
	if dsn != "" {
		model.ConnectMysql()
	} else {
		log.Fatal("未找到MYSQL配置，请检查相关配置文件重新启动。")
	}

	// 连接REDIS数据库
	dsn = os.Getenv("REDIS_DSN")
	if dsn != "" {
		cache.ConnctRedis()
	}

	// 初始化路由
	e := routers.Init()

	// 启动服务 端口2233
	e.Logger.Fatal(e.Start(":2233"))
}
