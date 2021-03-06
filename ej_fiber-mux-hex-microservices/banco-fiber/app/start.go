package app

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zeroidentidad/fiber-hex-api/domain"
	logs "github.com/zeroidentidad/fiber-hex-api/logger"
	"github.com/zeroidentidad/fiber-hex-api/service"
)

func Start() {
	sanityCheck()
	router := fiber.New()
	router.Use(logger.New())

	dbClient := getDbClient()
	storageDbCliente := domain.NewStorageDbCliente(dbClient)
	storageDbCuenta := domain.NewStorageDbCuenta(dbClient)

	hclientes := HandlersCliente{service.NewServiceCliente(storageDbCliente)}
	hcuentas := HandlersCuenta{service.NewServiceCuenta(storageDbCuenta)}

	router.Get("/clientes", hclientes.getAllClientes)
	router.Get("/clientes/:id", hclientes.getCliente)
	router.Post("/clientes/:id/cuenta", hcuentas.postNewCuenta)
	router.Post("/clientes/:id/cuenta/:id_cuenta", hcuentas.postNewTransaccion)

	mauth := MiddlewareAuth{domain.NewStorageAuth()}
	router.Use(mauth.authorizationHandler())

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logs.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	err := router.Listen(fmt.Sprintf("%s:%s", address, port))
	logs.Fatal(err.Error())
}
