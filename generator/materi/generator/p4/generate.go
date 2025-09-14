package p4

import (
	"context"
	"log"
	"skb_materi/genai/summary"
	"skb_materi/types"
)

type PurposeType int

const (
	SkbType PurposeType = iota
	PppkType
	PppkWawancaraType
)

type Fetch struct {
	sum *summary.Summary
}

func NewFetch(sum *summary.Summary) *Fetch {
	return &Fetch{
		sum: sum,
	}
}

func (f *Fetch) Fetch(ctx context.Context, ebook *types.Ebook, purposeType PurposeType) {
	var purposeText = "Seleksi Calon Pegawai Negeri"

	if purposeType == SkbType {
		purposeText = "Seleksi Kompetensi Bidang Calon Pegawai Negeri Sipil jabatan " + ebook.Title
	} else if purposeType == PppkType {
		purposeText = "Seleski Kompetensi Teknis Calon Pegawai Negeri Sipil jabatan " + ebook.Title
	} else if purposeType == PppkWawancaraType {
		purposeText = "Selesi Wawancara Calon Pegawai Negeri Sipil jabatan " + ebook.Title
	}

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
					Purpose:      purposeText,
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
