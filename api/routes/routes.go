package routes

import (
  "net/http"

  "github.com/RrNn/detector/app"
  _ "github.com/RrNn/detector/app" // blank
  "github.com/RrNn/detector/constants"
  "github.com/RrNn/detector/controller"
  "github.com/RrNn/detector/database"
  "github.com/RrNn/detector/jobs"
  _ "github.com/RrNn/detector/jobs" // blank
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

var env = constants.GetEnvironment()

// AttachRoutes exported
func AttachRoutes() (err error) {
  c := &controller.Controller{
    DB: database.Connect(),
  }
  jobs.StartPinging(c)
  app := app.App
  // app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  //   AllowOrigins: []string{"http://127.0.0.1:8080", "http://localhost:8080"},
  //   AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
  // }))
  app.GET("/", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{"message": "welcome"})
  })
  app.POST("/register", c.Register)
  auth := app.Group("/auth")
  auth.POST("/login", c.Login)
  auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey:  []byte(env.JwtString),
    TokenLookup: "header:Authorization",
    ContextKey:  "token",
  }))

  auth.POST("/url", c.AddURL)
  auth.GET("/urls", c.GetUrls)
  auth.GET("/urls/:id", c.GetURL)
  auth.DELETE("/urls/:id", c.DeleteURL)

  return err
}
