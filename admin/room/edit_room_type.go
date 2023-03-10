package room

import (
	"hotel/admin"
	"hotel/service/room"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditRoomType(c *gin.Context) {

	var service room.EditRoomTypeService
	if err := c.ShouldBind(&service); err == nil {

		typeID, _ := strconv.Atoi(c.Param("typeID"))
		service.TypeID = uint(typeID)
		res := service.EditRoomType()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
