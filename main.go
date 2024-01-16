package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swim0000/echo_htmx_templ/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Static("/static", "static")

	e.GET("/signup", handlers.SignupFormHandler)
	e.POST("/signup", handlers.SignupHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
