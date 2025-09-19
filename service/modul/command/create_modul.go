package service

import (
	"smart-edu-api/data/baseMateri/request"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func CreateModul(c *fiber.Ctx) error {
	request := new(request.MateriPokokRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Invalid request body",
		})
	}
	isValid, err := govalidator.ValidateStruct(request)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}
	materiPokok, err := repository.CreateModul(entity.Modul{
		MateriPokok: entity.MateriPokok{
			Namajabatan:  request.Namajabatan,
			Tugasjabatan: request.Tugasjabatan,
			Keterampilan: request.Keterampilan,
		},
		Status:    "ACTIVE",
		State:     "DRAFT",
		CreatedAt: helper.GetCurrentTime(),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "server error",
			})
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"message": "Modul Created Successfully",
			"data":    materiPokok,
		},
	)
}
