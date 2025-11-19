package main

import (
	"os"
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
	//Init .env file
	InitEnv()
	config.InitMongoDB()
	app := fiber.New()
	//cors url to front end
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:  os.Getenv("PATHFE"),
	//	AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
	//	AllowHeaders:  "*",
	//	ExposeHeaders: "Content-Disposition, X-Filename",
	//}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://smart-edu-production.up.railway.app, http://localhost:5173, http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	controllers.RouteAuth(app)
	controllers.RouteModel(app)
	controllers.RouteModul(app)
	controllers.RouteEbook(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	if err := app.Listen("0.0.0.0:" + port); err != nil {
		logrus.Fatal("Error on running fiber: ", err.Error())
	}
}
