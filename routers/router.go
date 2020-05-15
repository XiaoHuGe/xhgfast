package routers

import (
	"github.com/gin-gonic/gin"
	"xhgfast/controller"
	"xhgfast/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.CORSMiddleware())

	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	u := r.Group("/api/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/info", controller.Info)
	}
	return r

}
