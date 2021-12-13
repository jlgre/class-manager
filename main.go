package main

import (
	"jlgre/classManager/srv"
)

func main() {
	srv.Setup()
	srv.RegisterRoutesAndServe()
}
