package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User2 struct {
	ID        uint `gorm:"primary_key"`
	Age       int
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"uniqueIndex"`
}

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "go_postgres_demo"
)

var Postgres_db *gorm.DB

func ConnectToDatabase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User2{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	fmt.Println("Connected to database")
	Postgres_db = database
}

func CreateUser() {
	user := User2{FirstName: "Jafar", LastName: "Soltani", Age: 43, Email: "jafar@soltani.com"}
	Postgres_db.Create(&user)
}

func FetchAllUsers() {
	var users []User2
	Postgres_db.Find(&users)
	fmt.Println(users)
}
