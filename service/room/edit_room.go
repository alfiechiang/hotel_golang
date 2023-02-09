package room

import (
	"hotel/model"
	"hotel/serializer"
)

type EditRoomService struct {
	ID uint 
}

func (service *EditRoomService) EditRoom() serializer.Response {

	var room model.Room
	model.DB.Table("room").First(&room, service.ID)

	res := room
	return serializer.ResponseFormat(200, "請求成功", res)
}
