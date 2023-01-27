package room

import (
	"hotel/model"
	"hotel/serializer"
)

type EditRoomTypeService struct {
	RoomID uint `json:"roomID"`
}

func (service *EditRoomTypeService) EditRoomType() serializer.Response {

	var roomType model.RoomType
	model.DB.Table("room_type").First(&roomType, service.RoomID)
	return serializer.ResponseFormat(200, "請求成功", roomType)
}
