package note

import (
	"jlgre/classManager/srv/db"
	"jlgre/classManager/srv/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var notes []models.Note
	err := db.Connection.Find(&notes).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func Create(context *gin.Context) {
	var note models.Note

	if err := context.ShouldBindJSON(&note); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Connection.Save(&note)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusCreated, note)
	}
}

func Get(context *gin.Context) {
	var note models.Note

	result := db.Connection.First(&note, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusOK, note)
	}
}

func Update(context *gin.Context) {
	var note models.Note

	result := db.Connection.First(&note, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if err := context.ShouldBindJSON(&note); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = db.Connection.Save(&note)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	} else {
		context.JSON(http.StatusOK, note)
	}
}

func Delete(context *gin.Context) {
	result := db.Connection.Delete(&models.Note{}, context.Param("id"))

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		context.Status(http.StatusNoContent)
	}
}
