package main

import (
	"os"
	"smart-edu-api/auth"
	"smart-edu-api/config"
	"smart-edu-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // 1. Tambahkan import ini
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load .env file, using system environment")
	}
}

func main() {
	InitEnv()
	config.InitMongoDB()

	app := fiber.New()

	// 2. Terapkan middleware CORS di sini
	godotenv.Load()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("PATHFE"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Daftarkan rute-rute Anda setelah middleware CORS
	auth.RegisterGoogleRoutes(app)
	controllers.RouteModel(app)
	controllers.RouteMateriPokok(app)
	controllers.RouteOutline(app)

	if err := app.Listen(":3001"); err != nil {
		logrus.Fatal("Error on running fiber: ", err.Error())
	}
}
