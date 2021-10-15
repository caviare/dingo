package util

import (
	"dingo/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	dto.User
	jwt.StandardClaims
}

// 生成返回给用户的 token
func GenerateToken(userData dto.User, expireDuration time.Duration) (string, error) {
	// 由于 jwt 返回的 Token 中的数据仅做了 Base64 处理，没有加密，所以不应放入重要的信息。
	// jwt Token 由于是无状态的，任何获取到此 Token 的人都可以访问，所以为了减少盗用，可以将 Token 有效期设置短一些。对一些重要的操作，尽量再次进行认证。
	// 网站尽量使用 HTTPS，可以减少 Token 的泄漏。

	// token存储的数据、过期时间....
	// expireDuration：过期时间，单位（小时）
	claims := MyCustomClaims{
		userData,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * expireDuration).Unix(),
		},
	}
	// 将数据写入到token中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// SecretKey 用于对用户数据进行签名，不能暴露
	secretKey := os.Getenv("SECRET_KEY")
	// 生成返回给用户的 token
	return token.SignedString([]byte(secretKey))
}
