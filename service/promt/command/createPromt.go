package service

import (
	"smart-edu-api/data/promt/request"
	"smart-edu-api/model"
	"smart-edu-api/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)


func CreatePromt (app *fiber.Ctx) error{
	request := new(request.CreatePromtRequest)
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

	promt , errCreatePromt := utils.CreatePromt(model.Promt{
		Nama: request.Nama,
		Model: request.Model,
		Status: request.Status,
		PromtContext: request.PromtContext,
		PromtInstruction: request.PromtInstruction,
		CreatedAt: utils.GetCurrentTime(),
	})
	if errCreatePromt != nil {
		logrus.Printf("Terjadi error: %s\n", errCreatePromt.Error())
		return app.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "server error",
			})
	}

	return app.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":   promt,
			"message": "Promt created successfully",
		},
	)
}