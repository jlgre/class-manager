package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint
	FirstName    string
	LastName     string
	UserName     string
	Password     string
	GlobalAdmin  bool
	Classes      []Class `gorm:"many2many:user_classes;"`
	AdminClasses []Class `gorm:"many2many:user_admin_classes;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Class struct {
	gorm.Model
	ID        uint
	Name      string
	Code      string
	Notes     []Note
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Note struct {
	ID        uint
	Name      string
	Markdown  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
