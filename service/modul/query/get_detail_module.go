package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetDetailModule(app *fiber.Ctx) error {
	id := app.Params("id")
	res, err := repository.GetModulById(id)
	if err != nil {
		logrus.Error("Error while getting module: ", err.Error())
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}
	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    res,
		"message": "Success",
	})
}
