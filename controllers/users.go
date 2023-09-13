package controllers

import (
	"get-api-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(context *gin.Context) {
	var input models.CheckDataValid
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Status: input.Status}
	models.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"users": user})
}