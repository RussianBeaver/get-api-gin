package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
}

var users = []user{
	{ID: "1", Name: "Cringe Master", Status: "admin"},
	{ID: "2", Name: "John Doe", Status: "user"},
	{ID: "3", Name: "Julian Paul Assange", Status: "user"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/user_id_:id", getUserByID)
	router.GET("/:status", getUsersByStatus)
	router.POST("/users", addUser)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
	// context.IndentedJSON(http.StatusOK, users)
}
/*
The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
*/
func addUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, value := range users {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"messege": "user with this id not found"})
}

func getUsersByStatus(c *gin.Context) {
	status := c.Param("status")
	err := 1
	
	for _, value := range users {
		if value.Status == status {
			c.IndentedJSON(http.StatusOK, value)
			err = 0
		}
	}
	if err == 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messege": "user with this status not found"})
	}
}
