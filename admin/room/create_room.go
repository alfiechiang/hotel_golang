package room

import (
	"hotel/admin"
	"hotel/service/room"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {

	var service room.CreateRoomService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateRoom()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
