package controllers

import (
	auth "smart-edu-api/service/auth"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(app *fiber.App) {
	//Post Request to google
	app.Post("/api/auth/google", auth.HandleGoogleLogin)
	//Login & Register
	app.Post("/register", auth.HandleRegister)
	app.Post("/login", auth.HandleLogin)
	//get Data From Mongo
	profileGroup := app.Group("/api/profile", auth.AuthenticateToken)
	profileGroup.Get("/", auth.HandleGetProfile)
}
