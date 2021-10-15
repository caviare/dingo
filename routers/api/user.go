package api

import (
	"net/http"

	"dingo/dto"
	"dingo/util"

	"github.com/labstack/echo/v4"
)

// ping 测试程序是否正常连通
func User(api *echo.Group) {

	// api/v1
	v1 := api.Group("/v1")
	// api/v1/login
	v1.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		var userData dto.User
		// 判断账号密码是否正确
		if username != "admin" || password != "a123123.." {
			return echo.ErrUnauthorized
		} else {
			userData = dto.User{
				Id:       1,
				Username: "刘钏",
			}
		}
		token, err := util.GenerateToken(userData, 72)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": token,
		})

	})
}
