package srv

import (
	"jlgre/classManager/srv/controllers/class"
	"jlgre/classManager/srv/controllers/note"
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

	router.GET("/classes", class.Index)
	router.GET("/class/:id", class.Get)
	router.POST("/classes", class.Create)
	router.PUT("/class/:id", class.Update)
	router.DELETE("/class/:id", class.Delete)

	router.GET("/notes", note.Index)
	router.GET("/note/:id", note.Get)
	router.POST("/notes", note.Create)
	router.PUT("/note/:id", note.Update)
	router.DELETE("/note/:id", note.Delete)

	router.Run(":8080")
}
