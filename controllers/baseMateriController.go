package controllers
import (
	command "smart-edu-api/service/baseMateri/command"
	query "smart-edu-api/service/baseMateri/query"
	"github.com/gofiber/fiber/v2"
)

func RouteBaseMateri(app *fiber.App) {
	baseMateriGroup := app.Group("/base-materi")
	baseMateriGroup.Post("/", command.CreateBaseMateri)
	baseMateriGroup.Put("/:id", command.UpdateBaseMateri)
	baseMateriGroup.Get("/", query.GetBaseMateri)
	baseMateriGroup.Delete("/:id", command.DeleteBaseMateri)

}
