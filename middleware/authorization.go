package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 用户登录认证
func Authorization() echo.MiddlewareFunc {
	return middleware.JWT([]byte("secret"))
}
