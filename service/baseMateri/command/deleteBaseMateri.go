package service

import (
	"smart-edu-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)


func DeleteBaseMateri(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil data
	existing, err := utils.GetBaseMateriByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	// Set soft delete
	existing.Status = "DELETED"
	existing.DeleteAt = utils.GetCurrentTime()

	// Simpan perubahan
	_, err = utils.UpdateBaseMateri(existing)
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

