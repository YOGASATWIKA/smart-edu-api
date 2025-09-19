package base_competency_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"skb_materi/gemini"
	"skb_materi/genai/base_competency"
	"strings"
	"testing"
)

func TestBaseCompetency_Generate(t *testing.T) {
	godotenv.Load("./../../.env")

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, os.Getenv("GOOGLE_API_KEY"))
	assert.Nil(t, err)

	bc := base_competency.NewBaseCompetency(model)

	out, err := bc.Generate(ctx, base_competency.CompetencyPrompt{
		Purpose:      "Seleksi Kompetensi Bidang calon pegwai negeri sipili pada jabatan PRANATA KOMPUTER TERAMPIL",
		BaseMaterial: "Pengelolaan Sistem Informasi",
		SubMaterial:  "Analisis Kebutuhan Sistem",
		Learnings: []string{
			"Metodologi Analisis Kebutuhan",
			"Teknik Pengumpulan Data",
			"Dokumentasi Kebutuhan",
		},
	})

	assert.NoError(t, err)

	fmt.Println(strings.Join(out.TriggerQuestions, "\n"))
}
