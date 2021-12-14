package user

import (
	"jlgre/classManager/srv/db"
	"jlgre/classManager/srv/hash"
	"jlgre/classManager/srv/models"

	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	AccessCode string `json:"access_code" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
	Password  string `json:"password"`
}

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

func Create(context *gin.Context) {
	var request CreateRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.AccessCode != os.Getenv("ACCESS_CODE") {
		context.AbortWithStatus(http.StatusUnauthorized)
	} else {
		hash, err := hash.Hash(request.Password)

		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		user := models.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			UserName:  request.UserName,
			Password:  hash,
		}

		result := db.Connection.Create(&user)

		if result.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		} else {
			context.JSON(http.StatusCreated, user)
		}
	}
}

func Get(context *gin.Context) {
	var user models.User

	result := db.Connection.First(&user, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func Update(context *gin.Context) {
	var request UpdateRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Connection.First(&user, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
	}
}

func Delete(context *gin.Context) {
	result := db.Connection.Delete(&models.User{}, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		context.Status(http.StatusNoContent)
	}
}
