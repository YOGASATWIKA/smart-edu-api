package query

import (
	"smart-edu-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetDetailModel(app *fiber.Ctx) error {
	model := app.Params("model")
	res, err := repository.GetModelByModel(model)
	if err != nil {
		logrus.Error("Error while getting base materi: ", err.Error())
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}
	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    res,
		"message": "Success",
	})
}
