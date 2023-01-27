package room

import (
	"hotel/model"
	"hotel/serializer"
)

type CreateRoomTypeService struct {
	model.RoomType
}

func (service *CreateRoomTypeService) CreateRoomType() serializer.Response {

	if err := model.DB.Create(service).Error; err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.ResponseFormat(200, "請求成功", make([]string, 0))
}
