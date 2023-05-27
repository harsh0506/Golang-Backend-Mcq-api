package main

import (
	"simple-crud-api/initializers"
	"simple-crud-api/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.MCQ{})
	initializers.DB.AutoMigrate(&models.LEC{})
}
