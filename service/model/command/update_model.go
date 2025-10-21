package model

import (
	request2 "smart-edu-api/data/model/request"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func UpdateModel(app *fiber.Ctx) error {
	id := app.Params("id")

	existing, err := repository.GetModelById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Model Not Found",
		})
	}

	request := new(request2.ModelOutlineRequest)
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
	existing.Description = request.Description
	existing.Steps = request.Steps
	existing.Variables = request.Variables
	existing.IsActive = request.IsActive
	existing.UpdatedAt = helper.GetCurrentTime()

	updated, err := repository.UpdateModel(existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Error Updating Model",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Model Successfully",
		"data":    updated,
	})
}
