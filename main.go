package main

import (
	"jlgre/classManager/srv"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal(envErr)
	}

	srv.Setup()
	srv.RegisterRoutesAndServe()
}
