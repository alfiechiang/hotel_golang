package room

import (
	"hotel/admin"
	"hotel/service/room"

	"github.com/gin-gonic/gin"
)

func CreateRoomType(c *gin.Context) {

	var service room.CreateRoomTypeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateRoomType()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
