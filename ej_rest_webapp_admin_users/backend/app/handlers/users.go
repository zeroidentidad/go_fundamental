package handlers

import (
	"backend/domain"

	"github.com/gofiber/fiber/v2"
)

type HandlerUser struct {
}

func (h *HandlerUser) Register(c *fiber.Ctx) error {
	user := domain.User{
		ID:        1,
		FirstName: "testn name",
		LastName:  "test last",
		Email:     "test email",
		Password:  "test pass",
	}

	return c.JSON(user)
}
