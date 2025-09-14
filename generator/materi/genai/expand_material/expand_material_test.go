package expand_material_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"skb_materi/gemini"
	"skb_materi/genai/expand_material"
	"strings"
	"testing"
)

func TestNewExpandMaterial(t *testing.T) {
	godotenv.Load("../../.env")

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, os.Getenv("GOOGLE_API_KEY"))
	assert.NoError(t, err)

	exp := expand_material.NewExpandMaterial(model)

	prompt := expand_material.ExpandedMaterialPrompt{
		Purpose:      "Materi bahan ajar seleksi kompetensi bidang (SKB) calon aparatur sipil negara (seleksi CPNS) untuk formasi jabatan PRANATA KOMPUTER TERAMPIL",
		BaseMaterial: "Pengelolaan Sistem Informasi",
		SubMaterial:  "Analisis Kebutuhan Sistem",
		Learnings: []string{
			"Metodologi Analisis Kebutuhan",
			"Teknik Pengumpulan Data",
			"Dokumentasi Kebutuhan",
			"Prototyping",
			"Analisis SWOT",
		},
		MainTopics: []string{
			"Memahami konsep dasar dan tujuan analisis kebutuhan sistem informasi.",
			"Mampu mengidentifikasi dan menerapkan metodologi analisis kebutuhan yang tepat sesuai dengan karakteristik sistem.",
			"Menguasai teknik pengumpulan data yang relevan untuk memperoleh informasi kebutuhan sistem.",
			"Mampu mendokumentasikan kebutuhan sistem secara terstruktur dan komprehensif menggunakan alat bantu yang sesuai.",
			"Mampu menerapkan teknik prototyping untuk memperjelas dan memvalidasi kebutuhan sistem dengan stakeholders.",
			"Mampu melakukan analisis SWOT untuk mengevaluasi faktor-faktor internal dan eksternal yang mempengaruhi pengembangan sistem",
		},
		ChosenLearning: "Metodologi Analisis Kebutuhan",
		WordsToExpand:  "Menganalisis kebutuhan sistem adalah langkah yang sangat penting dalam siklus pengembangan sistem informasi. Proses ini bertujuan untuk mengidentifikasi kebutuhan dan masalah yang perlu diatasi melalui implementasi sistem yang baru atau yang sudah ada. Pada tahap ini kita akan mempelajari tentang berbagai macam metodologi analisis kebutuhan yang digunakan dalam pengembangan sistem informasi.",
	}

	out, err := exp.Generate(ctx, prompt)

	assert.NoError(t, err)

	fmt.Println(out.Expanded)

	fmt.Println("Words before: ", len(strings.Split(prompt.WordsToExpand, " ")))
	fmt.Println("Words after: ", len(strings.Split(out.Expanded, " ")))
}
