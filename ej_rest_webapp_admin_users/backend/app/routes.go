package app

import (
	"backend/app/handlers"
	"backend/domain"
	"backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes(prefix string) *fiber.App {
	router := fiber.New()
	router.Use(logger.New())

	db := dbclient()
	userStorage := domain.NewUserStorageDb(db)
	hu := handlers.HandlerUser{Svc: service.NewUserService(userStorage)}

	router.Post(prefix+"/register", hu.Register)

	return router
}
