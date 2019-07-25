package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	eco := echo.New()

	eco.Use(middleware.Logger())
	eco.Use(middleware.Recover())

	// CORS config
	// con metodos permitidos
	eco.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	eco.Static("/", "public")

	log.Error(eco.Start(":9000"))
}
