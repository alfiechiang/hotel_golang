package room

import (
	"hotel/model"
	"hotel/serializer"
)

type CreateRoomService struct {
	model.Room
}

func (service *CreateRoomService) CreateRoom() serializer.Response {



	var roomType model.RoomType
	model.DB.Where("type_id",service.TypeID).Find(&roomType)
	if roomType == (model.RoomType{}){
		return serializer.Err(40001, "房型ID不存在", nil)
	}

	if err := model.DB.Create(service).Error; err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.ResponseFormat(200, "請求成功", make([]string, 0))
}
