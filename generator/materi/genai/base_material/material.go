package base_material

import (
	"context"
	"encoding/json"
	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

type BaseMaterial struct {
	model llms.Model
}

func NewBaseMaterial(model llms.Model) *BaseMaterial {
	return &BaseMaterial{
		model: model,
	}
}

func (p *BaseMaterial) Generate(ctx context.Context, prompt MaterialPrompt) (MaterialGeneration, error) {
	var material MaterialGeneration

	tmpl := prompts.PromptTemplate{
		Template:       baseMaterialPromptTemplate,
		InputVariables: []string{"task"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	task, err := prompt.ToPrompt()
	if err != nil {
		return MaterialGeneration{}, err
	}

	executedPrompt, err := tmpl.Format(map[string]any{
		"task": task,
	})

	if err != nil {
		return MaterialGeneration{}, err
	}

	output, err := llms.GenerateFromSinglePrompt(ctx, p.model, executedPrompt)
	if err != nil {
		return MaterialGeneration{}, err
	}

	output, err = jsonrepair.RepairJSON(output)
	if err != nil {
		return MaterialGeneration{}, err
	}

	err = json.Unmarshal([]byte(output), &material)
	if err != nil {
		return MaterialGeneration{}, err
	}

	return material, nil
}
