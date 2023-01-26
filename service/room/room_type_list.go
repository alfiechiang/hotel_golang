package room

import (
	"hotel/model"
	"hotel/serializer"
	"hotel/util"
)

type RoomTypeListService struct {
	util.Pager
}

func (service *RoomTypeListService) RoomTypeList() serializer.Response {

	var res []model.RoomType
	var total int64
	model.DB.Scopes(util.Paginate(service.Page, service.PageSize)).Find(&res).Limit(-1).Offset(-1).Count(&total)
	pageRes := serializer.PageListFormat(service.Page, service.PageSize, total, res)

	return serializer.ResponseFormat(200, "請求成功", pageRes)
}
