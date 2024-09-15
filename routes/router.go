package routes

import (
	"grahamkatana/book-libray/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)
}

func BookRoutes(route *gin.Engine) {
	route.GET("/books", controllers.GetBooks)
	route.POST("/books", controllers.CreateBook)
	route.GET("/books/:id", controllers.FindBook)
	route.PATCH("/books/:id", controllers.UpdateBook)
	route.DELETE("/books/:id", controllers.DeleteBook)
	route.GET("/books/search/:key", controllers.SearchBook)

}

func ServiceFeeRoutes(route *gin.Engine) {
	route.GET("/service-fees", controllers.GetServiceFees)
	route.POST("/service-fees", controllers.CreateServiceFee)
	route.GET("/service-fees/:id", controllers.FindServiceFee)
	route.PATCH("/service-fees/:id", controllers.UpdateServiceFee)
	route.DELETE("/service-fees/:id", controllers.DeleteServiceFee)
}

func BookRentalRoutes(route *gin.Engine) {
	route.POST("/rentals", controllers.CreateBookRental)
	route.PATCH("/rentals/:id", controllers.UpdateBookRental)
}
