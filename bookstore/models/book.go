package models

import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Price float64 `json:"price"`
	// Author *Author `json:"author"`
}