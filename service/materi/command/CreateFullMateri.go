package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/data/model/request"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"smart-edu-api/llm"
	"smart-edu-api/repository"
	"smart-edu-api/service/generator/materi/genai/background_urgency"
	"smart-edu-api/service/generator/materi/genai/base_competency"
	"smart-edu-api/service/generator/materi/genai/base_material"
	"smart-edu-api/service/generator/materi/genai/summary"
	"smart-edu-api/service/generator/materi/process/p1"
	"smart-edu-api/service/generator/materi/process/p2"
	"smart-edu-api/service/generator/materi/process/p3"
	"smart-edu-api/service/generator/materi/process/p4"
	"strings"
	"sync"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/tmc/langchaingo/llms"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Process struct {
	Ctx    context.Context
	Model  llms.Model
	Worker int
}

func CreateFullMateri(app *fiber.Ctx) error {
	var model llms.Model
	ctx := context.Background()

	request := new(request.ModelRequest)
	if err := app.BodyParser(request); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Invalid request body",
		})
	}
	isValid, err := govalidator.ValidateStruct(request)
	if !isValid && err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}
	APIKEY := os.Getenv("API_KEY")
	if request.Model == "Default" {
		model = llm.NewDefault(ctx, APIKEY)
	} else {
		model = llm.NewModel(APIKEY, request.Model)
	}

	process := Process{
		Ctx:    ctx,
		Model:  model,
		Worker: 5,
	}

	lists := make([]entity.Outline, 0)
	chOutline := make(chan *entity.Outline)

	//Load Data
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("outline")

	var objectIDs primitive.A
	for _, idStr := range request.Id {
		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			fmt.Printf("Error konversi ID '%s': %v\n", idStr, err)
			continue
		}
		objectIDs = append(objectIDs, objID)
	}

	fil := bson.D{
		{"_id", bson.D{
			{"$in", objectIDs},
		}},
	}

	cursor, err := collection.Find(ctx, fil)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		outlineRoot := entity.Outline{}
		err := cursor.Decode(&outlineRoot)
		if err != nil {
			log.Fatal(err)
		}
		lists = append(lists, outlineRoot)
	}

	//end Load data
	go func() {
		for _, outline := range lists {
			fmt.Println("Job Jabatan: ", outline.MateriPokok.Namajabatan, " Started")
			chOutline <- &outline
		}
	}()

	c1 := process.Process1WithWorker(chOutline, process.Worker)
	c2 := process.ProcessWithWorker(c1, process.Process2, process.Worker)
	c3 := process.ProcessWithWorker(c2, process.Process3, process.Worker)
	c4 := process.ProcessWithWorker(c3, process.Process4, process.Worker)

	for ebook := range c4 {
		err = repository.CreateMateri(ctx, entity.Ebook{
			ID:        primitive.NewObjectID(),
			Title:     ebook.Title,
			Parts:     ebook.Parts,
			Lock:      ebook.Lock,
			Type:      "-",
			CreatedAt: helper.GetCurrentTime(),
		})
		log.Printf(fmt.Sprintf("Done : %s", ebook.Title))
	}
	return app.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Request accepted. Outline generation is processing in the background.",
	})

}

func (p *Process) Process1WithWorker(chOutline <-chan *entity.Outline, worker int) <-chan *entity.Ebook {

	out := make(chan *entity.Ebook)
	wg := &sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(out)
	}()

	var listWk []<-chan *entity.Ebook
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		listWk = append(listWk, p.Process1(chOutline))
	}

	go func() {
		for _, ebook := range listWk {
			go func() {
				defer wg.Done()
				for e := range ebook {
					out <- e
				}
			}()
		}
	}()
	return out
}

func (p *Process) ProcessWithWorker(in <-chan *entity.Ebook, process func(in <-chan *entity.Ebook) <-chan *entity.Ebook, worker int) <-chan *entity.Ebook {

	out := make(chan *entity.Ebook)
	wg := sync.WaitGroup{}
	wg.Add(worker)
	go func() {
		wg.Wait()
		close(out)
	}()

	var listWk []<-chan *entity.Ebook
	for i := 0; i < worker; i++ {
		listWk = append(listWk, process(in))
	}

	go func() {
		for _, ebook := range listWk {
			go func() {
				defer wg.Done()
				for e := range ebook {
					out <- e
				}
			}()
		}
	}()
	return out

}

