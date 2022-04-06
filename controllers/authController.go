package controllers

import (
	"grahamkatana/book-libray/config"
	"grahamkatana/book-libray/models"

	helper "grahamkatana/book-libray/utils"

	"golang.org/x/crypto/bcrypt"

	//"mime/multipart"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterStruct struct {
	Username      string `form:"user_name" binding:"required"`
	Email         string `form:"email" binding:"required"`
	PassCode      string `form:"pin" binding:"required"`
	Cell          string `form:"cell" binding:"required"`
	Token         string `form:"token"`
	Refresh_token string `form:"refresh_token"`
}

type LoginStruct struct {
	Email    string `form:"email" binding:"required"`
	PassCode string `form:"pin" binding:"required"`
}

func HashPin(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func CheckPin(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func Register(c *gin.Context) {
	var input RegisterStruct
	var count int64
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user []models.User

	if err := config.DB.Where("email = ?", input.Email).Find(&user).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"data": "The email already exists"})
		return

	}
	if err := config.DB.Where("cell = ?", input.Cell).Find(&user).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"data": "The cell number already exists"})
		return

	}
	pin := HashPin(input.PassCode)
	input.PassCode = pin
	token, refreshToken, _ := helper.GenerateAllTokens(input.Email)
	input.Token = token
	input.Refresh_token = refreshToken
	userCreate := models.User{Username: input.Username, Email: input.Email, PassCode: input.PassCode, Cell: input.Cell, Token: input.Token, Refresh_token: input.Refresh_token}
	config.DB.Create(&userCreate)
	c.JSON(http.StatusCreated, gin.H{"data": userCreate})
}

func Login(c *gin.Context) {
	var input LoginStruct
	var user models.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email does not exist!"})
		return
	}

	pinIsValid := CheckPin(input.PassCode, user.PassCode)
	if pinIsValid {
		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"email":         user.Email,
			"cell":          user.Cell,
			"token":         user.Token,
			"refresh_token": user.Refresh_token,
		}})

	}
}
