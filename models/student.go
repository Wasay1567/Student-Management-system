package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name        string       `json:"name"`
	Email       string       `json:"email" gorm:"unique"`
	Age         uint         `json:"age"`
	Country     string       `json:"country"`
	Enrollments []Enrollment `gorm:"foreignKey:StudentID"`
}
