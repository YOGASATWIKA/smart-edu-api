package main

import (
	"smart-edu-api/config"
	"smart-edu-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())
	controllers.RouteBaseMateri(app)
	controllers.RouteOutline(app)

	if err := app.Listen(":3001"); err != nil {
		logrus.Fatal("Error on running fiber: ", err.Error())
	}
}