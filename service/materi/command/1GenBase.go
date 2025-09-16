package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"smart-edu-api/embeded"
	"smart-edu-api/entity"
	"smart-edu-api/generator/materi/genai/base_material"
	"smart-edu-api/helper"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fetch struct {
	generator *base_material.BaseMaterial
	ctx       context.Context
	ebook     *entity.Ebook
	saveFile  *os.File
	outline   *entity.Outline
}

func FirstGenBase(app *fiber.Ctx) error {
	godotenv.Load()
	ctx := context.Background()
	id := app.Params("id")
	APIKEY := os.Getenv("API_KEY")
	//menggunakan open router
	//var req request.ModelRequest
	//model := llm.New(ctx, APIKEY, req.Model)
	//menggunakan gemini
	model := llm.New(ctx, APIKEY)

	outline, err := repository.GetOutlineByMateriPokokId(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Outline tidak ditemukan",
		})
	}

	var ebook = &entity.Ebook{
		Title: outline.MateriPokok.Namajabatan,
		Lock:  &sync.Mutex{},
	}

	bm := base_material.NewBaseMaterial(model)

	f := &fetch{
		generator: bm,
		ctx:       ctx,
		ebook:     ebook,
		outline:   outline,
	}

	err = f.fetchMaterial()
	if err != nil {
		log.Fatal(err)
	}

	err = repository.CreateMateri(ctx, entity.Ebook{
		ID:        primitive.NewObjectID(),
		Title:     ebook.Title,
		Parts:     ebook.Parts,
		Lock:      ebook.Lock,
		Type:      "1",
		CreatedAt: helper.GetCurrentTime(),
	})
	log.Printf("materi %s success inserted to database", ebook.Title)
	if err != nil {
		log.Printf("ERROR: Failed to save materi for %s: %v", ebook.Title, err)
	}
	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Request accepted",
		"jabatan": ebook,
	})
}

func (f *fetch) fetchMaterial() error {
	for _, listMateri := range f.outline.Outline.ListMateri {
		log.Println("[1_fetchMaterial] Fetching material :", listMateri)
		var part = &entity.Part{
			Subject: listMateri.MateriPokok,
		}

		err := f.fetchSubMateri(&listMateri, listMateri.MateriPokok, part)
		if err != nil {
			return err
		}

		f.ebook.Parts = append(f.ebook.Parts, part)
	}

	return nil
}

func (f *fetch) fetchSubMateri(m *embeded.MateriPokok, mp string, part *entity.Part) error {
	for _, listSubMateri := range m.ListSubMateri {
		log.Println("[2_fetchSubMaterial] Fetching sub material :", listSubMateri)
		var chapter = &entity.Chapter{
			Title: listSubMateri.SubMateriPokok,
		}

		err := f.fetchListMateri(&listSubMateri, mp, listSubMateri.SubMateriPokok, chapter)
		if err != nil {
			return err
		}

		part.Chapters = append(part.Chapters, chapter)
	}

	return nil
}

func (f *fetch) fetchListMateri(s *embeded.SubMateriPokok, mp, smp string, chapter *entity.Chapter) error {
	for _, materi := range s.ListMateri {
		log.Println("[3_fetchListMaterial] Fetching material :", materi)
		var g = &entity.Material{
			Title: materi,
		}

		try := 1

		for {
			log.Println("[3_fetchListMaterial] Fetching material try : ", try)
			out, err := f.generator.Generate(f.ctx, base_material.MaterialPrompt{
				Purpose:  fmt.Sprintf("Seleksi Kompetensi Bidang untuk calon pegawai negeri untuk jabatan %s", f.outline.MateriPokok.Namajabatan),
				Subject:  mp,
				Chapter:  smp,
				Material: materi,
			})

			if err != nil {
				log.Println("", err)
				try++
				continue
			}

			g.Short = out.Short

			for _, detail := range out.Details {
				g.Details = append(g.Details, &entity.Detail{
					Content: detail,
				})
			}

			log.Println("[3_fetchListMaterial] result : ", g.Details)

			break
		}

		chapter.Materials = append(chapter.Materials, g)
	}

	return nil
}
