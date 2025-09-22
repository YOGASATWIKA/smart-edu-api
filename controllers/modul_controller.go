package controllers

import (
	command "smart-edu-api/service/modul/command"
	query "smart-edu-api/service/modul/query"

	"github.com/gofiber/fiber/v2"
)

func RouteModul(app *fiber.App) {
	modulGroup := app.Group("/modul")
	modulGroup.Post("/base-materi", command.CreateModul)
	modulGroup.Post("/outline", command.GenerateOutline)
	modulGroup.Get("/", query.GetModul)
	modulGroup.Get("/ebook/", query.GetEbook)
	modulGroup.Get("/activity/", query.GetActivity)
	modulGroup.Get("/:id", query.GetDetailModule)
}
