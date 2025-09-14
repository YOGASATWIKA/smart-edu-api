package controllers

import (
	command "smart-edu-api/service/materiPokok/command"
	query "smart-edu-api/service/materiPokok/query"

	"github.com/gofiber/fiber/v2"
)

func RouteMateriPokok(app *fiber.App) {
	materiPokok := app.Group("/materi-pokok")
	materiPokok.Post("/", command.CreateMateriPokok)
	materiPokok.Get("/", query.GetAllMateriPokok)
	materiPokok.Delete("/:id", command.DeleteMateri)

}
