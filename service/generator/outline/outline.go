//package generator
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"os"
//	"smart-edu-api/entity"
//	"time"
//
//	jsonrepair "github.com/RealAlexandreAI/json-repair"
//	"github.com/tmc/langchaingo/llms"
//)
//
//type Params struct {
//	NamaJabatan  string
//	TugasJabatan string
//	Keterampilan string
//	Model        string
//}
//
//type Outliner struct {
//	model llms.Model
//}
//
//func New(model llms.Model) *Outliner {
//	return &Outliner{
//		model: model,
//	}
//}
//func (o *Outliner) Generate(model string, ctx context.Context, params Params) (entity.Outline, error) {
//	err := prepareLogs()
//	if err != nil {
//		return entity.Outline{}, err
//	}
//
//	var out = entity.Outline{}
//
//	contents := []llms.MessageContent{
//		llms.TextParts(llms.ChatMessageTypeSystem, "Anda adalah pejabat yang memiliki kompetensi dan pengalamaan dalam menyusun materi seleksi kompetensi bidang/teknis untuk jabatan dipemerintahan. Berdasarkan kompetensi dan pengalaman anda, selesaikan tugas berikut"),
//	}
//
//	continues := []llms.MessageContent{
//		llms.TextParts(llms.ChatMessageTypeHuman, fmt.Sprintf("Apa kompetensi bidang/teknis yang harus dimiliki oleh seseorang dengan\nNama Jabatan: %s \nTugas Jabatan: %s \nKeterampilan: %s\n\nIngat, Kompetensi bidang/teknis dalam seleksi ASN (Aparatur Sipil Negara) merujuk pada kemampuan atau pengetahuan spesifik yang berkaitan langsung dengan tugas dan fungsi jabatan yang dilamar. Kompetensi ini mencakup keterampilan teknis, pengetahuan, dan pengalaman yang diperlukan untuk melaksanakan pekerjaan tertentu secara efektif.", params.NamaJabatan, params.TugasJabatan, params.Keterampilan)),
//		llms.TextParts(llms.ChatMessageTypeHuman, "Berdasarkan kompetensi bidang/teknis tersebut, apa materi pokok yang perlu dipelajari dan dipahami agar dapat melaksanakan tugas jabatan secara baik dan benar. Buatkan minimal 5 materi pokok dimana masing-masing materi pokok berisi 5 sub materi pokok sertakan deskripsi yang mendetail tentang apa yang harus di perhatikan dalam penyusunan materi_pokok dan sub_materi_pokok."),
//		llms.TextParts(llms.ChatMessageTypeHuman, "Buatkan format dalam bentuk JSON sesuai template dibawah\n\n{\n\t\t\"list_materi\": [\n\t\t\t{\n\t\t\t\t\"materi_pokok\": \"\",\n\t\t\t\t\"list_sub_materi\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"sub_materi_pokok\": \"\",\n\t\t\t\t\t\t\"list_materi\": [\"\", \"\"]\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t}\n\t\t]\n\t }\n"),
//	}
//
//	var lastResponse *llms.ContentResponse
//
//	step := 1
//
//	for _, content := range continues {
//		log.Println(params.NamaJabatan, "generating step", step, "...")
//		contents = append(contents, content)
//
//		res, err := o.model.GenerateContent(ctx, contents, llms.WithJSONMode(), llms.WithMaxTokens(6000))
//		if err != nil {
//			return out, err
//		}
//
//		lastResponse = res
//
//		contents = append(contents, llms.TextParts(llms.ChatMessageTypeAI, res.Choices[0].Content))
//		step++
//	}
//
//	if lastResponse == nil {
//		return out, fmt.Errorf("empty response")
//	}
//
//	clean, err := jsonrepair.RepairJSON(lastResponse.Choices[0].Content)
//	if err != nil {
//		return out, err
//	}
//
//	err = json.Unmarshal([]byte(clean), &out)
//	if err != nil {
//		return out, err
//	}
//
//	_ = saveLogs(params.NamaJabatan, contents)
//
//	return out, nil
//}
//
//func prepareLogs() error {
//	err := os.MkdirAll("./logs", 0775)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//func saveLogs(name string, contents []llms.MessageContent) error {
//	text := "Nama Jabatan: " + name + "\n"
//
//	for _, content := range contents {
//		text = fmt.Sprintf("%s\nRole: %s\nContent: %s\n\n", text, content.Role, content.Parts[0])
//	}
//
//	filename := fmt.Sprintf("./logs/%s-%s.txt", name, time.Now().Format("2006-01-02_15-04-05"))
//
//	create, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//
//	defer create.Close()
//
//	_, err = create.Write([]byte(text))
//
//	return nil
//}

