package user

import (
	"jlgre/classManager/srv/db"
	"jlgre/classManager/srv/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var users []models.User
	err := db.Connection.Find(&users).Error

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
