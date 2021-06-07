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
	user := handlers.HandlerUser{Svc: service.NewUserService(userStorage)}
	roleStorage := domain.NewRoleStorageDb(db)
	role := handlers.HandlerRole{Svc: service.NewRoleService(roleStorage)}
	permissionStorage := domain.NewPermissionStorageDb(db)
	permission := handlers.HandlerPermission{Svc: service.NewPermissionService(permissionStorage)}
	productStorage := domain.NewProductStorageDb(db)
	product := handlers.HandlerProduct{Svc: service.NewProductService(productStorage)}
	orderStorage := domain.NewOrderStorageDb(db)
	order := handlers.HandlerOrder{Svc: service.NewOrderService(orderStorage)}

	router := fiber.New()
	router.Use(logger.New())
	router.Use(cors.New(cors.Config{AllowCredentials: true}))
	route := router.Group("/api")

	route.Post("/register", user.Register)
	route.Post("/login", user.Login)
	route.Static("/uploads", "./uploads")

	route.Use(user.AuthUser)

	route.Post("/logout", user.Logout)
	route.Get("/user", user.User)
	route.Put("/update-profile", user.UpdateProfile)
	route.Put("/update-password", user.UpdatePassword)
	route.Get("/users", user.Users)
	route.Post("/create-user", user.CreateUser)
	route.Get("/get-user/:id", user.GetUser)
	route.Put("/update-user", user.UpdateUser)
	route.Delete("/delete-user/:id", user.DeleteUser)

	route.Get("/roles", role.Roles)
	route.Post("/create-role", role.CreateRole)
	route.Get("/get-role/:id", role.GetRole)
	route.Put("/update-role", role.UpdateRole)
	route.Delete("/delete-role/:id", role.DeleteRole)

	route.Get("/permissions", permission.Permissions)

	route.Get("/products", product.Products)
	route.Post("/create-product", product.CreateProduct)
	route.Get("/get-product/:id", product.GetProduct)
	route.Put("/update-product", product.UpdateProduct)
	route.Delete("/delete-product/:id", product.DeleteProduct)
	route.Post("/upload-image", product.UploadImage)

	route.Get("/orders", order.Orders)
	route.Post("/export-orders", order.ExportOrders)
	route.Get("/chart-sales", order.ChartSales)

	return router
}