//Modifed

package generator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"smart-edu-api/entity"
	"smart-edu-api/repository"
	"strings"
	"time"

	jsonrepair "github.com/RealAlexandreAI/json-repair"
	"github.com/tmc/langchaingo/llms"
)

type Params struct {
	NamaJabatan  string
	TugasJabatan string
	Keterampilan string
	Model        string
}

type Outliner struct {
	model llms.Model
}

func New(model llms.Model) *Outliner {
	return &Outliner{
		model: model,
	}
}

// ==================================================================Promt In Code
//func (o *Outliner) Generate(ctx context.Context, params Params) (entity.Outline, error) {
//	err := prepareLogs()
//	if err != nil {
//		return entity.Outline{}, err
//	}
//
//	var out = entity.Outline{}
//
//	contents := []llms.MessageContent{
//		llms.TextParts(llms.ChatMessageTypeSystem, "Anda adalah pejabat yang memiliki kompetensi dan pengalamaan dalam menyusun materi seleksi kompetensi bidang/teknis untuk jabatan dipemerintahan. Berdasarkan kompetensi dan pengalaman anda, selesaikan tugas berikut"),
//	}
//
//	continues := []llms.MessageContent{
//		llms.TextParts(llms.ChatMessageTypeHuman, fmt.Sprintf("Apa kompetensi bidang/teknis yang harus dimiliki oleh seseorang dengan\nNama Jabatan: %s \nTugas Jabatan: %s \nKeterampilan: %s\n\nIngat, Kompetensi bidang/teknis dalam seleksi ASN (Aparatur Sipil Negara) merujuk pada kemampuan atau pengetahuan spesifik yang berkaitan langsung dengan tugas dan fungsi jabatan yang dilamar. Kompetensi ini mencakup keterampilan teknis, pengetahuan, dan pengalaman yang diperlukan untuk melaksanakan pekerjaan tertentu secara efektif.", params.NamaJabatan, params.TugasJabatan, params.Keterampilan)),
//		llms.TextParts(llms.ChatMessageTypeHuman, "Berdasarkan kompetensi bidang/teknis tersebut, apa materi pokok yang perlu dipelajari dan dipahami agar dapat melaksanakan tugas jabatan secara baik dan benar. Buatkan minimal 5 materi pokok dimana masing-masing materi pokok berisi 5 sub materi pokok."),
//		llms.TextParts(llms.ChatMessageTypeHuman, "Buatkan format dalam bentuk JSON sesuai template dibawah\n\n{\n\t\t\"list_materi\": [\n\t\t\t{\n\t\t\t\t\"materi_pokok\": \"\",\n\t\t\t\t\"list_sub_materi\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"sub_materi_pokok\": \"\",\n\t\t\t\t\t\t\"list_materi\": [\"\", \"\"]\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t}\n\t\t]\n\t }\n"),
//	}
//
//	var lastResponse *llms.ContentResponse
//
//	step := 1
//
//	for _, content := range continues {
//		log.Println(params.NamaJabatan, "generating step", step, "...")
//		contents = append(contents, content)
//
//		res, err := o.model.GenerateContent(ctx, contents, llms.WithJSONMode(), llms.WithMaxTokens(6000))
//		if err != nil {
//			return out, err
//		}
//
//		lastResponse = res
//
//		contents = append(contents, llms.TextParts(llms.ChatMessageTypeAI, res.Choices[0].Content))
//		step++
//	}
//
//	if lastResponse == nil {
//		return out, fmt.Errorf("empty response")
//	}
//
//	clean, err := jsonrepair.RepairJSON(lastResponse.Choices[0].Content)
//	if err != nil {
//		return out, err
//	}
//
//	err = json.Unmarshal([]byte(clean), &out)
//	if err != nil {
//		return out, err
//	}
//
//	_ = saveLogs(params.NamaJabatan, contents)
//
//	return out, nil
//}

