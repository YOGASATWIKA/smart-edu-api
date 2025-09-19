package background_urgency

import (
	"context"
	"encoding/json"

	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

type Intro struct {
	model llms.Model
}

func NewBackgroundUrgency(model llms.Model) *Intro {
	return &Intro{
		model: model,
	}
}

func (p *Intro) Generate(ctx context.Context, prompt Prompt) (Generated, error) {
	var gen Generated

	tmpl := prompts.PromptTemplate{
		Template:       template,
		InputVariables: []string{"task"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	task, err := prompt.ToPrompt()
	if err != nil {
		return gen, err
	}

	finalPrompt, err := tmpl.Format(map[string]any{
		"task": task,
	})
	if err != nil {
		return gen, err
	}

	out, err := llms.GenerateFromSinglePrompt(ctx, p.model, finalPrompt)
	if err != nil {
		return gen, err
	}

	out, err = jsonrepair.RepairJSON(out)
	if err != nil {
		return gen, err
	}

	err = json.Unmarshal([]byte(out), &gen)
	if err != nil {
		return gen, err
	}

	return gen, nil
}
