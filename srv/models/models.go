package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	UserName  string
	Password  string
}
