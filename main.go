package main

import (
	"net/http"
	"simple-crud-api/controllers"
	"simple-crud-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.POST("/MCQ", controllers.McqsCreate)

	r.GET("/MCQs/:num/:num2", controllers.ReturnMcq)
	r.GET("/MCQ/:id", controllers.ReturnSingleMcq)
	r.GET("/MCQ-sub/:sub", controllers.ReturnSingleMcqwITHsUB)
	r.GET("/MCQ-subjectsArray", controllers.GetAllTypeOFMcqs)

	r.PUT("/MCQ/:id", controllers.Update)

	r.DELETE("/MCQ/:id", controllers.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080

}
