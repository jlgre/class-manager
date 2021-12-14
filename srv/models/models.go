package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint      `json:"id"`
	FirstName    string    `gorm:"not null" binding:"required" json:"first_name"`
	LastName     string    `gorm:"not null" binding:"required" json:"last_name"`
	UserName     string    `gorm:"unique;not null" binding:"required" json:"user_name"`
	Password     string    `gorm:"not null" binding:"required" json:"password"`
	GlobalAdmin  bool      `gorm:"default:false" json:"admin"`
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
