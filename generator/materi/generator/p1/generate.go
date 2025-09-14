package p1

import (
	"context"
	"fmt"
	"log"
	"skb_materi/genai/base_material"
	"skb_materi/types"
)

type PurposeType int

const (
	SkbType PurposeType = iota
	PppkType
	PppkWawancaraType
)

type Fetch struct {
	Generator *base_material.BaseMaterial
	Outline   *types.OutlineRoot
}

func NewFetch(gen *base_material.BaseMaterial, outline *types.OutlineRoot) *Fetch {
	return &Fetch{
		Generator: gen,
		Outline:   outline,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *types.Ebook, purposeType PurposeType) error {
	for _, listMateri := range f.Outline.Outline.Output.ListMateri {
		log.Println("[1_fetchMaterial] Fetching material :", listMateri)
		var part = &types.Part{
			Subject: listMateri.MateriPokok,
		}

		err := f.fetchSubMateri(ctx, &listMateri, listMateri.MateriPokok, part, purposeType)
		if err != nil {
			return err
		}

		ebook.Parts = append(ebook.Parts, part)
	}

	return nil
}

func (f *Fetch) fetchSubMateri(ctx context.Context, m *types.MateriPokok, mp string, part *types.Part, purposeType PurposeType) error {
	for _, listSubMateri := range m.ListSubMateri {
		log.Println("[2_fetchSubMaterial] Fetching sub material :", listSubMateri)
		var chapter = &types.Chapter{
			Title: listSubMateri.SubMateriPokok,
		}

		err := f.fetchListMateri(ctx, &listSubMateri, mp, listSubMateri.SubMateriPokok, chapter, purposeType)
		if err != nil {
			return err
		}

		part.Chapters = append(part.Chapters, chapter)
	}

	return nil
}

func (f *Fetch) fetchListMateri(ctx context.Context, s *types.SubMateriPokok, mp, smp string, chapter *types.Chapter, purposeType PurposeType) error {
	var purposeText = ""
	if purposeType == SkbType {
		purposeText = fmt.Sprintf("Seleksi Kompetensi Bidang untuk calon pegawai negeri untuk jabatan %s", f.Outline.NamaJabatan)
	} else if purposeType == PppkType {
		purposeText = fmt.Sprintf("Seleksi Kompetensi Teknis untuk calon pegawai negeri untuk jabatan %s", f.Outline.NamaJabatan)
	} else if purposeType == PppkWawancaraType {
		purposeText = fmt.Sprintf("Seleksi Wawancara untuk calon pegawai negeri untuk jabatan %s", f.Outline.NamaJabatan)
	}

	for _, materi := range s.ListMateri {
		log.Println("[3_fetchListMaterial] Fetching material :", materi)
		var g = &types.Material{
			Title: materi,
		}

		try := 1

		for {
			log.Println("[3_fetchListMaterial] Fetching material try : ", try)
			out, err := f.Generator.Generate(ctx, base_material.MaterialPrompt{
				Purpose:  purposeText,
				Subject:  mp,
				Chapter:  smp,
				Material: materi,
			})

			if err != nil {
				try++
				continue
			}

			g.Short = out.Short

			for _, detail := range out.Details {
				g.Details = append(g.Details, &types.Detail{
					Content: detail,
				})
			}

			log.Println("[3_fetchListMaterial] result : ", g.Details)

			break
		}

		chapter.Materials = append(chapter.Materials, g)
	}

	return nil
}
