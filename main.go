package main

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/middleware"

	//"grahamkatana/book-libray/controllers"
	routes "grahamkatana/book-libray/routes"
	//"os"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	config.ConnectDatabase()
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.BookRoutes(router)
	//routes.UserRoutes(router)
	//router.POST("/register", controllers.Register)
	router.Run("localhost:8000")
}
