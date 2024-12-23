package migration

import (
	"belajar-crud-go/configs"
	"belajar-crud-go/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Categories{})
	configs.DB.AutoMigrate(&models.User{})
	configs.DB.AutoMigrate(&models.Book{})
}
