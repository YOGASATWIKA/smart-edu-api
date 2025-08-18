package controllers

import (
	command "smart-edu-api/service/promt/command"
	query "smart-edu-api/service/promt/query"
	"github.com/gofiber/fiber/v2"
)

func RoutePromt(app *fiber.App) {
	promtGroup := app.Group("/promt")
	promtGroup.Post("/", command.CreatePromt)
	promtGroup.Put("/:id", command.UpdatePromt)
	promtGroup.Get("/", query.Getpromt)
	promtGroup.Delete("/:id", command.DeletePromt)
}
