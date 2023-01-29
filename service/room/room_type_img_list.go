package room

import (
	"hotel/model"
	"hotel/serializer"
)

type RoomTypeImgListService struct {
	TypeID uint `form:"typeID"`
}

func (service *RoomTypeImgListService) RoomTypeImgList() serializer.Response {

	var imgs []model.RoomTypeImg
	model.DB.Table("room_type_img").Where("type_id", service.TypeID).Find(&imgs)
	return serializer.ResponseFormat(200, "請求成功", imgs)
}
