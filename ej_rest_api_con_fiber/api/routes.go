package api

import (
	"github.com/gofiber/fiber"
)

func SetupMoviesRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}

func SetupUsersRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)
	grp.Get("/video", s.ServeVideo)

	grp.Use(jwtMiddleware(tokenKey)).Post("/wishlist", s.WishListHandler)
}
