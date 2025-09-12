package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetModel(c *fiber.Ctx) error {
	model, err := repository.GetAllModel()
	if err != nil {
		logrus.Error("Error while getting promt: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    model,
		"message": "Success",
	})
}
