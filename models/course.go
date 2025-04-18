package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Instructor  string       `json:"instructor"`
	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}
