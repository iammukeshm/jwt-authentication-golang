package controllers

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(context *gin.Context) {
	var request LoginRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil || record.Error == gorm.ErrRecordNotFound {
		context.JSON(401, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

}

func Register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(500, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(201, user)
}
