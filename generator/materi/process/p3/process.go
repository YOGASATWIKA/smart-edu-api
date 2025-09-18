package p3

import (
	"context"
	"log"
	"smart-edu-api/entity"
	"smart-edu-api/generator/materi/genai/summary"
)

type Fetch struct {
	sum *summary.Summary
}

func NewFetch(sum *summary.Summary) *Fetch {
	return &Fetch{
		sum: sum,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *entity.Ebook) {

	for _, part := range ebook.Parts {
		for _, chapter := range part.Chapters {

			var materials = make(map[string]string)

			for _, material := range chapter.Materials {
				if _, ok := materials[material.Title]; !ok {
					materials[material.Title] = material.Short
				}
			}

			try := 1

			for {

				log.Println("Generate summary & reflection for: ", chapter.Title, " with try: ", try)

				out, err := f.sum.Generate(ctx, summary.Prompt{
					Purpose:      "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan " + ebook.Title,
					BaseMaterial: part.Subject,
					SubMaterial:  chapter.Title,
					Materials:    materials,
				})

				if err != nil {
					try++
					continue
				}

				chapter.Reflections = out.Reflections
				chapter.Conclusion = out.Summary

				break
			}
		}
	}
}
