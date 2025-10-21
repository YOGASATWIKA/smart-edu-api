package controllers

import (
	auth "smart-edu-api/service/auth"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(app *fiber.App) {
	app.Post("/api/auth/google", auth.HandleGoogleLogin)
	app.Post("/register", auth.HandleRegister)
	app.Post("/login", auth.HandleLogin)
	profileGroup := app.Group("/api/profile", auth.AuthenticateToken)
	profileGroup.Get("/", auth.HandleGetProfile)
}
