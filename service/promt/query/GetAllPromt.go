package query

import (
	"smart-edu-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Getpromt(c *fiber.Ctx) error {
	promt, err := utils.GetPromt()
	if err != nil {
		logrus.Error("Error while getting promt: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    promt,
		"message": "Success",
	})
}
