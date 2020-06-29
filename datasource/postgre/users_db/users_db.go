package users_db

import (
	"database/sql"
	"fmt"
	usersSchema "github.com/RobertMaulana/x-user-service/schema"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var (
	Client *sql.DB
	Database *gorm.DB
)

func Connect() {
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	schema := os.Getenv("PG_DBNAME")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, schema)
	var err error
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil{
		fmt.Printf("err %#v", err)
		panic(err)
	}
	Database = db
	Client = db.DB()
	//autoCreate := os.Getenv("DB_AUTO_CREATE")
	if true {
		fmt.Println("Dropping and recreating all tables...")
		usersSchema.AutoMigrate(db)
		fmt.Println("All tables recreated successfully...")
	}
}