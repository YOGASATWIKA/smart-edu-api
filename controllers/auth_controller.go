package controllers

import (
	"log"
	auth "smart-edu-api/service/auth"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(app *fiber.App) {

	cloudHandler, err := auth.NewCloudinaryHandler()
	if err != nil {
		log.Fatal("Cloudinary init error:", err)
	}

	app.Post("/api/auth/google", auth.HandleGoogleLogin)
	app.Post("/register", auth.HandleRegister)
	app.Post("/login", auth.HandleLogin)
	app.Post("/forgot-password", auth.HandleForgotPassword)
	app.Post("/reset-password", auth.HandleResetPassword)
	profileGroup := app.Group("/api/profile", auth.AuthenticateToken)
	profileGroup.Get("/", auth.HandleGetProfile)
	profileGroup.Put("/", auth.HandleUpdateProfile)
	profileGroup.Post("/upload", cloudHandler.UploadImage)
}
