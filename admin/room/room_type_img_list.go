package room

import (
	"hotel/admin"
	"hotel/service/room"
	"github.com/gin-gonic/gin"
)

func RoomTypeImgList(c *gin.Context) {

	var service room.RoomTypeImgListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RoomTypeImgList()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
