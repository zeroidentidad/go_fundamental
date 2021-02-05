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
	return resJSON(clientes, err, c)
}

func (hc *HandlersCliente) getCliente(c *fiber.Ctx) error {
	id := c.Params("id")
	cliente, err := hc.service.GetById(id)
	return resJSON(cliente, err, c)
}
