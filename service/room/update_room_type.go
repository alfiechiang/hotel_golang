package room

import (
	"hotel/model"
	"hotel/serializer"
)

type UpdateRoomTypeService struct {
	model.RoomType
}

func (service *UpdateRoomTypeService) UpdateRoomType() serializer.Response {

	var roomType model.RoomType
	model.DB.Table("room_type").First(&roomType, service.RoomID)
	roomType.BedNum = service.BedNum
	roomType.Name = service.Name
	roomType.Price = service.Price

	if err := model.DB.Table("room_type").Save(&roomType).Error; err != nil {
		return serializer.Err(40001, "系統錯誤", nil)
	}

	return serializer.Success()
}
