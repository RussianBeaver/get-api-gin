package main
//   docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=usersDB -d postgres

// https://medium.com/@com.berdin/%D1%81%D1%82%D1%80%D0%BE%D0%B8%D0%BC-%D0%BF%D1%80%D0%BE%D1%81%D1%82%D0%BE%D0%B9-golang-restful-api-%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80-c-gin-postgres-%D0%B8-gorm-e76ac21c275e


import (
	"get-api-gin/models"
	"get-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
   // Подключение к базе данных
	models.ConnectDatabase()
   // Маршруты
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.GetUser)
	router.PATCH("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
   // Запуск сервера
	router.Run("localhost:8080")
}
