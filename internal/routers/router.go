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
	r.Static("/static", "./internal/static")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLFiles("./internal/templates/index.html", "./internal/templates/register.html", "./internal/templates/info.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	r.GET("/info", func(c *gin.Context) {
		c.HTML(200, "info.html", nil)
	})
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
