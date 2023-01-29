package upload

import (
	"hotel/serializer"
	"os"

	"github.com/gin-gonic/gin"
)

type RemoveImageService struct {
	ImgName string `json:"img_name"`
}

func (service *RemoveImageService) RemoveImage(c *gin.Context) serializer.Response {

	dir, _ := os.Getwd()
	path := dir + "/img/" + service.ImgName
	if err := os.Remove(path); err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.Success()
}
