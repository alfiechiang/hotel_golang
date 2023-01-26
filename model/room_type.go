package model

type RoomType struct {
	ID     uint   `gorm:"primarykey"`
	RoomID uint   `json:"roomID"`
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	BedNum uint8  `json:"bed_num"`
}

func (RoomType) TableName() string {
	return "room_type"
}
