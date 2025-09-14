package summary

import (
	"context"
	"encoding/json"
	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

type Summary struct {
	model llms.Model
}

func NewSummary(model llms.Model) *Summary {
	return &Summary{model: model}
}

func (p *Summary) Generate(ctx context.Context, prompt Prompt) (Generated, error) {
	var generated Generated

	tmpl := prompts.PromptTemplate{
		Template:       template,
		InputVariables: []string{"task"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	task, err := prompt.ToPrompt()
	if err != nil {
		return generated, err
	}

	finalPrompt, err := tmpl.Format(map[string]any{
		"task": task,
	})
	if err != nil {
		return generated, err
	}

	out, err := llms.GenerateFromSinglePrompt(ctx, p.model, finalPrompt)
	if err != nil {
		return generated, err
	}

	out, err = jsonrepair.RepairJSON(out)
	if err != nil {
		return generated, err
	}

	err = json.Unmarshal([]byte(out), &generated)
	if err != nil {
		return generated, err
	}

	return generated, nil
}
