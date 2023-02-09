package room

import (
	"hotel/admin"
	"hotel/service/room"

	"github.com/gin-gonic/gin"
)

func RoomList(c *gin.Context) {

	var service room.RoomListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RoomList()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
