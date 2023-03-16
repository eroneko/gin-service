package routers

import (
	_ "github.com/eroneko/gin-service/docs"
	"github.com/eroneko/gin-service/internal/middlewares"
	"github.com/eroneko/gin-service/internal/routers/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	user := api.NewUser()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := r.Group("/")
	{
		apiGroup.POST("/users", user.Register)
		apiGroup.DELETE("/users/:id", middlewares.AuthMiddleware(), user.Delete)
		apiGroup.PUT("/users/:id", middlewares.AuthMiddleware(), user.Update)
		apiGroup.POST("/sessions", user.Login)
		apiGroup.GET("/user/info", middlewares.AuthMiddleware(), user.Info)
	}
	return r
}
