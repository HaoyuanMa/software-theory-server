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
		router.GET("build_connection", api.BuildConnection)

		router.POST("forward", api.Forward)
		router.POST("face_input", api.Register)
		router.POST("record_input", api.InputRecord)

		router.GET("user/test_get", api.TestGet)
		router.POST("user/test_post", api.TestPost)
	}

	var admin = r.Group("facemaskadmin")
	admin.Use(middleware.CORSMiddleware())
	{
		admin.GET("nomask/list", api.GetNoMaskList)
		admin.GET("employeesinfo/list", api.GetStaffList)
		admin.POST("user/login", api.Login)
	}

	_ = r.Run(":5000")
}
