package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	StudentId uint   `json:"student_id"`
	CourseId  uint   `json:"course_id"`
	Status    string `json:"status" default:"enrolled"`
}
