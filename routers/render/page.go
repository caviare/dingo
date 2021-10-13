package render

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// 若未匹配到路由，则对应VIEW文件夹下的HTML页面
func Page(e *echo.Echo) {

	e.GET("/*", func(c echo.Context) error {
		url := c.Request().URL.Path
		// 根路由默认渲染 index.html
		if url == "/" {
			return c.File("view/index.html")
		}
		path := strings.TrimSuffix(url, ".html")
		return c.File("view/" + path + ".html")
	})
}
