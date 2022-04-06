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
