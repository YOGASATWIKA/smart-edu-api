package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetModels(c *fiber.Ctx) error {
	model, err := repository.GetModels()
	if err != nil {
		logrus.Error("Error while getting model: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    model,
		"message": "Success",
	})
}
