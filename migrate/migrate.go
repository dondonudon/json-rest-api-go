package main

import (
	"github.com/dondonudon/json-rest-api/initializers"
	"github.com/dondonudon/json-rest-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
