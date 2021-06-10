package app

import (
	"backend/app/handlers"
	"backend/domain"
	"backend/service"

	"github.com/gofiber/fiber/v2"
)

func Routes(router *fiber.App) *fiber.App {
	db := dbclient()
	userStorage := domain.NewUserStorageDb(db)
	user := handlers.User{Svc: service.NewUserService(userStorage)}
	roleStorage := domain.NewRoleStorageDb(db)
	role := handlers.Role{Svc: service.NewRoleService(roleStorage)}
	permissionStorage := domain.NewPermissionStorageDb(db)
	permission := handlers.Permission{Svc: service.NewPermissionService(permissionStorage)}
	productStorage := domain.NewProductStorageDb(db)
	product := handlers.Product{Svc: service.NewProductService(productStorage)}
	orderStorage := domain.NewOrderStorageDb(db)
	order := handlers.Order{Svc: service.NewOrderService(orderStorage)}

	router.Static("/", "../frontend/build")

	api := router.Group("/api")

	api.Post("/register", user.Register)
	api.Post("/login", user.Login)
	api.Static("/uploads", "./uploads")

	api.Use(user.AuthUser)

	api.Post("/logout", user.Logout)
	api.Get("/user", user.User)
	api.Put("/update-profile", user.UpdateProfile)
	api.Put("/update-password", user.UpdatePassword)

	api.Use(func(c *fiber.Ctx) error {
		if err := permission.IsAuthorized(c, "users"); err != nil {
			return err
		}
		return c.Next()
	})
	api.Get("/users", user.Users)
	api.Post("/create-user", user.CreateUser)
	api.Get("/get-user/:id", user.GetUser)
	api.Put("/update-user", user.UpdateUser)
	api.Delete("/delete-user/:id", user.DeleteUser)

	api.Use(func(c *fiber.Ctx) error {
		if err := permission.IsAuthorized(c, "roles"); err != nil {
			return err
		}
		return c.Next()
	})
	api.Get("/roles", role.Roles)
	api.Post("/create-role", role.CreateRole)
	api.Get("/get-role/:id", role.GetRole)
	api.Put("/update-role", role.UpdateRole)
	api.Delete("/delete-role/:id", role.DeleteRole)

	api.Get("/permissions", permission.Permissions)

	api.Use(func(c *fiber.Ctx) error {
		if err := permission.IsAuthorized(c, "products"); err != nil {
			return err
		}
		return c.Next()
	})
	api.Get("/products", product.Products)
	api.Post("/create-product", product.CreateProduct)
	api.Get("/get-product/:id", product.GetProduct)
	api.Put("/update-product", product.UpdateProduct)
	api.Delete("/delete-product/:id", product.DeleteProduct)
	api.Post("/upload-image", product.UploadImage)

	api.Use(func(c *fiber.Ctx) error {
		if err := permission.IsAuthorized(c, "orders"); err != nil {
			return err
		}
		return c.Next()
	})
	api.Get("/orders", order.Orders)
	api.Post("/export-orders", order.ExportOrders)
	api.Get("/chart-sales", order.ChartSales)

	return router
}
