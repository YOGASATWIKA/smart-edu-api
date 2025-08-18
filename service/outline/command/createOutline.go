package command

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"smart-edu-api/llm"

// 	"github.com/gofiber/fiber/v2"
// )

// func CreateOutline(c *fiber.Ctx) error{
// 	ctx := context.Background()
// 	apiKey := os.Getenv("API_KEY")
// 	llmModel := os.Getenv("MODEL")

// 	client := llm.New(ctx, apiKey, llmModel)

// 	answer, err := client.Chat(ctx, "Apa itu kontol")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(
// 		map[string]any{
// 			"data":   answer,
// 			"message": "Test",
// 		},
// 	)
// }

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"smart-edu-api/config"
// 	"smart-edu-api/llm"
// 	"smart-edu-api/outline"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// 	"github.com/petermattis/goid"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Jabatan struct {
// 	JobID              uint   `json:"job_id" bson:"job_id"`
// 	NamaJabatan        string `json:"nama_jabatan" bson:"nama_jabatan"`
// 	TugasJabatan       string `json:"tugas_jabatan" bson:"tugas_jabatan"`
// 	Keterampilan       string `json:"keterampilan" bson:"keterampilan"`
// 	State              string `json:"state" bson:"state"`
// 	JumlahSoal         uint   `json:"jumlah_soal" bson:"jumlah_soal"`
// 	Remain             uint   `json:"remain" bson:"remain"`
// 	MaterialState      string `json:"material_state" bson:"material_state"`
// 	MaterialMergeState string `json:"material_merge_state" bson:"material_merge_state"`
// 	AvailableMateri    string `json:"available_materi" bson:"available_materi"`
// }

// type Job struct {
// 	Jabatan *Jabatan
// 	Outline *outline.Outline
// 	Err     error
// }

// func CreateOutline(c *fiber.Ctx) error{
// 	godotenv.Load()
// 	ctx := context.Background()

// 	mongo, err := config.New(config.GetMongoDBConnectionString())
// 	if err != nil {
// 		log.Fatal("error connecting to db")
// 	}

// 	jabatans := loadData(mongo)

// 	ch := make(chan *Job)

// 	go func() {
// 		defer close(ch)

// 		for _, j := range jabatans {
// 			ch <- &Job{
// 				Jabatan: &j,
// 			}
// 			log.Println("Spawn process for", j.NamaJabatan)
// 		}
// 	}()

// 	ch1 := RunGenerateOutlineWithWorker(ctx, ch, 10)

// 	for job := range ch1 {
// 		if job.Err != nil {
// 			fmt.Println(job.Jabatan.NamaJabatan, "is failed with error:", job.Err)
// 		} else {
// 			err := CreateOrUpdateOutline(ctx, mongo, job, "success")
// 			if err != nil {
// 				fmt.Println("failed to write db")
// 			}
// 			fmt.Println(job.Jabatan.NamaJabatan, "is processed")
// 		}
// 	}
	
// 	return c.Status(fiber.StatusOK).JSON(
// 		map[string]any{
// 			"data":   jabatans,
// 			"message": "Base materi created successfully",
// 		},
// 	)
// }


