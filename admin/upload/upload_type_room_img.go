package upload

import (
	"hotel/admin"
	"hotel/service/upload"

	"github.com/gin-gonic/gin"
)

func UploadTypeRoomImg(c *gin.Context) {

	var service upload.UploadTypeRoomImgService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadTypeRoomImg()
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
