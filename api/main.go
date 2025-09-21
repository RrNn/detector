package main

import (
	"log"

	"detector/app"
	"detector/database"
	"detector/routes"
	"github.com/joho/godotenv"
)

func loadEnvironment() {
	if error := godotenv.Load(); error != nil {
		log.Fatal("Error loading .env file", error)
	}
}

func main() {
	loadEnvironment()
	database.Migrate()
	routes.AttachRoutes()

	_ = app.Start()
}
