package base_material

import "fmt"

type MaterialPrompt struct {
	Purpose  string `json:"purpose"`
	Subject  string `json:"subject"`
	Chapter  string `json:"chapter"`
	Material string `json:"material"`
}

type MaterialGeneration struct {
	Short   string   `json:"short"`
	Details []string `json:"detail_materials"`
}

func (m MaterialPrompt) ToPrompt() (string, error) {

	return fmt.Sprintf(`Tujuan: %s
Materi Pokok: %s
Sub Materi Pokok: %s
Materi: %s`, m.Purpose, m.Subject, m.Chapter, m.Material), nil
}
