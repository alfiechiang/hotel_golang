package room

import (
	"hotel/model"
	"hotel/serializer"
)

type UpdateRoomService struct {
	model.Room
}

func (service *UpdateRoomService) UpdateRoom() serializer.Response {

	var room model.Room
	model.DB.Table("room").First(&room, service.ID)
	room.TypeID = service.TypeID
	room.Description = service.Description
	room.Status = service.Status

	if err := model.DB.Table("room_type").Save(&room).Error; err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.Success()
}
