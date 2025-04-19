package controllers

import (
	"net/http"

	"github.com/Wasay1567/edutrack/models"
	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context){
	var course models.RegCourse
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid payload"})
		return
	}
	



}

func GetCourseByID(c *gin.Context){

}

func GetAllCourses(c *gin.Context){

}

func UpdateCourseByID(c *gin.Context){

}

func DeleteCourseByID(c *gin.Context){

}

func GetStudentsByCourseID(c *gin.Context){

}