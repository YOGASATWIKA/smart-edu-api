package expand_material

import (
	"fmt"
	"strings"
)

type ExpandedMaterialPrompt struct {
	Purpose        string   `json:"purpose"`
	BaseMaterial   string   `json:"base_material"`
	SubMaterial    string   `json:"sub_material"`
	Learnings      []string `json:"learnings"`
	MainTopics     []string `json:"main_topics"`
	ChosenLearning string   `json:"chosen_learning"`
	WordsToExpand  string   `json:"words_to_expand"`
}

type ExpandedMaterialGenerated struct {
	Expanded string `json:"expanded"`
}

func (e ExpandedMaterialPrompt) ToPrompt() (string, error) {
	return fmt.Sprintf(`
Tujuan: %s
Materi Pokok: %s
Sub Materi: %s
Pembelajaran:
%s

Pokok Bahasan:
%s

Pembelajaran yang ingin dikembangakan: %s
Kalimat yang ingin dikembangkan:
%s
`, e.Purpose, e.BaseMaterial, e.SubMaterial, strings.Join(e.Learnings, "\n"), strings.Join(e.MainTopics, "\n"), e.ChosenLearning, e.WordsToExpand), nil
}
