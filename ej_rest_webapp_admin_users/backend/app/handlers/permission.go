package handlers

import (
	"backend/dto"
	"backend/service"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HandlerPermission struct {
	Svc service.PermissionService
}

func (h *HandlerPermission) IsAuthorized(c *fiber.Ctx, page string) error {
	usr := c.Locals("user")
	role_id := usr.(*dto.UserClaims).ResponseUser.RoleID

	role, err := h.Svc.RolePermissions(role_id)
	if err != nil {
		return errors.New("unauthorized error")
	}

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(http.StatusUnauthorized)
	return errors.New("unauthorized")
}

func (h *HandlerPermission) Permissions(c *fiber.Ctx) error {
	permissions, err := h.Svc.Permissions()

	return resJSON(c, permissions, err, http.StatusOK)
}
