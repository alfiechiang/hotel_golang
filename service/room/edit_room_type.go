package room

import (
	"hotel/model"
	"hotel/serializer"
)

type EditRoomTypeService struct {
	TypeID uint `json:"typeID"`
}

func (service *EditRoomTypeService) EditRoomType() serializer.Response {

	var roomType model.RoomType
	model.DB.Table("room_type").First(&roomType, service.TypeID)
	return serializer.ResponseFormat(200, "請求成功", roomType)
}
