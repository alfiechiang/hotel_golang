package upload

import (
	"hotel/admin"
	"hotel/service/upload"

	"github.com/gin-gonic/gin"
)

func RemoveImage(c *gin.Context) {

	var service upload.RemoveImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RemoveImage(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
