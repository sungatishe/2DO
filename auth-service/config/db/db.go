package db

import (
	"auth-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func InitDb() {
	dsn := os.Getenv("dsn")
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db ")
	}

	Db.AutoMigrate(&models.User{})
}
