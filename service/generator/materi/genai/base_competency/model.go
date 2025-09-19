package base_competency

import (
	"fmt"
	"strings"
)

type CompetencyPrompt struct {
	Purpose      string   `json:"purpose"`
	BaseMaterial string   `json:"base_material"`
	SubMaterial  string   `json:"sub_material"`
	Learnings    []string `json:"learnings"`
}

type CompetencyGenerated struct {
	Objectives       []string `json:"objectives"`
	TriggerQuestions []string `json:"trigger_questions"`
}

func (cp CompetencyPrompt) ToPrompt() (string, error) {
	return fmt.Sprintf(`Tujuan Materi: %s
Materi Pokok: %s
Sub Materi: %s
Pembelajaran:
%s
`, cp.Purpose, cp.BaseMaterial, cp.SubMaterial, strings.Join(cp.Learnings, "\n")), nil
}
