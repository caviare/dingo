package routers

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// 初始化路由
func InitRouters() *echo.Echo {
	e := echo.New()

	// v1版本路由组
	v1 := e.Group("/api/v1")
	v1.GET("/ping", func(c echo.Context) error {
		return c.String(200, "v1-pong")
	})
	// v2版本路由组
	v2 := e.Group("/api/v2")
	v2.GET("/ping", func(c echo.Context) error {
		return c.String(200, "v2-pong")
	})

	// 其他路由默认渲染页面
	e.GET("/*", func(c echo.Context) error {
		url := c.Request().URL.Path
		// 根路由默认渲染index.html
		if url == "/" {
			return c.File("view/index.html")

		}
		path := strings.TrimSuffix(url, ".html")
		return c.File("view/" + path + ".html")
	})
	return e
}
