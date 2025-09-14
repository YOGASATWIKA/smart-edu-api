package service

import (
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func DeleteMateri(c *fiber.Ctx) error {
	id := c.Params("id")
	existing, err := repository.GetMateriPokokByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	existing.Status = "DELETED"
	existing.DeleteAt = helper.GetCurrentTime()

	_, err = repository.DeleteMateri(existing)
	if err != nil {
		logrus.Printf("Gagal soft delete base materi: %s\n", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus base materi",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Base materi berhasil dihapus (soft delete)",
	})
}