func (o *Outliner) Generate(ctx context.Context, params Params) (entity.Outline, error) {
	err := prepareLogs()
	if err != nil {
		return entity.Outline{}, err
	}

	var out entity.Outline

	promt, err := repository.GetModelByModel(params.Model)
	if err != nil {
		return entity.Outline{}, fmt.Errorf("gagal mendapatkan model prompt dari repository: %w", err)
	}

	data := map[string]string{
		"NamaJabatan":  params.NamaJabatan,
		"TugasJabatan": params.TugasJabatan,
		"Keterampilan": params.Keterampilan,
	}

	var contents []llms.MessageContent
	for _, step := range promt.Steps {
		tmpl, err := template.New("prompt").Parse(step.Content)
		if err != nil {
			fmt.Println("Error parse template:", err)
			fmt.Println("Template content:", step.Content)
			return out, err
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return out, err
		}

		switch step.Role {
		case "system":
			contents = append(contents, llms.TextParts(llms.ChatMessageTypeSystem, buf.String()))
		case "human":
			contents = append(contents, llms.TextParts(llms.ChatMessageTypeHuman, buf.String()))
		}
	}

	res, err := o.model.GenerateContent(ctx, contents, llms.WithJSONMode(), llms.WithMaxTokens(6000))
	if err != nil {
		return out, err
	}

	if len(res.Choices) == 0 {
		return out, fmt.Errorf("empty response from model")
	}

	clean, err := jsonrepair.RepairJSON(res.Choices[0].Content)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal([]byte(clean), &out); err != nil {
		return out, err
	}

	_ = saveLogs(params.NamaJabatan, contents)
	return out, nil
}

//==================================================================Promt In Code

