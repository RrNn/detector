package main

import (
  "github.com/RrNn/detector/app"
  "github.com/RrNn/detector/database"
  "github.com/RrNn/detector/routes"
)

func main() {
  database.Migrate()
  routes.AttachRoutes()
  _ = app.Start()

}
