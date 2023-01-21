package main

import (
	"hotel/cmd/seed/mock"
	"hotel/model"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../../.env")
	model.Database(os.Getenv("MYSQL_DSN"))
	mock.MockUser()
}
