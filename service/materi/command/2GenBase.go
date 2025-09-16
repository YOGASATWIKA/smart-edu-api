package command

import (
	"context"
	"log"
	"os"
	"smart-edu-api/generator/materi/genai/base_competency"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TwoGenBase(app *fiber.Ctx) error {
	godotenv.Load()
	ctx := context.Background()
	id := app.Params("id")
	APIKEY := os.Getenv("API_KEY")
	model := llm.New(ctx, APIKEY)

	ebook, err := repository.GetFullMateriById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Materi tidak ditemukan",
		})
	}

	ebook.Lock = &sync.Mutex{}

	if err != nil {
		log.Fatal(err)
	}

	bc := base_competency.NewBaseCompetency(model)

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {
			var materials []string

			for _, material := range chapter.Materials {
				materials = append(materials, material.Title)
			}

			try := 1

			for {
				log.Println("Generate objective & trigger question for ", chapter.Title, "with try : ", try)

				out, err := bc.Generate(ctx, base_competency.CompetencyPrompt{
					Purpose:      "Seleksi Kompetensi Bidang calon pegwai negeri sipili pada jabatan " + ebook.Title,
					BaseMaterial: part.Subject,
					SubMaterial:  chapter.Title,
					Learnings:    materials,
				})

				if err != nil {
					try++
					continue
				}

				chapter.BaseCompetitions = out.Objectives
				chapter.TriggerQuestions = out.TriggerQuestions

				break
			}
		}
	}

	err = repository.CreateLog(ctx, ebook)
	log.Println("log materi success insert to database")
	ebook.Type = "2"
	updated, err := repository.UpdateMateri(ebook)
	log.Printf("materi %s success updated to database", ebook.Title)
	if err != nil {
		log.Printf("ERROR: Failed to save materi for %s: %v", ebook.Title, err)
	}
	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Request accepted",
		"jabatan": updated.Title,
	})
}
