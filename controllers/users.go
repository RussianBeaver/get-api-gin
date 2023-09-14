package controllers

import (
	"get-api-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)
// GET /users
// Получаем список всех пользователей
func GetAllUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"users": users})
}
// POST /user
// Создание пользователя
func CreateUser(context *gin.Context) {
	var input models.CheckUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Status: input.Status}
	models.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"users": user})
}
// GET /user/:id
// Получение одного пользователя по ID
func GetUserByID(context *gin.Context) {
   // Проверяем имеется ли запись
   var user models.User
   if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": "user with this id not found"})
      return
   }

	context.JSON(http.StatusOK, gin.H{"users": user})
}
// PATCH /user/:id
// Изменениe информации
func UpdateUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user with this id not found"})
		return
	}

	var input models.UpdateUser
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Update(input)

	context.JSON(http.StatusOK, gin.H{"users": user})
}
// DELETE /user/:id
// Удаление пользователя по айди
func DeleteUser (context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user with this id not found"})
		return
	}
	
	models.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"users": true})
}
// GET /users/:name
//вывод пользователей по имени
func GetUsersByName(context *gin.Context) {
	// Проверяем имеется ли запись
	var users []models.User
	models.DB.Where("name = ?", context.Param("name")).Find(&users)

 	context.JSON(http.StatusOK, gin.H{"users": users})
}
// GET /users/:status
//вывод пользователей по статусу
func GetUsersByStatus(context *gin.Context) {
	// Проверяем имеется ли запись
	var users []models.User
	models.DB.Where("status = ?", context.Param("status")).Find(&users)

 	context.JSON(http.StatusOK, gin.H{"users": users})
}