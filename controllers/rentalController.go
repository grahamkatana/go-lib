package controllers

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookRentalStruct struct {
	UserID uint64 `form:"user_id" binding:"required"`
	BookID uint64 `form:"book_id" bindig:"required"`
	// IsReturned    bool    `form:"is_returned" bindig:"required"`
	PayableAmount float64 `form:"payable_amount" binding:"required"`
	// LateReturnFee float64 `form:"late_return_fee" bindig:"required"`
	ReturnDate string `form:"return_date" binding:"required"`
}

type UpdateBookRentalStruct struct {
	// UserID uint64 `form:"user_id" binding:"required"`
	// BookID uint64 `form:"book_id" binding:"required"`
	// IsReturned    bool    `form:"is_returned" bindig:"required"`
	// PayableAmount float64 `form:"payable_amount" bindig:"required"`
	LateReturnFee float64 `form:"late_return_fee" binding:"required"`
	ReturnDate    string  `form:"return_date" binding:"required"`
}

func CreateBookRental(c *gin.Context) {
	var input CreateBookRentalStruct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var book models.Book
	var user models.User
	if err := config.DB.Where("id = ?", int(input.BookID)).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	book.Quantity = book.Quantity - 1
	config.DB.Model(&book).Updates(book)
	if err := config.DB.Where("id = ?", int(input.UserID)).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
	bookRental := models.Rental{
		UserID:        int(input.UserID),
		BookID:        int(input.BookID),
		IsReturned:    false,
		PayableAmount: input.PayableAmount,
		LateReturnFee: 0.0,
		ReturnDate:    input.ReturnDate,
	}
	result := config.DB.Create(&bookRental)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookRental})
}

func UpdateBookRental(c *gin.Context) {
	var bookRental models.Rental
	if err := config.DB.Where("id = ?", c.Param("id")).First(&bookRental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var book models.Book
	if err := config.DB.Where("id = ?", bookRental.BookID).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	book.Quantity = book.Quantity + 1
	config.DB.Model(&book).Updates(book)
	// Validate input
	var input UpdateBookRentalStruct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&bookRental).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": bookRental})
}
