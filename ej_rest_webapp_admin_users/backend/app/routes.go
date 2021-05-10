package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes() *fiber.App {
	router := fiber.New()
	router.Use(logger.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return router
}
