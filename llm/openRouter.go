package llm

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/tmc/langchaingo/llms"
// )

// type OpenRouterClient struct {
// 	apiKey string
// 	model  string
// }

// func New(ctx context.Context, apiKey string) llms.Model {
// 	model, err := 

// }

// func (c *OpenRouterClient) Chat(ctx context.Context, message string) (string, error) {
// 	url := "https://openrouter.ai/api/v1/chat/completions"

// 	payload := map[string]interface{}{
// 		"model": c.model,
// 		"messages": []map[string]string{
// 			{"role": "user", "content": message},
// 		},
// 	}

// 	data, _ := json.Marshal(payload)

// 	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
// 	if err != nil {
// 		return "", err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+c.apiKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	var result map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "", err
// 	}

// 	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
// 		if msg, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
// 			return msg["content"].(string), nil
// 		}
// 	}

// 	return "", fmt.Errorf("no response content")
// }
