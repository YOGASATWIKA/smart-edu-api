package background_urgency_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"skb_materi/gemini"
	"skb_materi/genai/background_urgency"
	"testing"
)

func TestNewBackgroundUrgency(t *testing.T) {
	godotenv.Load("../../.env")

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, os.Getenv("GOOGLE_API_KEY"))
	assert.NoError(t, err)

	bg := background_urgency.NewBackgroundUrgency(model)

	prompt := background_urgency.Prompt{
		Purpose:      "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan PRANATA KOMPUTER TERAMPIL",
		BaseMaterial: "Pengelolaan Sistem Informasi",
		SubMaterial:  "Analisis Kebutuhan Sistem",
		Materials: map[string]string{
			"Metodologi Analisis Kebutuhan": "Metodologi analisis kebutuhan sistem informasi adalah kerangka kerja terstruktur yang memandu proses identifikasi, pengumpulan, analisis, dan dokumentasi kebutuhan sistem. Memahami metodologi ini penting untuk memastikan pengembangan sistem yang efektif dan efisien sesuai dengan kebutuhan pengguna dan tujuan organisasi.",
			"Teknik Pengumpulan Data":       "Teknik pengumpulan data yang efektif sangat krusial dalam analisis kebutuhan sistem informasi, karena kualitas data yang diperoleh akan mempengaruhi kualitas analisis dan solusi yang dihasilkan. Pemilihan teknik yang tepat bergantung pada jenis data yang dibutuhkan dan karakteristik dari responden atau sumber data.",
			"Dokumentasi Kebutuhan":         "Dokumentasi kebutuhan sistem informasi adalah artefak penting yang merekam kebutuhan pengguna dan stakeholders secara terstruktur dan sistematis. Dokumentasi yang baik haruslah jelas, ringkas, lengkap, dan mudah dipahami untuk memfasilitasi komunikasi dan kolaborasi yang efektif dalam pengembangan sistem.",
			"Prototyping":                   "Prototyping dalam analisis kebutuhan sistem adalah membangun versi awal sistem yang disederhanakan untuk mendapatkan umpan balik dari pengguna, memperjelas kebutuhan, dan memvalidasi asumsi sebelum pengembangan skala penuh. Metode ini memfasilitasi komunikasi yang lebih baik antara pengembang dan pengguna.",
			"Analisis SWOT":                 "Analisis SWOT adalah alat yang membantu dalam memahami posisi sistem informasi saat ini dengan mengidentifikasi kekuatan (Strengths), kelemahan (Weaknesses), peluang (Opportunities), dan ancaman (Threats), sehingga mendukung dalam pengambilan keputusan strategis terkait pengembangan dan pengelolaan sistem informasi.",
		},
	}

	gen, err := bg.Generate(ctx, prompt)
	assert.NoError(t, err)

	fmt.Println(gen.Backgrounds)
	fmt.Println(gen.Urgencies)
}
