package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 用户登录认证
func Authorization() echo.MiddlewareFunc {
	secretKey := os.Getenv("SECRET_KEY")
	return middleware.JWT([]byte(secretKey))
}
