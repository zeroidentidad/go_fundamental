package main

import (
	"net/http"
	"time"

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

	eco.GET("/json", getJSON)

	log.Error(eco.Start(":9000"))
}

func getJSON(c echo.Context) error {
	time.Sleep(3 * time.Second)
	ns := make([]Notas, 0)
	ns = append(ns, Notas{Titulo: "xyz 123", Contenido: "lorem ipsum not"})
	ns = append(ns, Notas{Titulo: "asd 556", Contenido: "lorem ipsum not"})
	ns = append(ns, Notas{Titulo: "bnn 987", Contenido: "lorem ipsum not"})

	return c.JSON(http.StatusOK, ns)
	//return c.JSON(http.StatusUnauthorized, ns)
	//return c.JSON(http.StatusNotFound, ns)
	//return c.JSON(http.StatusInternalServerError, ns)
}

type Notas struct {
	Titulo    string
	Contenido string
}
