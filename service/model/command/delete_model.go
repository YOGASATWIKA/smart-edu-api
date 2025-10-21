package model

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
			"message": "Model Not Found",
		})
	}

	existing.IsActive = false
	existing.DeleteAt = helper.GetCurrentTime()

	_, err = repository.UpdateModel(existing)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Deleting Model",
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Model Deleted Successfully",
	})
}
