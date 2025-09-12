package service

import (
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func DeleteMateriPokok(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil data
	existing, err := repository.GetMateriPokokByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	// Set soft delete
	existing.Status = "DELETED"
	existing.DeleteAt = helper.GetCurrentTime()

	// Simpan perubahan
	_, err = repository.UpdateMateriPokok(existing)
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
