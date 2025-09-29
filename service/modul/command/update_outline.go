package service

import (
	modul "smart-edu-api/data/modul/request"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func UpdateOutline(app *fiber.Ctx) error {
	id := app.Params("id") // ambil ID dari path parameter

	// Cek apakah data dengan ID tersebut ada
	existing, err := repository.GetModulById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Model tidak ditemukan",
		})
	}

	request := new(modul.OutlineRequest)
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

	existing.Outline = *request.Outline

	updated, err := repository.UpdateModul(helper.GetContext(), existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Gagal mengupdate data",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Modul berhasil diperbarui",
		"data":    updated,
	})
}
