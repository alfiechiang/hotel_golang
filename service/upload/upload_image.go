package upload

import (
	"hotel/serializer"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type UploadImageService struct {
}

type UploadImageResponse struct {
	ImgUrl string `json:"img_url"`
	ImgName string `json:"img_name"`
}

func (service *UploadImageService) UploadImage(c *gin.Context) serializer.Response {

	file, header, err := c.Request.FormFile("image")
	filename := header.Filename
	dir, _ := os.Getwd()
	out, err := os.Create(dir + "/img/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	res := UploadImageResponse{
		ImgUrl: os.Getenv("DOMAIN") + "/img/" + filename,
		ImgName: filename,
	}

	return serializer.ResponseFormat(200, "請求成功", res)
}
