package command

import (
	"context"
	"log"
	"os"
	"smart-edu-api/generator/materi/genai/background_urgency"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var PREFIX = "pkb-teknis"

func FiveBackground(app *fiber.Ctx) error {
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
	bg := background_urgency.NewBackgroundUrgency(model)
	ebook.Lock = &sync.Mutex{}

	for _, part := range ebook.Parts {
		var competencies = make(map[string][]string)

		for _, chapter := range part.Chapters {
			competencies[chapter.Title] = chapter.BaseCompetitions
		}

		try := 1

		for {

			log.Println("Generate background & urgencies for: ", part.Subject, " with try: ", try)

			out, err := bg.Generate(ctx, background_urgency.Prompt{
				Purpose:      "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan " + ebook.Title,
				BaseMaterial: part.Subject,
				Competencies: competencies,
			})

			if err != nil {
				try++
				continue
			}

			part.Introductions = out.Backgrounds
			part.Urgencies = out.Urgencies

			break
		}
	}
	err = repository.CreateLog(ctx, ebook)
	log.Println("log materi success insert to database")
	ebook.Type = "5"
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
