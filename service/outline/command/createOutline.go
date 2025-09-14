package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/embeded"
	"smart-edu-api/entity"
	generator "smart-edu-api/generator/outline"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Job struct {
	Jabatan *entity.Materi
	Outline *embeded.Outline
	Err     error
}
type ModelRequest struct {
	Model string `json:"model"`
}

func CreateOutline(app *fiber.Ctx) error {
	godotenv.Load()
	ctx := context.Background()

	// ambil ID dari path parameter
	id := app.Params("id")

	var req ModelRequest
	if err := app.BodyParser(&req); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	// ambil jabatan dari DB berdasarkan _id
	jabatan, err := repository.GetMateriPokokByID(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	// proses generate outline untuk satu data
	APIKEY := os.Getenv("API_KEY")
	model := llm.New(ctx, APIKEY, req.Model)

	o := generator.New(model)

	otln, err := o.Generate(ctx, generator.Params{
		NamaJabatan:  jabatan.Namajabatan,
		TugasJabatan: strings.Join(jabatan.Tugasjabatan, ", "),
		Keterampilan: strings.Join(jabatan.Keterampilan, ", "),
	})
	log.Println("generating outline done")
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate outline",
			"error":   err.Error(),
		})
	}

	// simpan ke file
	filename := fmt.Sprintf("./output/%s.json", strings.ReplaceAll(jabatan.Namajabatan, "/", "|"))
	f, err := os.Create(filename)
	if err == nil {
		_ = json.NewEncoder(f).Encode(otln)
		_ = f.Close()
	}

	// update DB
	err = CreateOrUpdateOutline(ctx, &Job{
		Jabatan: jabatan,
		Outline: &otln,
	}, "OUTLINE")
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to save outline",
			"error":   err.Error(),
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Outline generated successfully",
		"jabatan": jabatan.Namajabatan,
	})
}

func CreateOrUpdateOutline(ctx context.Context, job *Job, stage string) error {

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.D{
		{
			"_id", job.Jabatan.ID,
		},
	}

	update := bson.D{
		{
			"$set", bson.D{
				{"stage", stage},
				{"outline", job.Outline},
			},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
