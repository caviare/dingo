package main

import "dingo/routers"

func main() {
	// 初始化路由
	e := routers.InitRouters()
	// 启动服务 端口2233
	e.Logger.Fatal(e.Start(":2233"))
}
