package config

// import (
// 	"grahamkatana/book-libray/models"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
// 	database, err := gorm.Open("sqlite3", "libray.db")

// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	database.AutoMigrate(&models.Book{}, &models.User{}, &models.Rental{})

// 	DB = database
// }

import (
	"fmt"
	"grahamkatana/book-libray/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	driver := GetConfigValues("DATABASE_DRIVER")
	// database := ""
	// databaseUser := GetConfigValues("DATABASE_USER")
	// databaseHost := GetConfigValues("DATABASE_URL")
	// databasePort := GetConfigValues("DATABASE_PORT")
	// databasePassword := GetConfigValues("DATABASE_PASSWORD")
	switch dbdriver := driver; dbdriver {
	case "mysql":
		// DATABASE_DRIVER=mysql
		// DATABASE_NAME=gorm
		// DATABASE_URL=localhost
		// DATABASE_PORT=3306
		// DATABASE_PASSWORD=
		// DATABASE_USER=root

		fmt.Println("mysql.")

		// fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", driver)
		dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database!")
		}

		database.AutoMigrate(&models.Book{}, &models.User{}, &models.Rental{})
		DB = database

	case "pgsql":
		fmt.Println("pgsql.")

	default:
		// database, err := gorm.Open("sqlite3", "libray.db")

		// if err != nil {
		// 	panic("Failed to connect to database!")
		// }

		// database.AutoMigrate(&models.Book{}, &models.User{}, &models.Rental{})

		// DB = database

	}

}
