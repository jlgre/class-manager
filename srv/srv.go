package srv

import (
	"jlgre/classManager/srv/controllers/user"
	"jlgre/classManager/srv/db"

	"github.com/gin-gonic/gin"
)

func Setup() {
	db.Connect()
}

func RegisterRoutesAndServe() {
	router := gin.Default()

	router.GET("/users", user.Index)
	router.GET("/user/:id", user.Get)
	router.POST("/users", user.Create)
	router.PUT("/user/:id", user.Update)
	router.DELETE("/user/:id", user.Delete)

	router.Run(":8080")
}
