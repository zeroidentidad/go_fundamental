package app

import "github.com/gofiber/fiber/v2"

func resJSON(data interface{}, err error, c *fiber.Ctx) error {
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"status": false,
			"error":  err,
		})
	}

	return c.JSON(&fiber.Map{
		"status": true,
		"data":   data,
	})
}
