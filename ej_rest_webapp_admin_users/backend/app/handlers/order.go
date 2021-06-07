package handlers

import (
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HandlerOrder struct {
	Svc service.OrderService
}

func (h *HandlerOrder) Orders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	orders, total, err := h.Svc.Orders(page)

	res := fiber.Map{
		"data": orders,
		"meta": fiber.Map{
			"total": total,
			"page":  page,
		},
	}

	return resJSON(c, res, err, http.StatusOK)
}

func (h *HandlerOrder) ExportOrders(c *fiber.Ctx) error {
	filePath := "./download/orders.csv"

	orders, err := h.Svc.ExportOrders()
	if err != nil {
		return resJSON(c, nil, err, http.StatusOK)
	}

	if err := CreateFile(filePath, orders); err != nil {
		return err
	}

	return c.Download(filePath)
}

func (h *HandlerOrder) ChartSales(c *fiber.Ctx) error {
	sales, err := h.Svc.ChartSales()

	return resJSON(c, sales, err, http.StatusOK)
}
