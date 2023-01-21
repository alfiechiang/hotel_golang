package mock

import "hotel/model"

func MockUser() {

	model.DB.Exec("TRUNCATE TABLE users")
	admin := model.User{UserName: "admin"}
	admin.SetPassword("123456")
	admin.Nickname = "超級管理員"
	admin.Status = true

	model.DB.Create(&admin)

}
