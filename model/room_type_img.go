package model

type RoomTypeImg struct {
	ID     uint   `gorm:"primarykey"`
	TypeID uint   `json:"typeID"`
	ImgUrl string `json:"img_url"`
}

func (RoomTypeImg) TableName() string {
	return "room_type_img"
}
