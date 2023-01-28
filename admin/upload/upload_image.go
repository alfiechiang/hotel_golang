package upload

import (
	"hotel/service/upload"
	"github.com/gin-gonic/gin"
	"hotel/admin"
)

func UploadImage(c *gin.Context) {


	var service upload.UploadImageService
	if err := c.ShouldBind(&service); err == nil {
		res:=service.UploadImage(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}
	
}
