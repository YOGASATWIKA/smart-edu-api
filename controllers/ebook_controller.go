package controllers

import (
	"smart-edu-api/service/ebook/command"
	"smart-edu-api/service/ebook/query"

	"github.com/gofiber/fiber/v2"
)

func RouteEbook(app *fiber.App) {
	eBook := app.Group("/ebook")
	eBook.Post("/", command.CreateEbook)
	eBook.Get("/:id", query.GetEbookModuleById)
	eBook.Get("/detail:id", command.GetEbookById)
	eBook.Get("/pdf/:id", query.DownloadEbookById)
	eBook.Get("/word/:id", query.DownloadEbookWordById)
	eBook.Put("/:id", command.UpdateEbookById)

}
