package initialisers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	//dbname = "go_jwt_demo"
)

var Postgres_db *gorm.DB

func ConnectToDatabase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, os.Getenv("DB"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	/*err = database.AutoMigrate(&User2{})
	if err != nil {
		panic("Failed to migrate database!")
	}*/

	fmt.Println("Connected to database")
	Postgres_db = database

}
