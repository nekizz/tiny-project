package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Gender    string
	Age       string
	Note      string
	CompanyID uint `gorm:"default:null"`
	Company   Company
}
