package main

import (
	"log"

	"github.com/danikg/go-todo-rest-api/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	api := app.NewApp()
	api.Run()
}
