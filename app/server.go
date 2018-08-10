package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/wrbss/learning-go-echo/app/controllers"
)

func main() {
	e:= echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.NewExample().GetExample)

	e.Start(":5000")
}
