package handlers

import (
	"backend/dto"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Svc service.UserService
}

func (h *User) AuthUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims, err := h.Svc.AuthUser(cookie)
	if err != nil {
		return resJSON(c, claims, err, http.StatusOK)
	}
	c.Locals("user", claims)

	return c.Next()
}

func (h *User) Register(c *fiber.Ctx) error {
	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}

	user, err := h.Svc.Register(*body)
	return resJSON(c, user, err, http.StatusCreated)
}

func (h *User) Login(c *fiber.Ctx) error {
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

func (h *User) Logout(c *fiber.Ctx) error {
	logout := *setCookie(c, "", -dto.TOKEN_DURATION)

	return resJSON(c, logout, nil, http.StatusOK)
}

func (h *User) User(c *fiber.Ctx) error {
	usr := c.Locals("user")

	user, err := h.Svc.User(usr.(*dto.UserClaims))

	return resJSON(c, user, err, http.StatusOK)
}

func (h *User) Users(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	users, total, err := h.Svc.Users(page)

	res := fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage(total),
		},
	}

	return resJSON(c, res, err, http.StatusOK)
}

func (h *User) CreateUser(c *fiber.Ctx) error {
	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}

	user, err := h.Svc.CreateUser(*body)
	return resJSON(c, user, err, http.StatusCreated)
}

func (h *User) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.Svc.GetUser(id)

	return resJSON(c, user, err, http.StatusOK)
}

func (h *User) UpdateUser(c *fiber.Ctx) error {
	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}

	user, err := h.Svc.UpdateUser(*body)
	return resJSON(c, user, err, http.StatusCreated)
}

func (h *User) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Svc.DeleteUser(id)

	return resJSON(c, "deleted", err, http.StatusOK)
}

func (h *User) UpdateProfile(c *fiber.Ctx) error {
	usr := c.Locals("user")

	id := usr.(*dto.UserClaims).ResponseUser.ID

	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}
	body.ID = id

	user, err := h.Svc.UpdateProfile(*body)
	return resJSON(c, user, err, http.StatusCreated)
}

func (h *User) UpdatePassword(c *fiber.Ctx) error {
	usr := c.Locals("user")

	id := usr.(*dto.UserClaims).ResponseUser.ID

	body := new(dto.RequestUser)
	if err := parseBody(c, body); err != nil {
		return err
	}
	body.ID = id

	user, err := h.Svc.UpdatePassword(*body)
	return resJSON(c, user, err, http.StatusCreated)
}
