package controllers

import (
	commandOutline "smart-edu-api/service/outline/command"
	queryBase "smart-edu-api/service/outline/query"

	"github.com/gofiber/fiber/v2"
)

func RouteOutline(app *fiber.App) {
	outlineGroup := app.Group("/outline")
	outlineGroup.Post("/:id", commandOutline.CreateOutline)
	outlineGroup.Get("/:id", queryBase.GetOutlineById)
	outlineGroup.Put("/:id", commandOutline.UpdateOutline)
	// baseMateriGroup.Delete("/:id", commandBase.DeleteBaseMateri)

}
