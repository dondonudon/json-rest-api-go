package main

import (
	"github.com/dondonudon/json-rest-api/controllers"
	"github.com/dondonudon/json-rest-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.PostsGetAll)
	r.GET("/posts/:id", controllers.PostsGet)
	r.POST("/posts", controllers.PostsCreate)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080

}
