package handlers

import (
	"backend/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HandlerPermission struct {
	Svc service.PermissionService
}

func (h *HandlerPermission) Permissions(c *fiber.Ctx) error {
	permissions, err := h.Svc.Permissions()

	return resJSON(c, permissions, err, http.StatusOK)
}
