package api

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"dingo/middleware"
)

// ping 测试程序是否正常连通
func Ping(api *echo.Group) {
	// api/v1
	v1 := api.Group("/v1")
	// api/v1/ping
	v1.GET("/ping", func(c echo.Context) error {
		return c.String(200, "v1-pong")
	})

	// api/v2
	v2 := api.Group("/v2")
	// api/v1/ping
	v2.GET("/ping", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		// user := util.GetUser()
		return c.HTML(200, "v2-pong <br /> 登录用户： "+name)
	}, middleware.Authorization())
}
