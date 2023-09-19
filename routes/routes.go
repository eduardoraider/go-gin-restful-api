package routes

import (
	"github.com/eduardoraider/gin-api-rest/controllers"
	docs "github.com/eduardoraider/gin-api-rest/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	docs.SwaggerInfo.Title = "GO API Rest with Gin Swagger"
	docs.SwaggerInfo.Description = "Example of an application made in GO with GORM and using Gin Swagger to document APIs made with the Gin Framework"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")
	r.GET("/index", controllers.ShowIndexPage)
	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchByCPF)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.NoRoute(controllers.RouteNotFound)
	r.Run(":8000")
}
