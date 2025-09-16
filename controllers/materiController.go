package controllers

import (
	"smart-edu-api/service/materi/command"
	"smart-edu-api/service/materi/query"

	"github.com/gofiber/fiber/v2"
)

func RouteMateri(app *fiber.App) {
	materiPokok := app.Group("/materi")
	materiPokok.Post("/first/:id", command.FirstGenBase)
	materiPokok.Post("/two/:id", command.TwoGenBase)
	materiPokok.Post("/third/:id", command.ThirdGenExpand)
	materiPokok.Post("/four/:id", command.FourSummary)
	materiPokok.Post("/five/:id", command.FiveBackground)
	materiPokok.Post("/six/:id", command.SixChunckExpand)
	materiPokok.Get("/:id", query.GetMateriById)

}
