package handlers

import "github.com/gofiber/fiber/v2"

type HandlerUser struct {
}

func (h *HandlerUser) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
