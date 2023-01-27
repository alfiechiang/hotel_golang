package room

import (
	"hotel/admin"
	"hotel/service/room"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRoomType(c *gin.Context) {

	var service room.UpdateRoomTypeService
	if err := c.ShouldBind(&service); err == nil {
		roomID, _ := strconv.Atoi(c.Param("roomID"))
		service.RoomID = uint(roomID)
		res := service.UpdateRoomType()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}