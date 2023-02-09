package room

import (
	"hotel/model"
	"hotel/serializer"
	"hotel/util"
	"strconv"

	"gorm.io/gorm"
)

type RoomListService struct {
	util.Pager
	TypeID uint   `form:"typeID" json:"typeID"`
	Status string `form:"status"`
}

type RoomListRes struct {
	model.Room
	Type model.RoomType `gorm:"foreignKey:type_id;references:type_id" json:"type"`
}

func (service *RoomListService) RoomList() serializer.Response {

	var res []RoomListRes
	var total int64
	model.DB.Scopes(util.Paginate(service.Page, service.PageSize)).
		Scopes(service.Search()).
		Preload("Type").
		Find(&res).Limit(-1).Offset(-1).Count(&total)
	pageRes := serializer.PageListFormat(service.Page, service.PageSize, total, res)

	return serializer.ResponseFormat(200, "請求成功", pageRes)
}

func (service *RoomListService) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if service.Status != "" {
			status, _ := strconv.Atoi(service.Status)
			db.Where("status", status)
		}

		if service.TypeID != 0 {
			db.Where("type_id", service.TypeID)
		}

		return db
	}
}
