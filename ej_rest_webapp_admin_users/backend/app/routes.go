package app

import (
	"backend/app/handlers"
	"backend/domain"
	"backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes() *fiber.App {
	db := dbclient()
	userStorage := domain.NewUserStorageDb(db)
	hu := handlers.HandlerUser{Svc: service.NewUserService(userStorage)}
	roleStorage := domain.NewRoleStorageDb(db)
	hr := handlers.HandlerRole{Svc: service.NewRoleService(roleStorage)}
	permissionStorage := domain.NewPermissionStorageDb(db)
	hp := handlers.HandlerPermission{Svc: service.NewPermissionService(permissionStorage)}

	router := fiber.New()
	router.Use(logger.New())
	router.Use(cors.New(cors.Config{AllowCredentials: true}))
	route := router.Group("/api")

	route.Post("/register", hu.Register)
	route.Post("/login", hu.Login)

	route.Use(hu.AuthUser)

	route.Post("/logout", hu.Logout)
	route.Get("/user", hu.User)
	route.Put("/update-profile", hu.UpdateProfile)
	route.Put("/update-password", hu.UpdatePassword)
	route.Get("/users", hu.Users)
	route.Post("/create-user", hu.CreateUser)
	route.Get("/get-user/:id", hu.GetUser)
	route.Put("/update-user", hu.UpdateUser)
	route.Delete("/delete-user/:id", hu.DeleteUser)

	route.Get("/roles", hr.Roles)
	route.Post("/create-role", hr.CreateRole)
	route.Get("/get-role/:id", hr.GetRole)
	route.Put("/update-role", hr.UpdateRole)
	route.Delete("/delete-role/:id", hr.DeleteRole)

	route.Get("/permissions", hp.Permissions)

	return router
}
