package service

import (
	"dingo/dto"
	"dingo/model"
	"dingo/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string
	Password string
}

// 用户登录后TOKEN存入用户ID
type User struct {
	UserId int `json:"user_id"`
}

type MyCustomClaims struct {
	User
	jwt.StandardClaims
}

// Login 用户登录函数
func Login(c echo.Context) dto.Response {
	username := c.FormValue("username")
	password := c.FormValue("password")
	// 加密数据 进行判断
	password = utils.Encrypt(password)
	var userData dto.User
	data := dto.User{}
	// 查找登录账号信息
	result := model.DB.Where("username = ? AND password = ?", username, password).First(&data)
	if result.Error != nil {
		return dto.ParamErr("账号不存在或密码错误", result.Error)
	}

	// 生成TOKEN
	token, err := GenerateToken(userData.Id, 72)
	if err != nil {
		return dto.Err(40001, "TOKEN生成失败", err)
	}
	data.Token = token
	return dto.Response{
		Data: data,
	}
}

// 生成返回给用户的 token
func GenerateToken(userId int, expireDuration time.Duration) (string, error) {
	// 由于 jwt 返回的 Token 中的数据仅做了 Base64 处理，没有加密，所以不应放入重要的信息。
	// jwt Token 由于是无状态的，任何获取到此 Token 的人都可以访问，所以为了减少盗用，可以将 Token 有效期设置短一些。对一些重要的操作，尽量再次进行认证。
	// 网站尽量使用 HTTPS，可以减少 Token 的泄漏。

	// token存储的数据、过期时间....
	// expireDuration：过期时间，单位（小时）
	claims := MyCustomClaims{
		User{
			UserId: userId,
		},
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
