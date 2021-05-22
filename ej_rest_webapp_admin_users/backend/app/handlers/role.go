package handlers

import (
	"backend/dto"
	"backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HandlerRole struct {
	Svc service.RoleService
}

func (h *HandlerRole) Roles(c *fiber.Ctx) error {
	roles, err := h.Svc.Roles()

	return resJSON(c, roles, err, http.StatusOK)
}

func (h *HandlerRole) CreateRole(c *fiber.Ctx) error {
	body := new(dto.RequestRole)
	if err := parseBody(c, body); err != nil {
		return err
	}

	role, err := h.Svc.CreateRole(*body)
	return resJSON(c, role, err, http.StatusCreated)
}

func (h *HandlerRole) GetRole(c *fiber.Ctx) error {
	id := c.Params("id")

	role, err := h.Svc.GetRole(id)

	return resJSON(c, role, err, http.StatusOK)
}

func (h *HandlerRole) UpdateRole(c *fiber.Ctx) error {
	body := new(dto.RequestRole)
	if err := parseBody(c, body); err != nil {
		return err
	}

	role, err := h.Svc.UpdateRole(*body)
	return resJSON(c, role, err, http.StatusCreated)
}

func (h *HandlerRole) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Svc.DeleteRole(id)

	return resJSON(c, "deleted", err, http.StatusOK)
}
