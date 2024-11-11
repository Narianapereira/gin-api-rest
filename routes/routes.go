package routes

import (
	"github.com/Narianapereira/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/student", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.GET("/student/:id", controllers.GetStudendById)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.PATCH("/student/:id", controllers.EditStudent)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.Run()
}
