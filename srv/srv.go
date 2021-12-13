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

	router.Run(":8080")
}
