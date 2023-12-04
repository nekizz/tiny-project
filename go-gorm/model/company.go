package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name          string
	ContactNumber string
	Description   string
	User          []User
}
