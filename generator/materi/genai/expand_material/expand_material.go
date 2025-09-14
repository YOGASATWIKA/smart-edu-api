package expand_material

import (
	"context"
	"encoding/json"
	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

type ExpandMaterial struct {
	model llms.Model
}

func NewExpandMaterial(model llms.Model) *ExpandMaterial {
	return &ExpandMaterial{model: model}
}

func (p *ExpandMaterial) Generate(ctx context.Context, prompt ExpandedMaterialPrompt) (ExpandedMaterialGenerated, error) {
	var exp ExpandedMaterialGenerated

	tmpl := prompts.PromptTemplate{
		Template:       promptTemplate,
		InputVariables: []string{"task"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	task, err := prompt.ToPrompt()
	if err != nil {
		return exp, err
	}

	finalPrompt, err := tmpl.Format(map[string]any{
		"task": task,
	})
	if err != nil {
		return exp, err
	}

	out, err := llms.GenerateFromSinglePrompt(ctx, p.model, finalPrompt)
	if err != nil {
		return exp, err
	}

	out, err = jsonrepair.RepairJSON(out)
	if err != nil {
		return exp, err
	}

	err = json.Unmarshal([]byte(out), &exp)
	if err != nil {
		return exp, err
	}

	return exp, nil
}
