package main

import (
	"jlgre/classManager/srv"
)

func main() {
	srv.Setup()
	srv.RegisterRoutes()
	srv.Srv()
}
