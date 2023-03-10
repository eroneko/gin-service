package routers

import (
	api "github.com/eroneko/gin-service/internal/routers/api"
	"github.com/eroneko/gin-service/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	user := api.NewUser()
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/user", user.Register)
		apiGroup.DELETE("/user/:id", middlewares.AuthMiddleware(), user.Delete)
		apiGroup.PUT("/user/:id", middlewares.AuthMiddleware(), user.Update)
		apiGroup.GET("/user", user.Login)
		apiGroup.GET("/user/info", middlewares.AuthMiddleware(), user.Info)
	}
	return r
}
