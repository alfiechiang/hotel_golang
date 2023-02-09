package model

type RoomImg struct {
	ID      uint   `gorm:"primarykey"`
	RoomID  uint   `json:"roomID"`
	ImgUrl  string `json:"img_url"`
	ImgName string `json:"img_name"`
}

func (RoomImg) TableName() string {
	return "room_img"
}
