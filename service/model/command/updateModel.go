package service

import (
	"smart-edu-api/data/model/request"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func UpdateModel(app *fiber.Ctx) error {
	id := app.Params("id") // ambil ID dari path parameter

	// Cek apakah data dengan ID tersebut ada
	existing, err := repository.GetModelById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Model tidak ditemukan",
		})
	}

	// Parse dan validasi body request
	request := new(request.UpdateModelRequest)
	if err := app.BodyParser(request); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Invalid request body",
		})
	}

	isValid, err := govalidator.ValidateStruct(request)
	if !isValid && err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	existing.Model = request.Model
	existing.PromtContext = request.PromtContext
	existing.PromtInstruction = request.PromtInstruction
	existing.UpdatedAt = helper.GetCurrentTime()

	// Simpan perubahan
	updated, err := repository.UpdateModel(existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Gagal mengupdate data",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Model berhasil diperbarui",
		"data":    updated,
	})
}
