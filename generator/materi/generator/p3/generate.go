package p3

import (
	"context"
	"log"
	"skb_materi/genai/expand_material"
	"skb_materi/types"
)

type PurposeType int

const (
	SkbType PurposeType = iota
	PppkType
	PppkWawancaraType
)

type Fetch struct {
	ep *expand_material.ExpandMaterial
}

func NewFetch(ep *expand_material.ExpandMaterial) *Fetch {
	return &Fetch{ep: ep}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *types.Ebook, purpose PurposeType) {
	var purposeText = ""

	if purpose == SkbType {
		purposeText = "Materi bahan ajar seleksi kompetensi bidang (SKB) calon aparatur sipil negara (seleksi CPNS) untuk formasi jabatan " + ebook.Title
	} else if purpose == PppkType {
		purposeText = "Materi bahan ajar kompetensi Teknis calon aparatur sipil negara (seleksi PPPK) untuk formasi jabatan " + ebook.Title
	} else if purpose == PppkWawancaraType {
		purposeText = "Materi bahan ajar seleksi wawancara calon aparatur sipil negara (seleksi PPPK) untuk formasi jabatan " + ebook.Title
	}

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {

			var learnings []string

			for _, material := range chapter.Materials {
				learnings = append(learnings, material.Title)
			}

			for _, material := range chapter.Materials {
				for _, detail := range material.Details {

					try := 1

					for {
						log.Println("Expanding words \"", detail.Content, "\" with try:", try)

						out, err := f.ep.Generate(ctx, expand_material.ExpandedMaterialPrompt{
							Purpose:        purposeText,
							BaseMaterial:   part.Subject,
							SubMaterial:    chapter.Title,
							Learnings:      learnings,
							MainTopics:     chapter.BaseCompetitions,
							ChosenLearning: material.Title,
							WordsToExpand:  detail.Content,
						})

						if err != nil {
							try++
							continue
						}

						detail.Expanded = out.Expanded

						break
					}
				}
			}
		}
	}
}
