package base_material_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"skb_materi/gemini"
	base_material2 "skb_materi/genai/base_material"
	"testing"
)

func TestBaseMaterial_Generate(t *testing.T) {
	godotenv.Load("../../.env")

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, os.Getenv("GOOGLE_API_KEY"))
	assert.NoError(t, err)

	bm := base_material2.NewBaseMaterial(model)

	mg, err := bm.Generate(ctx, base_material2.MaterialPrompt{
		Purpose:  "Seleksi Kompetensi Bidang untuk calon pegawai negeri untuk jabatan Pranana Komputer Terampil",
		Subject:  "Pengelolaan Sistem Informasi",
		Chapter:  "Analisis Kebutuhan Sistem",
		Material: "Metodologi Analisis Kebutuhan",
	})

	assert.NoError(t, err)

	fmt.Println(mg.Short)
	fmt.Println(mg.Details)
}
