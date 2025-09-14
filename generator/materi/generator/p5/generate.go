package p5

import (
	"context"
	"log"
	"skb_materi/genai/background_urgency"
	"skb_materi/types"
)

type PurposeType int

const (
	SkbType PurposeType = iota
	PppkType
	PppkWawancaraType
)

type Fetch struct {
	bg *background_urgency.Intro
}

func NewFetch(bu *background_urgency.Intro) *Fetch {
	return &Fetch{
		bg: bu,
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
		var competencies = make(map[string][]string)

		for _, chapter := range part.Chapters {
			competencies[chapter.Title] = chapter.BaseCompetitions
		}

		try := 1

		for {

			log.Println("Generate background & urgencies for: ", part.Subject, " with try: ", try)

			out, err := f.bg.Generate(ctx, background_urgency.Prompt{
				Purpose:      purposeText,
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
