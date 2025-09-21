package routes

import (
	"net/http"

	"detector/app"
	_ "detector/app" // blank
	"detector/constants"
	"detector/controller"
	"detector/database"
	"detector/jobs"
	_ "detector/jobs" // blank
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

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "welcome"})
	})
	app.POST("/register", c.Register)
	auth := app.Group("/auth")
	auth.POST("/login", c.Login)

	v1 := app.Group("/api/v1")
	v1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(env.JwtString),
		TokenLookup: "header:Authorization",
		ContextKey:  "token",
	}))

	v1.POST("/url", c.AddURL)
	v1.GET("/urls", c.GetUrls)
	v1.GET("/urls/:id", c.GetURL)
	v1.GET("/urls/search", c.SearchURL)
	v1.DELETE("/urls/:id", c.DeleteURL)

	return err
}
