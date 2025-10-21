package service

import (
	"context"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteModul(app *fiber.Ctx) error {
	id := app.Params("id")
	ctx := context.Background()

	existing, err := repository.GetModulById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Modul tidak ditemukan",
		})
	}

	existing.IsActive = false
	existing.DeleteAt = helper.GetCurrentTime()

	_, err = repository.UpdateModul(ctx, existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus model",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "model berhasil dihapus (soft delete)",
	})
}
