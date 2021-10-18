package main

import (
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
		log.Fatal("配置文件加载失败。")
	}

	// 连接MYSQL数据库
	model.ConnectMysql()
	rdsn := os.Getenv("REDIS_DSN")
	if rdsn != "" {
		// 连接REDIS数据库
		model.ConnctRedis()
	}

	// 初始化路由
	e := routers.InitRouters()

	// 启动服务 端口2233
	e.Logger.Fatal(e.Start(":2233"))
}
