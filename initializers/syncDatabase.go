package initializers

import "go-jwt/models"

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
