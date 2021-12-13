package srv

import (
	"fmt"
	"jlgre/classManager/srv/db"
)

func Setup() {
	db.Connect()
}

func RegisterRoutes() {
	fmt.Println("Registered Routes")
}

func Srv() {
	fmt.Println("Serving application")
}
