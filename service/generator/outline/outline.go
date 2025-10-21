package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"smart-edu-api/entity"
	"smart-edu-api/repository"
	"strings"
	"time"

	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
)

type Params struct {
	NamaJabatan  string
	TugasJabatan string
	Keterampilan string
	Model        string
}

type Outliner struct {
	model llms.Model
}

func New(model llms.Model) *Outliner {
	return &Outliner{
		model: model,
	}
}

func (o *Outliner) Generate(ctx context.Context, params Params) (entity.Outline, error) {
	err := prepareLogs()
	if err != nil {
		return entity.Outline{}, err
	}

	var out entity.Outline

	promt, err := repository.GetModelByModel(params.Model)
	if err != nil {
		return entity.Outline{}, fmt.Errorf("Model Not Found: %w", err)
	}

	data := map[string]string{
		"NamaJabatan":  params.NamaJabatan,
		"TugasJabatan": params.TugasJabatan,
		"Keterampilan": params.Keterampilan,
	}

	var contents []llms.MessageContent
	var continues []llms.MessageContent
	for i, step := range promt.Steps {
		tmpl, err := template.New("prompt").Parse(step.Content)
		if err != nil {
			fmt.Println("Error parse template:", err)
			fmt.Println("Template content:", step.Content)
			return out, err
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return out, err
		}

		switch step.Role {
		case "system":
			contents = append(contents, llms.TextParts(llms.ChatMessageTypeSystem, buf.String()))
		case "human":
			continues = append(continues, llms.TextParts(llms.ChatMessageTypeHuman, buf.String()))
		}
		if i == len(promt.Steps)-1 {
			continues = append(continues, llms.TextParts(llms.ChatMessageTypeHuman, "\"Buatkan hasilnya dalam format JSON seperti berikut:\\\\n{\\\\n\\\\\\\"list_materi\\\\\\\": [\\\\n{\\\\n\\\\\\\"materi_pokok\\\\\\\": \\\\\\\"\\\\\\\",\\\\n\\\\\\\"list_sub_materi\\\\\\\": [\\\\n{\\\\n\\\\\\\"sub_materi_pokok\\\\\\\": \\\\\\\"\\\\\\\",\\\\n\\\\\\\"list_materi\\\\\\\": [\\\\\\\"\\\\\\\", \\\\\\\"\\\\\\\"]\\\\n}\\\\n]\\\\n}\\\\n]\\\\n}\\\"\""))
		}

	}

	var lastResponse *llms.ContentResponse
	generateStep := 1
	for _, content := range continues {
		log.Println(params.NamaJabatan, "generating step", generateStep, "...")
		contents = append(contents, content)

		res, err := o.model.GenerateContent(ctx, contents, llms.WithJSONMode(), llms.WithMaxTokens(6000))
		if err != nil {
			return out, err
		}

		lastResponse = res

		contents = append(contents, llms.TextParts(llms.ChatMessageTypeAI, res.Choices[0].Content))
		generateStep++
	}

	if lastResponse == nil {
		return out, fmt.Errorf("empty response")
	}

	clean, err := jsonrepair.RepairJSON(lastResponse.Choices[0].Content)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal([]byte(clean), &out); err != nil {
		return out, err
	}

	_ = saveLogs(params.NamaJabatan, contents)
	return out, nil
}

func prepareLogs() error {
	err := os.MkdirAll("./logs", 0775)
	if err != nil {
		return err
	}
	return nil
}

func saveLogs(name string, contents []llms.MessageContent) error {
	var textBuilder strings.Builder
	textBuilder.WriteString("Nama Jabatan: " + name + "\n\n")

	for _, content := range contents {
		var contentText string
		if len(content.Parts) > 0 {
			if partStr, ok := content.Parts[0].(llms.TextContent); ok {
				contentText = partStr.Text
			}
		}
		textBuilder.WriteString(fmt.Sprintf("Role: %s\nContent: %s\n\n", content.Role, contentText))
	}

	filename := fmt.Sprintf("./logs/%s-%s.txt", name, time.Now().Format("2006-01-02_15-04-05"))

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(textBuilder.String())
	return err
}
