package llm

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func New(ctx context.Context, apiKey string) llms.Model {
	godotenv.Load()
	model, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithModel(os.Getenv("MODEL")),
		openai.WithBaseURL("https://openrouter.ai/api/v1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return model
}
