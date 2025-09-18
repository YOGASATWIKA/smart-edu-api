package p4

import (
	"context"
	"log"
	"smart-edu-api/entity"
	"smart-edu-api/generator/materi/genai/background_urgency"
)

type Fetch struct {
	bg *background_urgency.Intro
}

func NewFetch(bu *background_urgency.Intro) *Fetch {
	return &Fetch{
		bg: bu,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *entity.Ebook) {
	for _, part := range ebook.Parts {
		var competencies = make(map[string][]string)

		for _, chapter := range part.Chapters {
			competencies[chapter.Title] = chapter.BaseCompetitions
		}

		try := 1

		for {

			log.Println("Generate background & urgencies for: ", part.Subject, " with try: ", try)

			out, err := f.bg.Generate(ctx, background_urgency.Prompt{
				Purpose:      "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan " + ebook.Title,
				BaseMaterial: part.Subject,
				Competencies: competencies,
			})

			if err != nil {
				try++
				continue
			}

			part.Introductions = out.Backgrounds
			part.Urgencies = out.Urgencies

			break
		}
	}
}
