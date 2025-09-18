package query

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"smart-edu-api/data/materi/query"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	processChannels = make(map[string]chan query.StatusMessage)
	mu              sync.Mutex
)

// query/get_streaming_materi.go

func GetStreamingMateri(c *fiber.Ctx) error {
	id := c.Params("id")

	mu.Lock()
	updateChan, ok := processChannels[id]
	mu.Unlock()

	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Process not found"})
	}
	defer func() {
		mu.Lock()
		delete(processChannels, id)
		mu.Unlock()
		log.Printf("Cleaned up resources for process ID: %s", id)
	}()

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		log.Printf("Client connected for streaming status of ID: %s", id)

		for msg := range updateChan {
			jsonData, _ := json.Marshal(msg)
			fmt.Fprintf(w, "data: %s\n\n", jsonData)
			w.Flush()
		}
		log.Printf("Streaming finished for ID: %s", id)
	})

	return nil
}
