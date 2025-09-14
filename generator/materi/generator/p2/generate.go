package p2

import (
	"context"
	"log"
	"skb_materi/genai/base_competency"
	"skb_materi/types"
)

type PurposeType int

const (
	SkbType PurposeType = iota
	PppkType
	PppkWawancaraType
)

type Fetch struct {
	bc *base_competency.BaseCompetency
}

func NewFetch(bc *base_competency.BaseCompetency) *Fetch {
	return &Fetch{
		bc: bc,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *types.Ebook, purposeType PurposeType) {
	var purposeText = ""

	if purposeType == SkbType {
		purposeText = "Seleksi Kompetensi Bidang calon pegawai negeri sipil pada jabatan " + ebook.Title
	} else if purposeType == PppkType {
		purposeText = "Seleksi Kompetensi Teknis calon pegawai negeri sipil pada jabatan " + ebook.Title
	} else if purposeType == PppkWawancaraType {
		purposeText = "Seleksi Wawancara calon pegawai negeri sipili pada jabatan " + ebook.Title
	}

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
					Purpose:      purposeText,
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
