package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-apidoc/dto"
	"github.com/zeroidentidad/fiber-hex-apidoc/service"
)

type HandlersCuenta struct {
	service service.ServiceCuenta
}

func (hc *HandlersCuenta) PostNewCuenta(c *fiber.Ctx) error {
	cliente_id := c.Params("id")
	body := new(dto.RequestCuenta)
	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}

	body.ClienteID = cliente_id
	cuenta, err := hc.service.PostNew(*body)
	return resJSON(cuenta, err, c)
}
