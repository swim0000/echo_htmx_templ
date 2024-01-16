package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swim0000/echo_htmx_templ/handlers"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/signup", handlers.SignupFormHandler)
	e.POST("/signup", handlers.SignupHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
