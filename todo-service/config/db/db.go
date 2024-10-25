package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"todo-service/internal/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("dsn")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db ")
	}
	DB.AutoMigrate(&models.Todo{})
}
