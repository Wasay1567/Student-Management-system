package controllers

import (
	"net/http"
	"strconv"

	"github.com/Wasay1567/edutrack/config"
	"github.com/Wasay1567/edutrack/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterStudent(c *gin.Context) {
	var std models.RegStudent
	if err := c.ShouldBindJSON(&std); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var existingStudent models.Student
	if err := config.DB.Where("email = ?", std.Email).First(&existingStudent).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	student := models.Student{
		Name:    std.Name,
		Email:   std.Email,
		Age:     std.Age,
		Country: std.Country,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create account. Try again later"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "student": student})
}

func GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch student"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"student": student})
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	if err := config.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "students": students})
}

func UpdateStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData models.RegStudent
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch student"})
		}
		return
	}

	student.Name = updateData.Name
	student.Email = updateData.Email
	student.Age = updateData.Age
	student.Country = updateData.Country

	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "student": student})
}

func DeleteStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch student"})
		}
		return
	}

	if err := config.DB.Delete(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Student deleted"})
}
