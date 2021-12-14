package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint
	FirstName    string    `gorm:"not null"`
	LastName     string    `gorm:"not null"`
	UserName     string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null"`
	GlobalAdmin  bool      `gorm:"default:false"`
	Classes      []Class   `gorm:"many2many:user_classes;"`
	AdminClasses []Class   `gorm:"many2many:user_admin_classes;"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
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
	ClassID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
