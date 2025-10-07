package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/data/modul/request"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"smart-edu-api/service/generator/outline"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/petermattis/goid"
	"github.com/tmc/langchaingo/llms"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateOutline(app *fiber.Ctx) error {
	request := new(modul.ModelRequest)
	ctx := context.Background()
	if err := app.BodyParser(&request); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	mongo, err := config.New(config.GetMongoDBConnectionString())
	if err != nil {
		log.Fatal("error connecting to db")
	}

	materiPokok := loadData(request.Id, mongo)
	ch := make(chan *entity.Modul)

	go func() {
		defer close(ch)
		for _, j := range materiPokok {
			ch <- &j
			log.Println("Spawn process for", j.MateriPokok.Namajabatan)
		}
	}()

	ch1 := RunGenerateOutlineWithWorker(request.Model, ctx, ch, 3)

	for job := range ch1 {

		modul := entity.Modul{
			ID:          job.ID,
			MateriPokok: job.MateriPokok,
			Outline:     job.Outline,
			Status:      job.Status,
			State:       "OUTLINE",
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   helper.GetCurrentTime(),
		}
		updated, err := repository.UpdateModul(ctx, modul)
		if err != nil {
			log.Printf("ERROR: Failed to save outline for %s: %v", updated.MateriPokok.Namajabatan, err)
			return nil
		}
		log.Printf("Successfully processed and saved outline for: %s", updated.MateriPokok.Namajabatan)
	}

	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Request accepted. Outline generation is processing in the background.",
	})
}

func RunGenerateOutlineWithWorker(request string, ctx context.Context, ch <-chan *entity.Modul, worker int) <-chan *entity.Modul {
	out := make(chan *entity.Modul)
	var workers []<-chan *entity.Modul
	wg := &sync.WaitGroup{}
	wg.Add(worker)

	go func() {
		wg.Wait()
		close(out)
	}()

	for i := 0; i < worker; i++ {
		workers = append(workers, RunGenerateOutline(request, ctx, ch))
	}

	go func() {
		for _, j := range workers {
			go func() {
				defer wg.Done()

				for job := range j {
					out <- job
				}
			}()
		}
	}()

	return out
}

func RunGenerateOutline(request string, ctx context.Context, in <-chan *entity.Modul) <-chan *entity.Modul {

	var out = make(chan *entity.Modul)

	err := os.MkdirAll("./output", 0775)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(out)
		var model llms.Model
		DEFAULTMODEL := os.Getenv("GEMINI_API_KEY")
		CUSTOMEMODEL := os.Getenv("OPENAI_API_KEY")
		if request == "Default" {
			model = llm.NewDefault(ctx, DEFAULTMODEL)
		} else {
			model = llm.NewModel(CUSTOMEMODEL, request)
		}
		o := generator.New(model)

		for {
			select {
			case <-ctx.Done():
				log.Println("receiving close signal, terminating goroutine")
				return
			case job, ok := <-in:
				if !ok {
					log.Println("channel closing, terminating goroutine", goid.Get())
					return
				}

				log.Println("Processing job", job.MateriPokok.Namajabatan)

				otln, err := o.Generate(request, ctx, generator.Params{
					NamaJabatan:  job.MateriPokok.Namajabatan,
					TugasJabatan: strings.Join(job.MateriPokok.Tugasjabatan, ", "),
					Keterampilan: strings.Join(job.MateriPokok.Keterampilan, ", "),
				})

				if err != nil {
					out <- job
					continue
				}

				job.Outline = otln
				// job.Outline.NamaJabatan = fmt.Sprintf("%s", job.Jabatan.NamaJabatan)

				filename := fmt.Sprintf("./output/%s.json", strings.ReplaceAll(job.MateriPokok.Namajabatan, "/", "|"))

				f, err := os.Create(filename)
				if err != nil {
					out <- job
					continue
				}

				err = json.NewEncoder(f).Encode(job.Outline)
				if err != nil {
					out <- job
					continue
				}
				_ = f.Close()

				log.Println("successfully saved outline:", filename)

				out <- job
			}
		}
	}()

	return out
}

func loadData(id []string, db *config.Database) []entity.Modul {
	var objectIDs primitive.A
	var jabatans = make([]entity.Modul, 0)
	for _, idStr := range id {
		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			fmt.Printf("Error konversi ID '%s': %v\n", idStr, err)
			continue
		}
		objectIDs = append(objectIDs, objID)
	}
	filter := bson.D{
		{"_id", bson.D{
			{"$in", objectIDs},
		}},
	}
	ctx := context.Background()
	collection := db.Conn.Database("smart_edu").Collection("modul")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []entity.Modul{}
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result entity.Modul
		err := cursor.Decode(&result)

		if err != nil {
			fmt.Println("error collecting data")
			return []entity.Modul{}
		}

		jabatans = append(jabatans, result)
	}

	return jabatans
}
