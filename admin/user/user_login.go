package user

import (
	"hotel/service/user"
	"hotel/admin"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {

	var service user.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserLogin(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, admin.ErrorResponse(err))
	}

}
