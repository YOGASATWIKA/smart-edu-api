package service

//
//import (
//	"log"
//	"smart-edu-api/data/model/request"
//	"smart-edu-api/entity"
//	"smart-edu-api/repository"
//
//	"github.com/gofiber/fiber/v2"
//)
//
//func GenerateOutline(app *fiber.Ctx) error  {
//	id := app.Params("id")
//	//var req request.ModelRequest
//	//if err := app.BodyParser(&req); err != nil {
//	//	return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//	//		"message": "Cannot parse JSON",
//	//		"error":   err.Error(),
//	//	})
//	//}
//	materiPokok, err := repository.GetMateriPokokByID(id)
//	if err != nil {
//		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"message": "Base Materi tidak ditemukan",
//		})
//	}
//
//	ch := make(chan *entity.Outline)
//
//	go func() {
//
//		defer close(ch)
//		for _, mp := range materiPokok {
//			ch <- &Job{
//				Jabatan: &j,
//			}
//			log.Println("Spawn process for", j.NamaJabatan)
//		}
//	}()
//
//
//}
