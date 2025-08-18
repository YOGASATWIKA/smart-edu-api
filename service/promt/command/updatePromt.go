package service

import (
	"smart-edu-api/data/promt/request"
	"smart-edu-api/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)



func UpdatePromt (app *fiber.Ctx) error{
	id := app.Params("id") // ambil ID dari path parameter

	// Cek apakah data dengan ID tersebut ada
	existing, err := utils.GetPromtById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Promt tidak ditemukan",
		})
	}

	// Parse dan validasi body request
	request := new(request.UpdatePromtRequest)
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

	existing.Nama = request.Nama
	existing.Model = request.Model
	existing.PromtContext = request.PromtContext
	existing.PromtInstruction = request.PromtInstruction
	existing.UpdatedAt = utils.GetCurrentTime()


	// Simpan perubahan
	updated, err := utils.UpdatePromt(existing)
	if err != nil {
		logrus.Printf("Terjadi error saat update: %s\n", err.Error())
		return app.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Gagal mengupdate data",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Promt berhasil diperbarui",
		"data":    updated,
	})
}