package service

import (
	"smart-edu-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func DeletePromt(app *fiber.Ctx) error{
	id := app.Params("id")

	// Ambil data
	existing, err := utils.GetPromtById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Promt tidak ditemukan",
		})
	}

	// Set soft delete
	existing.Status = "DELETED"
	existing.DeleteAt = utils.GetCurrentTime()

	// Simpan perubahan
	_, err = utils.UpdatePromt(existing)
	if err != nil {
		logrus.Printf("Gagal soft delete promt: %s\n", err.Error())
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus promt",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "promt berhasil dihapus (soft delete)",
	})
}