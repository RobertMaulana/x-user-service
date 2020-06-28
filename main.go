package main

import (
	"github.com/RobertMaulana/x-user-service/app"
	"github.com/RobertMaulana/x-user-service/datasource/postgre/users_db"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
	users_db.Connect()
}

func main() {
	app.StartApplication()
}
