package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
)

func resJSON(data interface{}, err *errors.AppError, c *fiber.Ctx) error {
	if err != nil {
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	return c.JSON(&fiber.Map{
		"data": data,
	})
}
