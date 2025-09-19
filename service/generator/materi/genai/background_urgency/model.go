package background_urgency

import "fmt"

type Prompt struct {
	Purpose      string              `json:"purpose"`
	BaseMaterial string              `json:"base_material"`
	Competencies map[string][]string `json:"materials"`
}

type Generated struct {
	Backgrounds []string `json:"backgrounds"`
	Urgencies   []string `json:"urgencies"`
}

func (s Prompt) ToPrompt() (string, error) {
	text := ""
	index := 1
	for k, v := range s.Competencies {
		text += fmt.Sprintf("%d. %s\n", index, k)
		ntxt := ""

		for _, n := range v {
			ntxt += fmt.Sprintf("- %s\n", n)
		}

		text += fmt.Sprintf("%s\n\n", ntxt)
		index++
	}

	return fmt.Sprintf(`
Tujuan Materi: %s
Materi Pokok: %s
Kompetensi Dasar:
%s
`, s.Purpose, s.BaseMaterial, text), nil
}
