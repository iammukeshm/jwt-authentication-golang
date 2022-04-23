package controllers

import (
	"jwt-authentication-golang/auth"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(500, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(401, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err:= auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"token": tokenString})
}