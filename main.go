package main

import (
	"jlgre/classManager/srv"
)

func main() {
	srv.RegisterRoutes()
	srv.Srv()
}
