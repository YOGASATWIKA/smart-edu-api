package query

import (
	"smart-edu-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetOutlineById(c *fiber.Ctx) error {
	id := c.Params("id") // ambil ID dari path parameter
	existing, err := utils.GetOutlineById(id)
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