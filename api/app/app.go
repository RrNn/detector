package app

import "github.com/labstack/echo"

var App = echo.New()

func Start() (err error) {
	App.Logger.Fatal(App.Start(":5555"))
	return err
}
