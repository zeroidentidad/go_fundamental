package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-apidoc/service"
)

type HandlersCliente struct {
	service service.ServiceCliente
}

func (hc *HandlersCliente) getAllClientes(c *fiber.Ctx) error {
	//status := c.Query("status")
	clientes, err := hc.service.GetAll() //status
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"status": false,
			"error":  err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": true,
		"data":   clientes,
	})
}
