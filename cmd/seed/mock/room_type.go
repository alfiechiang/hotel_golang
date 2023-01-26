package mock

import "hotel/model"

func MockRoomType() {

	model.DB.Exec("TRUNCATE TABLE room_type")

	types := []model.RoomType{
		{RoomID: 1, Name: "標準間", Price: 200, BedNum: 2},
		{RoomID: 2, Name: "單人間", Price: 150, BedNum: 1},
		{RoomID: 3, Name: "大床房", Price: 280, BedNum: 1},
		{RoomID: 4, Name: "三人間", Price: 300, BedNum: 3},
		{RoomID: 5, Name: "標準套房", Price: 380, BedNum: 2},
		{RoomID: 6, Name: "豪華套房", Price: 580, BedNum: 2},
		{RoomID: 7, Name: "商務套房", Price: 880, BedNum: 3},
		{RoomID: 8, Name: "行政套房", Price: 1000, BedNum: 3},
		{RoomID: 9, Name: "總統套房", Price: 2000, BedNum: 4},
	}

	model.DB.Create(&types)

}
