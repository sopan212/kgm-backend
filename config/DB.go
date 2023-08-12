package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kgm-backend/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/kgm-app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to conect to database")
	}
	DB.AutoMigrate(&models.Books{})
}
