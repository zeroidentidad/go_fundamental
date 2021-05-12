package app

import (
	"backend/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes() *fiber.App {
	router := fiber.New()
	router.Use(logger.New())

	hu := handlers.HandlerUser{}

	prefix := "/api"
	router.Post(prefix+"/register", hu.Register)

	return router
}
