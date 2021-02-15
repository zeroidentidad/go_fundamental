package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-api/errors"
)

func resJSON(data interface{}, err *errors.AppError, c *fiber.Ctx) error {
	if err != nil {
		c.Status(err.Code)
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	c.Status(http.StatusCreated)
	return c.JSON(&fiber.Map{
		"data": data,
	})
}
