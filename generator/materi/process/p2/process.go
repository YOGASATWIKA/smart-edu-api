package p2

import (
	"context"
	"log"
	"smart-edu-api/entity"
	"smart-edu-api/generator/materi/genai/base_competency"
)

type Fetch struct {
	bc *base_competency.BaseCompetency
}

func NewFetch(bc *base_competency.BaseCompetency) *Fetch {
	return &Fetch{
		bc: bc,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *entity.Ebook) {
	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {
			var materials []string

			for _, material := range chapter.Materials {
				materials = append(materials, material.Title)
			}

			try := 1

			for {
				log.Println("Generate objective & trigger question for ", chapter.Title, "with try : ", try)

				out, err := f.bc.Generate(ctx, base_competency.CompetencyPrompt{
					Purpose:      "Seleksi Kompetensi Bidang calon pegwai negeri sipili pada jabatan " + ebook.Title,
					BaseMaterial: part.Subject,
					SubMaterial:  chapter.Title,
					Learnings:    materials,
				})

				if err != nil {
					try++
					continue
				}

				chapter.BaseCompetitions = out.Objectives
				chapter.TriggerQuestions = out.TriggerQuestions

				break
			}
		}
	}
}
