package models

import (
	"crypto/rand"
	"fmt"
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
	ID        uint   `json:"id"`
	Name      string `gorm:"not null" binding:"required" json:"name"`
	Code      string `gorm:"unqiue;not null" json:"code"`
	Notes     []Note
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (cls *Class) GenCode() error {
	bytes := make([]byte, 3)

	_, err := rand.Read(bytes)

	s := fmt.Sprintf("%X", bytes)

	cls.Code = s

	return err
}

type Note struct {
	ID        uint
	Name      string
	Markdown  string
	ClassID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
