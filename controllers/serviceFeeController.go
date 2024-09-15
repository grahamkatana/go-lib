package controllers

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateServiceFeeStruct struct {
	LateServiceFee float64 `form:"late_service_fee" binding:"required"`
	InitiationFee  float64 `form:"initiation_fee" binding:"required"`
	LostFee        float64 `form:"lost_fee" binding:"required"`
}

type UpdateServiceFeeStruct struct {
	LateServiceFee float64 `form:"late_service_fee"`
	InitiationFee  float64 `form:"initiation_fee"`
	LostFee        float64 `form:"lost_fee"`
}

// POST /service-fee
// Create new service fee
func CreateServiceFee(c *gin.Context) {
	// Validate input
	var input CreateServiceFeeStruct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create service fee
	serviceFee := models.ServiceFee{
		LateServiceFee: input.LateServiceFee,
		InitiationFee:  input.InitiationFee,
		LostFee:        input.LostFee,
	}
	config.DB.Create(&serviceFee)
	c.JSON(http.StatusCreated, gin.H{"data": serviceFee})
}

// GET /service-fee
// Get all service fees
func GetServiceFees(c *gin.Context) {
	var serviceFees []models.ServiceFee
	config.DB.Find(&serviceFees)
	c.JSON(http.StatusOK, gin.H{"data": serviceFees})
}

// Get /service-fee/:id
// Find a service fee
func FindServiceFee(c *gin.Context) {
	var serviceFee models.ServiceFee

	if err := config.DB.Where("id = ?", c.Param("id")).First(&serviceFee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": serviceFee})
}

// PATCH /service-fee/:id
// Update a service fee
func UpdateServiceFee(c *gin.Context) {
	var serviceFee models.ServiceFee
	if err := config.DB.Where("id = ?", c.Param("id")).First(&serviceFee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input UpdateServiceFeeStruct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&serviceFee).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": serviceFee})
}

// DELETE /service-fee/:id
// Delete a service fee
func DeleteServiceFee(c *gin.Context) {
	var serviceFee models.ServiceFee
	if err := config.DB.Where("id = ?", c.Param("id")).First(&serviceFee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&serviceFee)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
