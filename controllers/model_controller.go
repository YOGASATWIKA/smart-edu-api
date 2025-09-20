package controllers

import (
	"github.com/gofiber/fiber/v2"
	command "smart-edu-api/service/model/command"
	query "smart-edu-api/service/model/query"
)

func RouteModel(app *fiber.App) {
	modelGroup := app.Group("/model")
	modelGroup.Post("/", command.CreateModel)
	modelGroup.Put("/:id", command.UpdateModel)
	modelGroup.Get("/", query.GetModel)
	modelGroup.Delete("/:id", command.DeleteModel)
}
