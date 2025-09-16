package command

//
//import (
//	"context"
//	"os"
//	"smart-edu-api/llm"
//	"smart-edu-api/repository"
//
//	"github.com/gofiber/fiber/v2"
//	"github.com/joho/godotenv"
//	"github.com/tmc/langchaingo/llms"
//)
//
//type Process struct {
//	Ctx    context.Context
//	Model  llms.Model
//	Worker int
//}
//
//func CreateFullMateri(app *fiber.Ctx) error {
//	godotenv.Load()
//	ctx := context.Background()
//	id := app.Params("id")
//	APIKEY := os.Getenv("API_KEY")
//	model := llm.New(ctx, APIKEY)
//	ebook, err := repository.GetFullMateriById(id)
//	if err != nil {
//		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"message": "Materi tidak ditemukan",
//		})
//	}
//	process := Process{
//		Ctx:    ctx,
//		Model:  model,
//		Worker: 10,
//	}
//
//	lists := make([]entity.OutlineRoot, 0)
//
//	chOutline := make(chan *entity.OutlineRoot)
//}
