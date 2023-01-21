package user

import (
	"hotel/auth"
	"hotel/model"

	"hotel/serializer"

	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) UserLogin(c *gin.Context) serializer.Response {

	var user model.User
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	token, _ := auth.GenerateToken(user)
	data := make(map[string]interface{}, 0)
	data["access_token"] = token

	return serializer.ResponseFormat(200, "登入成功", data)
}
