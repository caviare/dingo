package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"dingo/routers/api"
	"dingo/routers/render"
)

// 初始化路由
func InitRouters() *echo.Echo {
	e := echo.New()
	// 打印log记录
	e.Use(middleware.Logger())

	// 程序崩溃时打印错误信息 && 恢复正常执行
	e.Use(middleware.Recover())

	// 跨域配置
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 指定域名示例
		// AllowOrigins: []string{"http://192.168.1.16, http://www.baidu.com"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// api分组对象
	apiGroup := e.Group("/api")

	// api
	api.Ping(apiGroup)
	api.User(apiGroup)

	// 渲染页面
	render.Page(e)

	return e
}
