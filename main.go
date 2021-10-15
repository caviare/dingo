package main

import (
	"dingo/model"
	"dingo/routers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("配置文件加载失败。")
	}

	// 连接数据库
	model.ConnectMysql()

	// 初始化路由
	e := routers.InitRouters()

	// 启动服务 端口2233
	e.Logger.Fatal(e.Start(":2233"))
}
