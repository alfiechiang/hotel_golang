package model

type Room struct {
	ID          uint   `gorm:"primarykey"`
	RoomID      uint   `gorm:"unique" json:"roomID"`
	TypeID      uint   `json:"typeID"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      bool
}

func (Room) TableName() string {
	return "room"
}
