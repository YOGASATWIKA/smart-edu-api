package controllers

import (
	command "smart-edu-api/service/model/command"
	query "smart-edu-api/service/model/query"

	"github.com/gofiber/fiber/v2"
)

func RouteModel(app *fiber.App) {
	modelGroup := app.Group("/model")
	modelGroup.Post("/outline", command.CreateModelOutline)
	modelGroup.Post("/ebook", command.CreateModelEbook)
	modelGroup.Put("/:id", command.UpdateModel)
	modelGroup.Get("/", query.GetModel)
	modelGroup.Delete("/:id", command.DeleteModel)
}
