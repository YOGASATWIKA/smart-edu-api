package query

import (
	modul "smart-edu-api/data/modul/response"
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetActivity(c *fiber.Ctx) error {
	baseMateri, err := repository.GetActivity()
	if err != nil {
		logrus.Error("Error while getting base materi: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}
	responses := make([]modul.GetActivity, 0, len(baseMateri))
	for _, result := range baseMateri {
		response := modul.GetActivity{
			ID:          result.ID,
			Namajabatan: result.MateriPokok.Namajabatan,
			Status:      result.Status,
			State:       result.State,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    responses,
		"message": "Success",
	})
}
