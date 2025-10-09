package query

import (
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
)

func GetEbookModuleById(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := helper.GetContext()
	existing, err := repository.GetEbookByModulId(ctx, id)
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
