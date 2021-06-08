package handlers

import (
	"backend/dto"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Svc service.ProductService
}

func (h *Product) Products(c *fiber.Ctx) error {
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

func (h *Product) CreateProduct(c *fiber.Ctx) error {
	body := new(dto.RequestProduct)
	if err := parseBody(c, body); err != nil {
		return err
	}

	product, err := h.Svc.CreateProduct(*body)
	return resJSON(c, product, err, http.StatusCreated)
}

func (h *Product) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.Svc.GetProduct(id)

	return resJSON(c, product, err, http.StatusOK)
}

func (h *Product) UpdateProduct(c *fiber.Ctx) error {
	body := new(dto.RequestProduct)
	if err := parseBody(c, body); err != nil {
		return err
	}

	product, err := h.Svc.UpdateProduct(*body)
	return resJSON(c, product, err, http.StatusCreated)
}

func (h *Product) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Svc.DeleteProduct(id)

	return resJSON(c, "deleted", err, http.StatusOK)
}

func (h *Product) UploadImage(c *fiber.Ctx) error {
	product_id := c.FormValue("product_id")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	var filename string
	for _, file := range files {
		filename = product_id + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	id, _ := strconv.Atoi(product_id)
	p := dto.RequestProduct{}
	p.ID = uint(id)
	p.Image = filename

	product, errs := h.Svc.UpdateProduct(p)

	upload := fiber.Map{
		"data": product,
		"url":  c.BaseURL() + "/api/uploads/" + filename,
	}

	return resJSON(c, upload, errs, http.StatusCreated)
}
