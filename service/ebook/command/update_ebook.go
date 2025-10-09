package command

import (
	"smart-edu-api/entity"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
)

func GetEbookById(c *fiber.Ctx) error {
	id := c.Params("id")
	ebook, err := repository.GetEbookById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ebook tidak ditemukan",
		})
	}
	return c.JSON(ebook)
}

func UpdateEbookById(c *fiber.Ctx) error {
	id := c.Params("id")

	var ebook entity.Ebook
	if err := c.BodyParser(&ebook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Request body tidak valid",
		})
	}

	err := repository.UpdateEbookById(c.Context(), id, ebook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memperbarui ebook",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Ebook berhasil diperbarui",
		"id":      id,
	})
}
