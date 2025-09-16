package command

import (
	"context"
	"log"
	"os"
	"smart-edu-api/generator/materi/genai/expand_material"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ThirdGenExpand(app *fiber.Ctx) error {
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

	expand := expand_material.NewExpandMaterial(model)

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {

			var learnings []string

			for _, material := range chapter.Materials {
				learnings = append(learnings, material.Title)
			}

			for _, material := range chapter.Materials {
				for _, detail := range material.Details {

					try := 1

					for {
						log.Println("Expanding words \"", detail.Content, "\" with try:", try)

						out, err := expand.Generate(ctx, expand_material.ExpandedMaterialPrompt{
							Purpose:        "Materi bahan ajar seleksi kompetensi bidang (SKB) calon aparatur sipil negara (seleksi CPNS) untuk formasi jabatan " + ebook.Title,
							BaseMaterial:   part.Subject,
							SubMaterial:    chapter.Title,
							Learnings:      learnings,
							MainTopics:     chapter.BaseCompetitions,
							ChosenLearning: material.Title,
							WordsToExpand:  detail.Content,
						})

						if err != nil {
							try++
							continue
						}

						detail.Expanded = out.Expanded

						break
					}
				}
			}
		}
	}

	err = repository.CreateLog(ctx, ebook)
	log.Println("log materi success insert to database")
	ebook.Type = "3"
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
