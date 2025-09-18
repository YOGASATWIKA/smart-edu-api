package p1

import (
	"context"
	"fmt"
	"log"
	"smart-edu-api/embeded"
	"smart-edu-api/entity"
	"smart-edu-api/generator/materi/genai/base_material"
)

type Fetch struct {
	Generator *base_material.BaseMaterial
	ctx       context.Context
	ebook     *entity.Ebook
	Outline   *entity.Outline
}

func NewFetch(gen *base_material.BaseMaterial, outline *entity.Outline) *Fetch {
	return &Fetch{
		Generator: gen,
		Outline:   outline,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *entity.Ebook) error {
	for _, listMateri := range f.Outline.Outline.ListMateri {
		log.Println("[1_fetchMaterial] Fetching material :", listMateri)
		var part = &entity.Part{
			Subject: listMateri.MateriPokok,
		}

		err := f.fetchSubMateri(ctx, &listMateri, listMateri.MateriPokok, part)
		if err != nil {
			return err
		}

		ebook.Parts = append(ebook.Parts, part)
	}

	return nil
}

func (f *Fetch) fetchSubMateri(ctx context.Context, m *embeded.MateriPokok, mp string, part *entity.Part) error {
	for _, listSubMateri := range m.ListSubMateri {
		log.Println("[2_fetchSubMaterial] Fetching sub material :", listSubMateri)
		var chapter = &entity.Chapter{
			Title: listSubMateri.SubMateriPokok,
		}

		err := f.fetchListMateri(ctx, &listSubMateri, mp, listSubMateri.SubMateriPokok, chapter)
		if err != nil {
			return err
		}

		part.Chapters = append(part.Chapters, chapter)
	}

	return nil
}

func (f *Fetch) fetchListMateri(ctx context.Context, s *embeded.SubMateriPokok, mp, smp string, chapter *entity.Chapter) error {
	for _, materi := range s.ListMateri {
		log.Println("[3_fetchListMaterial] Fetching material :", materi)
		var g = &entity.Material{
			Title: materi,
		}

		try := 1

		for {
			log.Println("[3_fetchListMaterial] Fetching material try : ", try)
			out, err := f.Generator.Generate(ctx, base_material.MaterialPrompt{
				Purpose:  fmt.Sprintf("Seleksi Kompetensi Bidang untuk calon pegawai negeri untuk jabatan %s", f.Outline.MateriPokok.Namajabatan),
				Subject:  mp,
				Chapter:  smp,
				Material: materi,
			})

			if err != nil {
				log.Println("", err)
				try++
				continue
			}

			g.Short = out.Short

			for _, detail := range out.Details {
				g.Details = append(g.Details, &entity.Detail{
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
