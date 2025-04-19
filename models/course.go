package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Instructor  string       `json:"instructor"`
	Members     int          `json:"members" gorm:"default:0"`
	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}

type RegCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
}
