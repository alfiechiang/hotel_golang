package mock

import "hotel/model"

func MockRoomType() {

	model.DB.Exec("TRUNCATE TABLE room_type")

	types := []model.RoomType{
		{TypeID: 1, Name: "標準間", Price: 200, BedNum: 2},
		{TypeID: 2, Name: "單人間", Price: 150, BedNum: 1},
		{TypeID: 3, Name: "大床房", Price: 280, BedNum: 1},
		{TypeID: 4, Name: "三人間", Price: 300, BedNum: 3},
		{TypeID: 5, Name: "標準套房", Price: 380, BedNum: 2},
		{TypeID: 6, Name: "豪華套房", Price: 580, BedNum: 2},
		{TypeID: 7, Name: "商務套房", Price: 880, BedNum: 3},
		{TypeID: 8, Name: "行政套房", Price: 1000, BedNum: 3},
		{TypeID: 9, Name: "總統套房", Price: 2000, BedNum: 4},
	}

	model.DB.Create(&types)

}
