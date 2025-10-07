package controllers

import (
	command "smart-edu-api/service/model/command"
	query "smart-edu-api/service/model/query"

	"github.com/gofiber/fiber/v2"
)

func RouteModel(app *fiber.App) {
	modelGroup := app.Group("/model")
	modelGroup.Post("/", command.CreateModel)
	modelGroup.Put("/outline/:id", command.UpdateModel)
	modelGroup.Get("/outline", query.GetOutlineModel)
	modelGroup.Delete("/:id", command.DeleteModel)
}
