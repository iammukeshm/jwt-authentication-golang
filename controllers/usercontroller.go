package controllers

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
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

	context.JSON(201, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}