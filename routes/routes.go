package routes

import (
	"github.com/Wasay1567/edutrack/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	students := r.Group("/student")
	{
		students.POST("", controllers.RegisterStudent)
		students.GET("/:id", controllers.GetStudentByID)
		students.GET("", controllers.GetAllStudents)
		students.PUT("/:id", controllers.UpdateStudentByID)
		students.DELETE("/:id", controllers.DeleteStudentByID)
	}
}
