package routers

import (
	"github.com/eroneko/gin-service/internal/routers/api"
	"github.com/eroneko/gin-service/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	user := api.NewUser()
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/users", user.Register)
		apiGroup.DELETE("/users/:id", middlewares.AuthMiddleware(), user.Delete)
		apiGroup.PUT("/users/:id", middlewares.AuthMiddleware(), user.Update)
		apiGroup.POST("/sessions", user.Login)
		apiGroup.GET("/user/info", middlewares.AuthMiddleware(), user.Info)
	}
	return r
}