//==================================================================Promt In Database
//func (o *Outliner) Generate(model string, ctx context.Context, params Params) (entity.Outline, error) {
//	err := prepareLogs()
//	if err != nil {
//		return entity.Outline{}, err
//	}
//
//	var out = entity.Outline{}
//
//	existing, err := repository.GetModelByModel(model)
//	if err != nil {
//		return entity.Outline{}, fmt.Errorf("gagal mendapatkan model prompt dari repository: %w", err)
//	}
//
//	var promptSteps []string
//
//	contextPrompt := fmt.Sprintf("Apa kompetensi bidang/teknis yang harus dimiliki oleh seseorang dengan\nNama Jabatan: %s \nTugas Jabatan: %s \nKeterampilan: %s\n\nIngat, Kompetensi bidang/teknis dalam seleksi ASN (Aparatur Sipil Negara) merujuk pada kemampuan atau pengetahuan spesifik yang berkaitan langsung dengan tugas dan fungsi jabatan yang dilamar. Kompetensi ini mencakup keterampilan teknis, pengetahuan, dan pengalaman yang diperlukan untuk melaksanakan pekerjaan tertentu secara efektif.", params.NamaJabatan, params.TugasJabatan, params.Keterampilan)
//	promptSteps = append(promptSteps, contextPrompt)
//	promptSteps = append(promptSteps, existing.Promt.UserPrompts...)
//
//		llms.TextParts(llms.ChatMessageTypeHuman, "Berdasarkan kompetensi bidang/teknis tersebut, apa materi pokok yang perlu dipelajari dan dipahami agar dapat melaksanakan tugas jabatan secara baik dan benar. Buatkan minimal 5 materi pokok dimana masing-masing materi pokok berisi 5 sub materi pokok sertakan deskripsi yang mendetail tentang apa yang harus di perhatikan dalam penyusunan materi_pokok dan sub_materi_pokok."),
//
//	jsonFormatPrompt := "Buatkan format dalam bentuk JSON sesuai template dibawah\n\n{\n\t\t\"list_materi\": [\n\t\t\t{\n\t\t\t\t\"materi_pokok\": \"\",\n\t\t\t\t\"list_sub_materi\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"sub_materi_pokok\": \"\",\n\t\t\t\t\t\t\"list_materi\": [\"\", \"\"]\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t}\n\t\t]\n\t }\n"
//	promptSteps = append(promptSteps, jsonFormatPrompt)
//
//	conversationHistory := []llms.MessageContent{
//		llms.TextParts(llms.ChatMessageTypeSystem, fmt.Sprintf (existing.Promt.SystemPrompt)),
//		//llms.TextParts(llms.ChatMessageTypeHuman, fmt.Sprintf("Apa kompetensi bidang/teknis yang harus dimiliki oleh seseorang dengan\nNama Jabatan: %s \nTugas Jabatan: %s \nKeterampilan: %s\n\nIngat, Kompetensi bidang/teknis dalam seleksi ASN (Aparatur Sipil Negara) merujuk pada kemampuan atau pengetahuan spesifik yang berkaitan langsung dengan tugas dan fungsi jabatan yang dilamar. Kompetensi ini mencakup keterampilan teknis, pengetahuan, dan pengalaman yang diperlukan untuk melaksanakan pekerjaan tertentu secara efektif.", params.NamaJabatan, params.TugasJabatan, params.Keterampilan)),
//
//	}
//
//	var lastResponse *llms.ContentResponse
//	step := 1
//
//	for _, stepPrompt := range promptSteps {
//		log.Println(params.NamaJabatan, "generating step", step, "...")
//
//		conversationHistory = append(conversationHistory, llms.TextParts(llms.ChatMessageTypeHuman, stepPrompt))
//
//		res, err := o.model.GenerateContent(ctx, conversationHistory, llms.WithJSONMode(), llms.WithMaxTokens(6000))
//		if err != nil {
//			return out, fmt.Errorf("error pada langkah %d saat generate content: %w", step, err)
//		}
//
//		lastResponse = res
//
//		if len(res.Choices) > 0 {
//			conversationHistory = append(conversationHistory, llms.TextParts(llms.ChatMessageTypeAI, res.Choices[0].Content))
//		} else {
//			log.Println("Peringatan: AI tidak memberikan konten pada langkah", step)
//		}
//		step++
//	}
//
//	if lastResponse == nil || len(lastResponse.Choices) == 0 {
//		return out, fmt.Errorf("respons akhir dari AI kosong")
//	}
//
//	clean, err := jsonrepair.RepairJSON(lastResponse.Choices[0].Content)
//	if err != nil {
//		return out, fmt.Errorf("gagal memperbaiki JSON dari respons AI: %w", err)
//	}
//
//	if err := json.Unmarshal([]byte(clean), &out); err != nil {
//		return out, fmt.Errorf("gagal unmarshal JSON yang sudah diperbaiki: %w", err)
//	}
//
//	_ = saveLogs(params.NamaJabatan, conversationHistory)
//
//	return out, nil
//}

func prepareLogs() error {
	err := os.MkdirAll("./logs", 0775)
	if err != nil {
		return err
	}
	return nil
}

func saveLogs(name string, contents []llms.MessageContent) error {
	var textBuilder strings.Builder
	textBuilder.WriteString("Nama Jabatan: " + name + "\n\n")

	for _, content := range contents {
		var contentText string
		// Ambil bagian pertama dari Parts, karena TextParts hanya mengisi satu bagian.
		if len(content.Parts) > 0 {
			if partStr, ok := content.Parts[0].(llms.TextContent); ok {
				contentText = partStr.Text
			}
		}
		textBuilder.WriteString(fmt.Sprintf("Role: %s\nContent: %s\n\n", content.Role, contentText))
	}

	filename := fmt.Sprintf("./logs/%s-%s.txt", name, time.Now().Format("2006-01-02_15-04-05"))

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(textBuilder.String())
	return err
}
