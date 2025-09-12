package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	outline "smart-edu-api/data/outline/request"
	"time"

	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
)

type Params struct {
	NamaJabatan  string
	TugasJabatan string
	Keterampilan string
}

type Outliner struct {
	model llms.Model
}

func New(model llms.Model) *Outliner {
	return &Outliner{
		model: model,
	}
}
func (o *Outliner) GenerateWithOfficialMaterial(ctx context.Context, params Params) (outline.Outline, error) {
	err := prepareLogs()
	if err != nil {
		return outline.Outline{}, err
	}

	var out = outline.Outline{}

	contents := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "Anda adalah pejabat yang memiliki kompetensi dan pengalamaan dalam menyusun materi seleksi kompetensi bidang/teknis untuk jabatan dipemerintahan. Berdasarkan kompetensi dan pengalaman anda, selesaikan tugas berikut"),
	}

	continues := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, fmt.Sprintf("Apa kompetensi bidang/teknis yang harus dimiliki oleh seseorang dengan\nNama Jabatan: %s\nMateri Pokok: %s\n\nIngat, Kompetensi bidang/teknis dalam seleksi ASN (Aparatur Sipil Negara) merujuk pada kemampuan atau pengetahuan spesifik yang berkaitan langsung dengan tugas dan fungsi jabatan yang dilamar. Kompetensi ini mencakup keterampilan teknis, pengetahuan, dan pengalaman yang diperlukan untuk melaksanakan pekerjaan tertentu secara efektif.", params.NamaJabatan, params.TugasJabatan)),
		llms.TextParts(llms.ChatMessageTypeHuman, "Berdasarkan kompetensi bidang/teknis tersebut, apa sub materi pokok yang perlu dipelajari dan dipahami agar dapat melaksanakan tugas jabatan secara baik dan benar. Buatkan masing-masing materi pokok berisi 5 sub materi pokok."),
		llms.TextParts(llms.ChatMessageTypeHuman, "Buatkan format dalam bentuk JSON sesuai template dibawah\n\n{\n\t\t\"list_materi\": [\n\t\t\t{\n\t\t\t\t\"materi_pokok\": \"\",\n\t\t\t\t\"list_sub_materi\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"sub_materi_pokok\": \"\",\n\t\t\t\t\t\t\"list_materi\": [\"\", \"\"]\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t}\n\t\t]\n\t }\n"),
	}

	var lastResponse *llms.ContentResponse

	step := 1

	for _, content := range continues {
		log.Println(params.NamaJabatan, "generating step", step, "...")
		contents = append(contents, content)

		res, err := o.model.GenerateContent(ctx, contents, llms.WithJSONMode(), llms.WithMaxTokens(6000))
		if err != nil {
			return out, err
		}

		lastResponse = res

		contents = append(contents, llms.TextParts(llms.ChatMessageTypeAI, res.Choices[0].Content))
		step++
	}

	if lastResponse == nil {
		return out, fmt.Errorf("empty response")
	}

	clean, err := jsonrepair.RepairJSON(lastResponse.Choices[0].Content)
	if err != nil {
		return out, err
	}

	err = json.Unmarshal([]byte(clean), &out)
	if err != nil {
		return out, err
	}

	_ = saveLogs(params.NamaJabatan, contents)

	return out, nil
}

func prepareLogs() error {
	err := os.MkdirAll("./logs", 0775)
	if err != nil {
		return err
	}

	return nil
}
func saveLogs(name string, contents []llms.MessageContent) error {
	text := "Nama Jabatan: " + name + "\n"

	for _, content := range contents {
		text = fmt.Sprintf("%s\nRole: %s\nContent: %s\n\n", text, content.Role, content.Parts[0])
	}

	filename := fmt.Sprintf("./logs/%s-%s.txt", name, time.Now().Format("2006-01-02_15-04-05"))

	create, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer create.Close()

	_, err = create.Write([]byte(text))

	return nil
}
