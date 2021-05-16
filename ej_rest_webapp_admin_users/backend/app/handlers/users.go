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
	if err := parseBody(c, body); err != nil {
		return err
	}

	user, err := h.Svc.Register(*body)
	return resJSON(c, user, err, http.StatusCreated)
}

func (h *HandlerUser) Login(c *fiber.Ctx) error {
	var login fiber.Map
	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}

	jwt, err := h.Svc.Login(*body)
	if err == nil {
		login = *setCookie(c, jwt.Token, dto.TOKEN_DURATION)
	}

	return resJSON(c, login, err, http.StatusCreated)
}

func (h *HandlerUser) Logout(c *fiber.Ctx) error {
	logout := *setCookie(c, "", -dto.TOKEN_DURATION)

	return resJSON(c, logout, nil, http.StatusOK)
}

func (h *HandlerUser) AuthUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims, err := h.Svc.AuthUser(cookie)

	return resJSON(c, claims, err, http.StatusOK)
}
