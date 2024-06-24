package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Firstname string
	Lastname string
	Email string
}

