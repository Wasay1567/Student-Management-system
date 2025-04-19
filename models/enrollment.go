package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseID  uint    `json:"course_id"`
	Course    Course  `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status    string  `json:"status" default:"enrolled"`
}
