package main

import (
	"gorm.io/gorm"
	
)

func init(){	
	initializers.ConnecctToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Book{})
}

