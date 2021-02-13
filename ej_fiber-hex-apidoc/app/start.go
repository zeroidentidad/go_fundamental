package app

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/zeroidentidad/fiber-hex-apidoc/domain"
	logs "github.com/zeroidentidad/fiber-hex-apidoc/logger"
	"github.com/zeroidentidad/fiber-hex-apidoc/service"
)

func Start() {
	sanityCheck()
	router := fiber.New()
	router.Use(logger.New())

	//hc := HandlersCliente{service.NewServiceCliente(domain.NewStorageStubCliente())}
	hc := HandlersCliente{service.NewServiceCliente(domain.NewStorageDbCliente())}

	router.Get("/clientes", hc.getAllClientes)
	router.Get("/clientes/:id", hc.getCliente)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logs.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	err := router.Listen(fmt.Sprintf("%s:%s", address, port))
	logs.Fatal(err.Error())

}

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		logs.Fatal("Error loading .env file")
	}

	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logs.Fatal(fmt.Sprintf("Env variable %s not defined. Terminating app...", k))
		}
	}
}
