package server

import (
	"hotel/admin"
	"hotel/admin/room"
	"hotel/admin/user"

	"hotel/admin/upload"

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
			Admin.GET("roomtype", room.RoomTypeList)
			Admin.POST("roomtype", room.CreateRoomType)
			Admin.GET(":typeID/roomtype", room.EditRoomType)
			Admin.PUT("roomtype/:typeID", room.UpdateRoomType)

			Admin.POST("upload", upload.UploadImage)
		}
	}
	return r
}