func (p *Process) Process1(in <-chan *entity.Outline) <-chan *entity.Ebook {
	out := make(chan *entity.Ebook)
	go func(ctx context.Context) {

		defer close(out)

		gen := base_material.NewBaseMaterial(p.Model)

		for {

			select {

			case <-ctx.Done():

				log.Println("receive terminate signal")

				return

			case outline, ok := <-in:

				if !ok {

					continue

				}

				ebook := &entity.Ebook{

					Title: outline.MateriPokok.Namajabatan,

					Lock: &sync.Mutex{},
				}

				fetch := p1.NewFetch(gen, outline)

				err := fetch.Fetch(p.Ctx, ebook)

				if err != nil {

					log.Println(err)

				}
				err = repository.CreateLog(ctx, entity.Ebook{

					ID: primitive.NewObjectID(),

					Title: ebook.Title,

					Parts: ebook.Parts,

					Lock: ebook.Lock,

					Type: fmt.Sprintf("%s.1", strings.ReplaceAll(outline.MateriPokok.Namajabatan, "/", "|")),

					CreatedAt: helper.GetCurrentTime(),
				})

				if err != nil {

					log.Fatal(err)

				}

				fmt.Println(fmt.Sprintf("Done : %s", ebook.Title))

				//f.Close()

				out <- ebook

			}

		}

	}(p.Ctx)

	return out

}

func (p *Process) Process2(in <-chan *entity.Ebook) <-chan *entity.Ebook {

	out := make(chan *entity.Ebook)

	go func(ctx context.Context) {

		defer close(out)

		bc := base_competency.NewBaseCompetency(p.Model)

		fetch := p2.NewFetch(bc)

		for {

			select {

			case <-ctx.Done():

				log.Println("receive terminate signal")

				return

			case ebook, ok := <-in:

				if !ok {

					continue

				}

				fetch.Fetch(p.Ctx, ebook)
				err := repository.CreateLog(ctx, entity.Ebook{
					ID:        primitive.NewObjectID(),
					Title:     ebook.Title,
					Parts:     ebook.Parts,
					Lock:      ebook.Lock,
					Type:      fmt.Sprintf("%s.2", strings.ReplaceAll(ebook.Title, "/", "|")),
					CreatedAt: helper.GetCurrentTime(),
				})

				if err != nil {

					log.Fatal(err)

				}

				out <- ebook

			}

		}

	}(p.Ctx)

	return out

}

func (p *Process) Process3(in <-chan *entity.Ebook) <-chan *entity.Ebook {

	out := make(chan *entity.Ebook)

	sum := summary.NewSummary(p.Model)

	fetch := p3.NewFetch(sum)

	go func(ctx context.Context) {

		defer close(out)

		for {

			select {

			case <-ctx.Done():

				log.Println("receive terminate signal")

				return

			case ebook, ok := <-in:

				if !ok {

					continue

				}

				fetch.Fetch(ctx, ebook)
				err := repository.CreateLog(ctx, entity.Ebook{

					ID: primitive.NewObjectID(),

					Title: ebook.Title,

					Parts: ebook.Parts,

					Lock: ebook.Lock,

					Type: fmt.Sprintf("%s.3", strings.ReplaceAll(ebook.Title, "/", "|")),

					CreatedAt: helper.GetCurrentTime(),
				})

				if err != nil {

					log.Fatal(err)

				}

				out <- ebook

			}

		}

	}(p.Ctx)

	return out

}

func (p *Process) Process4(in <-chan *entity.Ebook) <-chan *entity.Ebook {
	out := make(chan *entity.Ebook)
	go func(ctx context.Context) {
		defer close(out)
		bu := background_urgency.NewBackgroundUrgency(p.Model)
		fetch := p4.NewFetch(bu)
		for {
			select {
			case <-ctx.Done():
				log.Println("receive terminate signal")
				return
			case ebook, ok := <-in:

				if !ok {
					continue
				}
				fetch.Fetch(ctx, ebook)

				err := repository.CreateLog(ctx, entity.Ebook{
					ID:        primitive.NewObjectID(),
					Title:     ebook.Title,
					Parts:     ebook.Parts,
					Lock:      ebook.Lock,
					Type:      fmt.Sprintf("%s.4", strings.ReplaceAll(ebook.Title, "/", "|")),
					CreatedAt: helper.GetCurrentTime(),
				})
				if err != nil {

					log.Fatal(err)

				}
				out <- ebook
			}
		}
	}(p.Ctx)
	return out
}
