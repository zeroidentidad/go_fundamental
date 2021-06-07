package handlers

import (
	"backend/dto"
	"backend/errs"
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func parseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func resJSON(c *fiber.Ctx, data interface{}, err *errs.AppError, status int) error {
	if err != nil {
		c.Status(err.Code)
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	c.Status(status)
	return c.JSON(data)
}

func setCookie(c *fiber.Ctx, jwt string, duration time.Duration) *fiber.Map {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Now().Add(duration),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return &fiber.Map{
		"message": "success",
	}
}

func CreateFile(filePath string, orders *[]dto.ResponseOrder) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// header
	_ = writer.Write([]string{"ID", "Name", "Email", "Product Title", "Price", "Quantity"})

	for _, order := range *orders {
		// body
		data := []string{
			strconv.Itoa(int(order.ID)),
			order.Name,
			order.Email,
			"",
			"",
			"",
		}
		if err := writer.Write(data); err != nil {
			return err
		}

		// detail
		for _, item := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				item.ProductTitle,
				strconv.Itoa(int(item.Price)),
				strconv.Itoa(int(item.Quantity)),
			}
			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}

	return nil
}
