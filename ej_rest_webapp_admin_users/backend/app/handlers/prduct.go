package handlers

import (
	"backend/dto"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HandlerProduct struct {
	Svc service.ProductService
}

func (h *HandlerProduct) Products(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	products, total, err := h.Svc.Products(page)

	res := fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total": total,
			"page":  page,
		},
	}

	return resJSON(c, res, err, http.StatusOK)
}

func (h *HandlerProduct) CreateProduct(c *fiber.Ctx) error {
	body := new(dto.RequestProduct)
	if err := parseBody(c, body); err != nil {
		return err
	}

	product, err := h.Svc.CreateProduct(*body)
	return resJSON(c, product, err, http.StatusCreated)
}

func (h *HandlerProduct) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.Svc.GetProduct(id)

	return resJSON(c, product, err, http.StatusOK)
}

func (h *HandlerProduct) UpdateProduct(c *fiber.Ctx) error {
	body := new(dto.RequestProduct)
	if err := parseBody(c, body); err != nil {
		return err
	}

	product, err := h.Svc.UpdateProduct(*body)
	return resJSON(c, product, err, http.StatusCreated)
}

func (h *HandlerProduct) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Svc.DeleteProduct(id)

	return resJSON(c, "deleted", err, http.StatusOK)
}
