package controllers
import (
	commandBase "smart-edu-api/service/baseMateri/command"
	queryBase "smart-edu-api/service/baseMateri/query"
	"github.com/gofiber/fiber/v2"
)

func RouteBaseMateri(app *fiber.App) {
	baseMateriGroup := app.Group("/base-materi")
	baseMateriGroup.Post("/", commandBase.CreateBaseMateri)
	baseMateriGroup.Put("/:id", commandBase.UpdateBaseMateri)
	baseMateriGroup.Get("/", queryBase.GetBaseMateri)
	baseMateriGroup.Delete("/:id", commandBase.DeleteBaseMateri)

}
