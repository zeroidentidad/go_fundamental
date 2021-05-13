package app

import (
	"backend/app/handlers"
	"backend/domain"
	"backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes() *fiber.App {
	router := fiber.New()
	router.Use(logger.New())

	dbUsers := dbclient()
	userStorage := domain.NewUserStorageDb(dbUsers)

	hu := handlers.HandlerUser{Svc: service.NewUserService(userStorage)}

	prefix := "/api"
	router.Post(prefix+"/register", hu.Register)

	return router
}
