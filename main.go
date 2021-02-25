package main

import (
	"log"

	"github.com/RrNn/detector/app"
	"github.com/RrNn/detector/database"
	"github.com/RrNn/detector/routes"
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
