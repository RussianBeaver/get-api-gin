package models

type User struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
}

type CheckDataValid struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}