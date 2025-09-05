package service

import (
	request "smart-edu-api/data/baseMateri/request"
	"smart-edu-api/model"
	"smart-edu-api/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func CreateBaseMateri(c *fiber.Ctx) error {

	request := new(request.CreateBaseMateriRequest)
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

	existsName, errCheckName := utils.IsNameExists(request.Namajabatan)
	if errCheckName == nil && existsName {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama Jabatan sudah terdaftar, silakan edit data yang sudah ada.",
		})
	}

	baseMateri, errCreateBaseMateri := utils.CreateBaseMateri(model.Materi{
		Namajabatan:  request.Namajabatan,
		Tugasjabatan: request.Tugasjabatan,
		Keterampilan: request.Keterampilan,
		Klasifikasi:  request.Klasifikasi,
		Status:      "ACTIVE",
		CreatedAt:    utils.GetCurrentTime(),
	})
	if errCreateBaseMateri != nil {
		logrus.Printf("Terjadi error: %s\n", errCreateBaseMateri.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "server error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":   baseMateri,
			"message": "Base materi created successfully",
		},
	)
}

