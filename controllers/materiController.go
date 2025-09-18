package controllers

import (
	"smart-edu-api/service/materi/command"
	"smart-edu-api/service/materi/query"

	"github.com/gofiber/fiber/v2"
)

func RouteMateri(app *fiber.App) {
	materiPokok := app.Group("/materi")
	materiPokok.Post("/", command.CreateFullMateri)
	materiPokok.Get("/:id", query.GetMateriById)

}
