package handlers

import (
	"backend/errs"
	"time"

	"github.com/gofiber/fiber/v2"
)

func parseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func resJSON(c *fiber.Ctx, data interface{}, err *errs.AppError, status int) error {
	if err != nil {
		c.Status(err.Code)
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	c.Status(status)
	return c.JSON(data)
}

func setCookie(c *fiber.Ctx, jwt string, duration time.Duration) *fiber.Map {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Now().Add(duration),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return &fiber.Map{
		"message": "success",
	}
}
