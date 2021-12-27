package class

import (
	"errors"
	"jlgre/classManager/srv/db"
	"jlgre/classManager/srv/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func Index(context *gin.Context) {
	var classes []models.Class
	err := db.Connection.Find(&classes).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"classes": classes,
	})
}

func Create(context *gin.Context) {
	var class models.Class
	var mysqlerr *mysql.MySQLError

	if err := context.ShouldBindJSON(&class); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for {
		err := class.GenCode()

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result := db.Connection.Create(&class)

		if result.Error == nil {
			context.JSON(http.StatusCreated, class)
			return
		} else if errors.As(result.Error, &mysqlerr) && mysqlerr.Number != 1062 {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func Get(context *gin.Context) {
	var class models.Class

	result := db.Connection.First(&class, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		context.JSON(http.StatusOK, class)
	}
}

func Update(context *gin.Context) {
	var class models.Class

	result := db.Connection.First(&class, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if err := context.ShouldBindJSON(&class); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = db.Connection.Save(&class)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusOK, class)
	}
}

func Delete(context *gin.Context) {
	result := db.Connection.Delete(&models.Class{}, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.Status(http.StatusNoContent)
	}
}
