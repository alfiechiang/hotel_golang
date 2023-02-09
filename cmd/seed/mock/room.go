package mock

import (
	"hotel/model"
	"math/rand"
	"time"
)

func MockRoom() {

	model.DB.Exec("TRUNCATE TABLE room")

	var roomtypes []model.RoomType
	model.DB.Table("room_type").Find(&roomtypes)

	insertData := make([]model.Room, 0)
	for _, v := range roomtypes {
		t1 := time.Now().UnixNano()
		t2 := time.Now().UnixNano()

		r1 := rand.New(rand.NewSource(t1))
		r2 := rand.New(rand.NewSource(t2 + 1))

		room1 := model.Room{}
		room1.TypeID = v.TypeID
		room1.RoomID = uint(r1.Int())
		room1.Status = true

		room2 := model.Room{}
		room2.TypeID = v.TypeID
		room2.RoomID = uint(r2.Int())
		room2.Status = true

		insertData = append(insertData, room1)
		insertData = append(insertData, room2)

	}

	model.DB.Create(&insertData)

}
