package model

import (
	"smart-edu-api/data/model/request"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateModel(app *fiber.Ctx) error {
	request := new(request.ModelOutlineRequest)
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

	model, err := repository.CreateModel(entity.Model{
		ID:        primitive.NewObjectID(),
		Model:     request.Model,
		Promt:     entity.Promt(request.Promt),
		Type:      "OUTLINE",
		Status:    "ACTIVE",
		CreatedAt: helper.GetCurrentTime(),
	})
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "server error",
			})
	}

	return app.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    model,
			"message": "Model created successfully",
		},
	)
}
