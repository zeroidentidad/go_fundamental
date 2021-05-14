package handlers

import (
	"backend/dto"
	"backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HandlerUser struct {
	Svc service.UserService
}

func (h *HandlerUser) Register(c *fiber.Ctx) error {
	body := new(dto.RequestUser)
	if err := ParseBody(c, body); err != nil {
		return err
	}

	user, err := h.Svc.Register(*body)
	return resJSON(user, err, c, http.StatusCreated)
}
