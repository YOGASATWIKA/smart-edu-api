package llm

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

func NewModel(apiKey string, selectedModel string) llms.Model {

	if selectedModel == "" {
		selectedModel = os.Getenv("MODEL")
	}

	godotenv.Load()
	model, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithModel(selectedModel),
		openai.WithBaseURL("https://openrouter.ai/api/v1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return model
}

func NewDefault(ctx context.Context, apiKey string) llms.Model {
	model, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultTemperature(1),
		googleai.WithHarmThreshold(googleai.HarmBlockNone),
		googleai.WithDefaultModel("gemini-2.0-flash"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return model
}
