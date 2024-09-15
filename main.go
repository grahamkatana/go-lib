package main

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/middleware"
	routes "grahamkatana/book-libray/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	config.ConnectDatabase()
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.BookRoutes(router)
	routes.ServiceFeeRoutes(router)
	routes.BookRentalRoutes(router)

	//routes.UserRoutes(router)
	//router.POST("/register", controllers.Register)
	router.Run("localhost:8000")
}
