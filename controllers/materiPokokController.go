package controllers

import (
	"github.com/gofiber/fiber/v2"
	command "smart-edu-api/service/materiPokok/command"
	query "smart-edu-api/service/materiPokok/query"
)

func RouteMateriPokok(app *fiber.App) {
	materiPokok := app.Group("/materi-pokok")
	materiPokok.Post("/", command.CreateMateriPokok)
	materiPokok.Put("/:id", command.UpdateMateriPokok)
	materiPokok.Get("/", query.GetAllMateriPokok)
	materiPokok.Delete("/:id", command.DeleteMateriPokok)

}
