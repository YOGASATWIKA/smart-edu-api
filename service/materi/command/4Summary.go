package command

import (
	"context"
	"log"
	"os"
	"smart-edu-api/generator/materi/genai/summary"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func FourSummary(app *fiber.Ctx) error {
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

	sum := summary.NewSummary(model)

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {

			var materials = make(map[string]string)

			for _, material := range chapter.Materials {
				if _, ok := materials[material.Title]; !ok {
					materials[material.Title] = material.Short
				}
			}

			try := 1

			for {

				log.Println("Generate summary & reflection for: ", chapter.Title, " with try: ", try)

				out, err := sum.Generate(ctx, summary.Prompt{
					Purpose:      "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan " + ebook.Title,
					BaseMaterial: part.Subject,
					SubMaterial:  chapter.Title,
					Materials:    materials,
				})

				if err != nil {
					try++
					continue
				}

				chapter.Reflections = out.Reflections
				chapter.Conclusion = out.Summary

				break
			}
		}
	}
	err = repository.CreateLog(ctx, ebook)
	log.Println("log materi success insert to database")
	ebook.Type = "4"
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
