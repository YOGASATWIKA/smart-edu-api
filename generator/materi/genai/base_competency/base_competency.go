package base_competency

import (
	"context"
	"encoding/json"
	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

type BaseCompetency struct {
	model llms.Model
}

func NewBaseCompetency(model llms.Model) *BaseCompetency {
	return &BaseCompetency{
		model: model,
	}
}

func (b *BaseCompetency) Generate(ctx context.Context, prompt CompetencyPrompt) (CompetencyGenerated, error) {
	var competency CompetencyGenerated

	tmpl := prompts.PromptTemplate{
		Template:       baseCompetencyTemplate,
		InputVariables: []string{"task"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	task, err := prompt.ToPrompt()
	if err != nil {
		return competency, err
	}

	executedPrompt, err := tmpl.Format(map[string]any{
		"task": task,
	})
	if err != nil {
		return competency, err
	}

	output, err := llms.GenerateFromSinglePrompt(ctx, b.model, executedPrompt)
	if err != nil {
		return competency, err
	}

	output, err = jsonrepair.RepairJSON(output)
	if err != nil {
		return competency, err
	}

	err = json.Unmarshal([]byte(output), &competency)
	if err != nil {
		return competency, err
	}

	return competency, nil
}
