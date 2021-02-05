package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zeroidentidad/fiber-hex-apidoc/domain"
	"github.com/zeroidentidad/fiber-hex-apidoc/service"
)

func Start() {
	router := fiber.New()
	router.Use(logger.New())

	//hc := HandlersCliente{service.NewServiceCliente(domain.NewStorageStubCliente())}
	hc := HandlersCliente{service.NewServiceCliente(domain.NewStorageDbCliente())}

	router.Get("/clientes", hc.getAllClientes)
	router.Get("/clientes/:id", hc.getCliente)

	_ = router.Listen(":3000")
}
