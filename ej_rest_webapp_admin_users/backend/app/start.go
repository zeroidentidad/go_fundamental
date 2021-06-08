package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	config()
	router := fiber.New()
	router.Use(logger.New())
	router.Use(cors.New(cors.Config{AllowCredentials: true}))
	serve(Routes(router))
}
