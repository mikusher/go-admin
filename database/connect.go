package database

import (
	"go-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=mikusher password=mikusher dbname=go_admin port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connection to database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{})
}
