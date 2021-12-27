package user

import (
	"jlgre/classManager/srv/db"
	"jlgre/classManager/srv/hash"
	"jlgre/classManager/srv/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var users []models.User
	err := db.Connection.Find(&users).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func Create(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := hash.Hash(user.Password)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user.Password = hash

	result := db.Connection.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusCreated, user)
	}
}

func Get(context *gin.Context) {
	var user models.User

	result := db.Connection.First(&user, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func Update(context *gin.Context) {
	var user models.User

	result := db.Connection.First(&user, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		currentHash := user.Password

		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			if !hash.Compare(user.Password, currentHash) && user.Password != currentHash {
				newHash, err := hash.Hash(user.Password)

				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}

				user.Password = newHash
			} else {
				user.Password = currentHash
			}

			result := db.Connection.Save(&user)

			if result.Error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			} else {
				context.JSON(http.StatusOK, user)
			}
		}
	}
}

func Delete(context *gin.Context) {
	result := db.Connection.Delete(&models.User{}, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.Status(http.StatusNoContent)
	}
}
