package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetEbook(c *fiber.Ctx) error {
	baseMateri, err := repository.GetAllEbook()
	if err != nil {
		logrus.Error("Error while getting base materi: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    baseMateri,
		"message": "Success",
	})
}
