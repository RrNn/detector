package routes

import (
  "net/http"

  "github.com/RrNn/detector/app"
  _ "github.com/RrNn/detector/app"
  "github.com/RrNn/detector/constants"
  "github.com/RrNn/detector/controller"
  "github.com/RrNn/detector/database"
  "github.com/RrNn/detector/jobs"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func AttachRoutes() (err error) {
  c := &controller.Controller{
    DB: database.Connect(),
  }
  jobs.StartPinging(c)
  app := app.App
  app.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "welcome")
  })
  app.POST("/register", c.Register)
  auth := app.Group("/auth")
  auth.POST("/login", c.Login)
  auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey:  []byte(constants.JwtString),
    TokenLookup: "header:Authorization",
    ContextKey:  "token",
  }))

  auth.POST("/url", c.AddUrl)
  auth.GET("/urls", c.GetUrls)

  return err
}
