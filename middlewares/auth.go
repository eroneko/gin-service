package middlewares

import (
	"github.com/eroneko/gin-service/global"
	"github.com/eroneko/gin-service/internal/dao"
	"github.com/eroneko/gin-service/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取Authorization Header
		tokenString := c.GetHeader("Authorization")
		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := app.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}
		//验证通过后，获取Claims中的UserID
		userID := claims.UserID
		d := dao.New(global.DBEngine)
		if !d.IsUserExist(dao.User{ID: userID}) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户不存在",
			})
			c.Abort()
			return
		}
		response, err := d.GetUserByID(dao.User{ID: userID})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "查询用户信息失败",
			})
			c.Abort()
			return
		}
		//用户存在将User信息写入上下文
		c.Set("user", response)
		c.Next()
	}
}
