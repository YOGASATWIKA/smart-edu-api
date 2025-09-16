package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
)

func GetMateriById(c *fiber.Ctx) error {
	id := c.Params("id")
	existing, err := repository.GetFullMateriById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Outline tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    existing,
		"message": "Success",
	})
}
