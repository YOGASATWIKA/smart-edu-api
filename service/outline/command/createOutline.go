package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"smart-edu-api/config"
	outline "smart-edu-api/data/outline/request"
	"smart-edu-api/generator"
	"smart-edu-api/llm"
	"smart-edu-api/model"
	"smart-edu-api/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Job struct {
	Jabatan *model.MateriPokok
	Outline *outline.Outline
	Err     error
}
func CreateOutline(app *fiber.Ctx) error {
	godotenv.Load()
	ctx := context.Background()
	
	// ambil ID dari path parameter
	id := app.Params("id")

	// ambil jabatan dari DB berdasarkan _id
	jabatan, err := utils.GetBaseMateriByID(id)
	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	// proses generate outline untuk satu data
	APIKEY := os.Getenv("API_KEY")
	model := llm.New(ctx, APIKEY)
	o := generator.New(model)

	otln, err := o.GenerateWithOfficialMaterial(ctx, generator.Params{
		NamaJabatan:     jabatan.Namajabatan,
		TugasJabatan:    strings.Join(jabatan.Tugasjabatan, ", "),
		Keterampilan:    strings.Join(jabatan.Keterampilan, ", "),
	})
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
	}, "success")
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

func CreateOrUpdateOutline(ctx context.Context, job *Job, state string) error {
		
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
				{"state", state},
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