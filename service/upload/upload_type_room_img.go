package upload

import (
	"hotel/model"
	"hotel/serializer"

)

type UploadTypeRoomImgService struct {
	model.RoomTypeImg
}

func (service *UploadTypeRoomImgService) UploadTypeRoomImg() serializer.Response {

	if err := model.DB.Create(&service).Error; err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.Success()
}
