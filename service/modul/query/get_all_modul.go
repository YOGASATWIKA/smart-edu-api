package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetModul(c *fiber.Ctx) error {
	stateParam := c.Query("state")

	baseMateri, err := repository.GetAllModul(stateParam)
	if err != nil {
		logrus.Error("Error while getting module: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    baseMateri,
		"message": "Success",
	})
}
