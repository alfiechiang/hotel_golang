package room

import (
	"hotel/admin"
	"hotel/service/room"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRoom(c *gin.Context) {

	var service room.UpdateRoomService
	if err := c.ShouldBind(&service); err == nil {
		ID, _ := strconv.Atoi(c.Param("ID"))
		service.ID = uint(ID)
		res := service.UpdateRoom()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
