package models

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=usersDB password=postgres sslmode=disable")

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database
	
  DB.AutoMigrate(&User{})
}