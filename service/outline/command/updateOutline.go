package service

import (
	"smart-edu-api/data/outline"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func UpdateOutline(app *fiber.Ctx) error {
	godotenv.Load()
	// ambil ID dari path parameter
	id := app.Params("id")
	// Check if outline exists
	currentOutline, err := repository.GetOutlineByMateriPokokId(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Outline tidak ditemukan",
		})
	}

	request := new(outline.Request)
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

	currentOutline.Outline = request.Outline
	currentOutline.UpdatedAt = helper.GetCurrentTime()

	updated, err := repository.UpdateOutline(currentOutline)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to save outline",
			"error":   err.Error(),
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Outline updated successfully",
		"jabatan": updated.MateriPokok.Namajabatan,
	})
}
