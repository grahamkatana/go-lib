package controllers

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/models"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookStruct struct {
	Title    string                `form:"title" binding:"required"`
	Author   string                `form:"author" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	Quantity uint                  `form:"quantity" binding:"required"`
}

type UpdateBookStruct struct {
	Title    string                `form:"title"`
	Author   string                `form:"author"`
	Image    *multipart.FileHeader `form:"image"`
	Quantity uint                  `form:"quantity"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookStruct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {

		c.String(500, " Error uploading file ")
		return
	}
	c.SaveUploadedFile(file, "./uploads/books/"+file.Filename)

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author, Image: input.Image.Filename, Quantity: input.Quantity}
	config.DB.Create(&book)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// Get /books
// Search a book
func SearchBook(c *gin.Context) {
	var books models.Book
	var searchQuery string = c.Param("key")
	if err := config.DB.Where("id = ?", searchQuery).Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Match not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})

}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book
	//db.Where("name LIKE ?", "%jin%").Find(&users)

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateBookStruct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
