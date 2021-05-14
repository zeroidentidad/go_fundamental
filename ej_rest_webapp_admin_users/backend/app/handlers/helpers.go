package handlers

import (
	"backend/errs"

	"github.com/gofiber/fiber/v2"
)

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func resJSON(data interface{}, err *errs.AppError, c *fiber.Ctx, status int) error {
	if err != nil {
		c.Status(err.Code)
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	c.Status(status)
	return c.JSON(data)
}
