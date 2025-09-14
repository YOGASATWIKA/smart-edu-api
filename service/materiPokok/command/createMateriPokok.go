package service

import (
	request "smart-edu-api/data/baseMateri/request"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func CreateMateriPokok(c *fiber.Ctx) error {

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

	existsName, errCheckName := helper.IsNameExists(request.Namajabatan)
	if errCheckName == nil && existsName {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama Jabatan sudah terdaftar, silakan edit data yang sudah ada.",
		})
	}

	materiPokok, err := repository.CreateMateriPokok(entity.Materi{
		Namajabatan:  request.Namajabatan,
		Tugasjabatan: request.Tugasjabatan,
		Keterampilan: request.Keterampilan,
		Klasifikasi:  request.Klasifikasi,
		Status:       "ACTIVE",
		Stage:        "MATERI_POKOK",
		CreatedAt:    helper.GetCurrentTime(),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "server error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    materiPokok,
			"message": "Materi pokok created successfully",
		},
	)
}
