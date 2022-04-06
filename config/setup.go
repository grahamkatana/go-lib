package config

import (
	"grahamkatana/book-libray/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "libray.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{}, &models.User{}, &models.Rental{})

	DB = database
}