// func RunGenerateOutlineWithWorker(ctx context.Context, ch <-chan *Job, worker int) <-chan *Job {
// 	out := make(chan *Job)
// 	var workers []<-chan *Job
// 	wg := &sync.WaitGroup{}
// 	wg.Add(worker)

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	for i := 0; i < worker; i++ {
// 		workers = append(workers, RunGenerateOutline(ctx, ch))
// 	}

// 	go func() {
// 		for _, j := range workers {
// 			go func() {
// 				defer wg.Done()

// 				for job := range j {
// 					out <- job
// 				}
// 			}()
// 		}
// 	}()

// 	return out
// }

// func RunGenerateOutline(ctx context.Context, in <-chan *Job) <-chan *Job {

// 	var out = make(chan *Job)

// 	err := os.MkdirAll("./output", 0775)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	go func() {
// 		defer close(out)
// 		// Load .env file
// 		err := godotenv.Load(".env")
// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 		}

// 		APIKEY := os.Getenv("GOOGLE_API_KEY")

// 		model := llm.New(ctx, APIKEY)

// 		o := outline.New(model)

// 		for {
// 			select {
// 			case <-ctx.Done():
// 				log.Println("receiving close signal, terminating goroutine")
// 				return
// 			case job, ok := <-in:
// 				if !ok {
// 					log.Println("channel closing, terminating goroutine", goid.Get())
// 					return
// 				}

// 				log.Println("Processing job", job.Jabatan.NamaJabatan)

// 				otln, err := o.GenerateWithOfficialMaterial(ctx, outline.Params{
// 					NamaJabatan:     job.Jabatan.NamaJabatan,
// 					TugasJabatan:    job.Jabatan.TugasJabatan,
// 					Keterampilan:    job.Jabatan.Keterampilan,
// 					AvailableMateri: job.Jabatan.AvailableMateri,
// 				})

// 				if err != nil {
// 					job.Err = err
// 					out <- job
// 					continue
// 				}

// 				job.Outline = &otln
// 				// job.Outline.NamaJabatan = fmt.Sprintf("%s", job.Jabatan.NamaJabatan)

// 				filename := fmt.Sprintf("./output/%s.json", strings.ReplaceAll(job.Jabatan.NamaJabatan, "/", "|"))

// 				f, err := os.Create(filename)
// 				if err != nil {
// 					job.Err = err
// 					out <- job
// 					continue
// 				}

// 				err = json.NewEncoder(f).Encode(job.Outline)
// 				if err != nil {
// 					job.Err = err
// 					out <- job
// 					continue
// 				}
// 				_ = f.Close()

// 				log.Println("successfully saved outline:", filename)

// 				out <- job
// 			}
// 		}
// 	}()

// 	return out
// }

// func loadData(db *config.Database) []Jabatan {

// 	var jabatans = make([]Jabatan, 0)

// 	// For bulk
// 	// filter := bson.D{
// 	// 	{"state", bson.D{
// 	// 		{"$in", bson.A{"not_started", "failed"}},
// 	// 	}},
// 	// }

// 	//For single
// 	filter := bson.M{"job_id": 2}

// 	// for more than 1
// 	// sss := []uint{1,2,3}
// 	// filter := bson.D{
// 	// 	{"job_id", bson.D{
// 	// 		{"$in", sss},
// 	// 	}},
// 	// }

// 	//for filter by files
// 	//sss := getListJabatanJSON()
// 	//filter := bson.D{
// 	//	{"nama_jabatan", bson.D{
// 	//		{"$in", sss},
// 	//	}},
// 	//	{"state", bson.D{
// 	//		{"$in", bson.A{"not_started", "failed"}},
// 	//	}},
// 	//}

// 	ctx := context.Background()
// 	collection := db.Conn.Database("materi").Collection("skb")
// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return []Jabatan{}
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var result Jabatan
// 		err := cursor.Decode(&result)

// 		if err != nil {
// 			fmt.Println("error collecting data")
// 			return []Jabatan{}
// 		}

// 		jabatans = append(jabatans, result)
// 	}

// 	return jabatans
// }

// func CreateOrUpdateOutline(ctx context.Context, db *config.Database, job *Job, state string) error {
// 	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()

// 	collection := db.Conn.Database("materi").Collection("skb")

// 	filter := bson.D{
// 		{
// 			"job_id", job.Jabatan.JobID,
// 		},
// 	}

// 	update := bson.D{
// 		{
// 			"$set", bson.D{
// 				{"state", state},
// 				{"outline", job.Outline},
// 			},
// 		},
// 	}

// 	opts := options.Update().SetUpsert(true)

// 	_, err := collection.UpdateOne(ctx, filter, update, opts)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // for Create By Jabatan List
// func getListJabatanJSON() []string {
// 	type jsonFile struct {
// 		NamaJabatan string `json:"jabatan_name" bson:"jabatan_name"`
// 	}

// 	content, err := os.ReadFile("./files/task_generate.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var (
// 		res    []jsonFile
// 		result []string
// 	)
// 	json.Unmarshal(content, &res)

// 	for _, v := range res {
// 		a := v.NamaJabatan

// 		result = append(result, a)

// 	}

// 	return result
// }
