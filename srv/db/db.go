package db

import (
	"fmt"
	"jlgre/classManager/srv/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect() {

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	db.AutoMigrate(&models.User{}, &models.Note{}, &models.Class{})
	Connection = db
}
