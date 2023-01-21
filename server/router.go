package server

import (
	"hotel/admin"
	"hotel/admin/user"

	"hotel/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// 路由
	Admin := r.Group("/admin")
	{
		Admin.POST("ping", admin.Ping)
		// 需要登录保护的
		auth := Admin.Group("")
		auth.Use(middleware.AuthRequired())
		{
			Admin.POST("login", user.UserLogin)
		}
	}
	return r
}
