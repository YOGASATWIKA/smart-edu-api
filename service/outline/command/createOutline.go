package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"smart-edu-api/data/model/request"
	"smart-edu-api/embeded"
	"smart-edu-api/entity"
	generator "smart-edu-api/generator/outline"
	"smart-edu-api/helper"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOutline(app *fiber.Ctx) error {
	id := app.Params("id")

	var req request.ModelRequest
	if err := app.BodyParser(&req); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	materiPokok, err := repository.GetMateriPokokByID(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}
	go func() {
		ctx := context.Background()
		log.Printf("Starting background process for Jabatan: %s", materiPokok.Namajabatan)

		APIKEY := os.Getenv("API_KEY")
		//model := llm.New(ctx, APIKEY, req.Model)
		model := llm.New(ctx, APIKEY)
		o := generator.New(model)

		otln, err := o.Generate(ctx, generator.Params{
			NamaJabatan:  materiPokok.Namajabatan,
			TugasJabatan: strings.Join(materiPokok.Tugasjabatan, ", "),
			Keterampilan: strings.Join(materiPokok.Keterampilan, ", "),
		})
		if err != nil {
			log.Printf("ERROR: Failed to generate outline for %s: %v", materiPokok.Namajabatan, err)
			return // Hentikan goroutine jika gagal
		}
		log.Printf("Generating outline done for: %s", materiPokok.Namajabatan)

		filename := fmt.Sprintf("./output/%s.json", strings.ReplaceAll(materiPokok.Namajabatan, "/", "|"))
		f, _ := os.Create(filename)
		if f != nil {
			_ = json.NewEncoder(f).Encode(otln)
			_ = f.Close()
		}

		res := embeded.MateriPokokEmbed{
			ID:          materiPokok.ID,
			Namajabatan: materiPokok.Namajabatan,
		}

		err = repository.CreateOutline(ctx, entity.Outline{
			ID:          primitive.NewObjectID(),
			MateriPokok: res,
			Outline:     otln,
			Model:       req.Model,
			CreatedAt:   helper.GetCurrentTime(),
		})
		if err != nil {
			log.Printf("ERROR: Failed to save outline for %s: %v", materiPokok.Namajabatan, err)
			return
		}

		log.Printf("Successfully processed and saved outline for: %s", materiPokok.Namajabatan)
	}()
	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Request accepted. Outline generation is processing in the background.",
		"jabatan": materiPokok.Namajabatan,
	})
}
