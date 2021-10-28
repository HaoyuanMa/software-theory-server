package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())

	var router = r.Group("api")
	{
		router.GET("user/test_get", api.Test_get)
		router.POST("user/test_post", api.Test_post)
	}

	var auth = r.Group("api")
	auth.Use(middleware.JwtToken())
	{

	}

	_ = r.Run(":8080")
}
