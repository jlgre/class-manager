package models

import (
	"crypto/rand"
	"fmt"
	"time"
)

type User struct {
	ID           uint      `gorm:"primary_key" json:"id"`
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
	ID        uint   `gorm:"primary_key" json:"id"`
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
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"not null" binding:"required" json:"name"`
	Markdown  string    `gorm:"not null" binding:"required" json:"markdown"`
	ClassID   uint      `binding:"required" json:"class_id"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
