package command

import (
	"context"
	"log"
	"smart-edu-api/repository"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SixChunckExpand(app *fiber.Ctx) error {
	godotenv.Load()
	ctx := context.Background()
	id := app.Params("id")
	ebook, err := repository.GetFullMateriById(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Materi tidak ditemukan",
		})
	}

	ebook.Lock = &sync.Mutex{}

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {
			for _, material := range chapter.Materials {
				for _, detail := range material.Details {
					chunks := strings.Split(detail.Expanded, "\n\n")

					for _, chunk := range chunks {
						f := strings.TrimSpace(chunk)
						if f == "" {
							continue
						}

						detail.ExpandChunks = append(detail.ExpandChunks, f)
					}
				}
			}
		}
	}

	err = repository.CreateLog(ctx, ebook)
	log.Println("log materi success insert to database")
	ebook.Type = "6"
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
