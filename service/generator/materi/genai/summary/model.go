package summary

import "fmt"

type Prompt struct {
	Purpose      string            `json:"purpose"`
	BaseMaterial string            `json:"base_material"`
	SubMaterial  string            `json:"sub_material"`
	Materials    map[string]string `json:"materials"`
}

type Generated struct {
	Summary     string   `json:"summary"`
	Reflections []string `json:"reflections"`
}

func (s Prompt) ToPrompt() (string, error) {
	text := ""
	index := 1
	for k, v := range s.Materials {
		text += fmt.Sprintf("%d. %s\n", index, k)
		text += fmt.Sprintf("%s\n\n", v)
		index++
	}

	return fmt.Sprintf(`
Tujuan Materi: %s
Materi Pokok: %s
Sub Materi: %s
Materi yang dibahas:
%s
`, s.Purpose, s.BaseMaterial, s.SubMaterial, text), nil
}
