package room

import (
	"hotel/admin"
	"hotel/service/room"

	"github.com/gin-gonic/gin"
)

func RoomTypeList(c *gin.Context) {

	var service room.RoomTypeListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RoomTypeList()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
