package initialisers

import "go-jwt/models"

func SyncDatabase() {
	Postgres_db.AutoMigrate(&models.User{})
}
