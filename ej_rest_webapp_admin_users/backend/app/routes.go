package app

import (
	"backend/app/handlers"
	"backend/domain"
	"backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes(prefix string) *fiber.App {
	router := fiber.New()
	router.Use(logger.New())
	router.Use(cors.New(cors.Config{AllowCredentials: true}))

	db := dbclient()
	userStorage := domain.NewUserStorageDb(db)
	hu := handlers.HandlerUser{Svc: service.NewUserService(userStorage)}

	router.Post(prefix+"/register", hu.Register)
	router.Post(prefix+"/login", hu.Login)

	router.Use(hu.AuthUser)

	router.Post(prefix+"/logout", hu.Logout)
	router.Get(prefix+"/user", hu.User)
	router.Get(prefix+"/users", hu.Users)
	router.Post(prefix+"/create-user", hu.CreateUser)
	router.Get(prefix+"/get-user/:id", hu.GetUser)
	router.Put(prefix+"/update-user", hu.UpdateUser)
	router.Delete(prefix+"/delete-user/:id", hu.DeleteUser)

	return router
}
