package api

import (
	"dingo/service"

	"github.com/labstack/echo/v4"
)

// ping 测试程序是否正常连通
func User(api *echo.Group) {

	// api/v1
	v1 := api.Group("/v1")
	// api/v1/login
	v1.POST("/login", func(c echo.Context) error {
		return c.JSON(200, service.Login(c))
	})
}
