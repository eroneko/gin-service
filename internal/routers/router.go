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
		// FIXME POST /users or /accounts
		apiGroup.POST("/user", user.Register)
		// FIXME DELETE /users or /accounts
		apiGroup.DELETE("/user/:id", middlewares.AuthMiddleware(), user.Delete)
		// FIXME PUT /users or /accounts
		apiGroup.PUT("/user/:id", middlewares.AuthMiddleware(), user.Update)
		// FIXME POST /sessions
		apiGroup.GET("/user", user.Login)
		// FIXME GET /sessions
		apiGroup.GET("/user/info", middlewares.AuthMiddleware(), user.Info)
		// Also see: https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design
	}
	return r
}
