package service

import (
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteModel(app *fiber.Ctx) error {
	id := app.Params("id")

	existing, err := repository.GetModelById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Promt tidak ditemukan",
		})
	}

	existing.Status = "DELETED"
	existing.DeleteAt = helper.GetCurrentTime()

	_, err = repository.UpdateModel(existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus model",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "model berhasil dihapus (soft delete)",
	})
}
