package helper

import (
	"fmt"
	"smart-edu-api/entity"
	"strings"
)

func TransformEbookToHTML(ebook *entity.Ebook) string {
	if ebook == nil {
		return "<p>Data ebook tidak ditemukan.</p>"
	}

	var html strings.Builder
	html.WriteString(fmt.Sprintf("<h1>%s</h1>", safeString(ebook.Title)))

	if len(ebook.Parts) == 0 {
		html.WriteString("<p>Konten tidak tersedia</p>")
		return html.String()
	}

	for i, part := range ebook.Parts {
		if part == nil {
			continue
		}
		html.WriteString(fmt.Sprintf("<h2>Bab %d: %s</h2>", i+1, safeString(part.Subject)))

		// Introductions
		for _, intro := range part.Introductions {
			html.WriteString(fmt.Sprintf("<p><strong>Pendahuluan:</strong> %s</p>", intro))
		}

		// Urgencies
		for _, urg := range part.Urgencies {
			html.WriteString(fmt.Sprintf("<p><em>Urgensi:</em> %s</p>", urg))
		}

		for j, chapter := range part.Chapters {
			if chapter == nil {
				continue
			}
			html.WriteString(fmt.Sprintf("<h3>%d.%d %s</h3>", i+1, j+1, safeString(chapter.Title)))

			// Base Competencies
			for _, base := range chapter.BaseCompetitions {
				html.WriteString(fmt.Sprintf("<p><strong>Kompetensi Dasar:</strong> %s</p>", base))
			}

			// Trigger Questions
			for _, tq := range chapter.TriggerQuestions {
				html.WriteString(fmt.Sprintf("<p><em>Pertanyaan Pemicu:</em> %s</p>", tq))
			}

			for _, material := range chapter.Materials {
				if material == nil {
					continue
				}
				html.WriteString(fmt.Sprintf("<h4>%s</h4>", safeString(material.Title)))

				if material.Short != "" {
					html.WriteString(fmt.Sprintf("<p><em>\"%s\"</em></p>", material.Short))
				}

				for _, detail := range material.Details {
					if detail == nil {
						continue
					}
					if detail.Content != "" {
						html.WriteString(fmt.Sprintf("<p>%s</p>", detail.Content))
					}
					if detail.Expanded != "" {
						html.WriteString(fmt.Sprintf("<p>%s</p>", detail.Expanded))
					}
					for _, chunk := range detail.ExpandChunks {
						html.WriteString(fmt.Sprintf("<p>%s</p>", chunk))
					}
				}
			}

			if chapter.Conclusion != "" {
				html.WriteString(fmt.Sprintf("<p><strong>Kesimpulan:</strong> %s</p>", chapter.Conclusion))
			}

			for _, reflection := range chapter.Reflections {
				html.WriteString(fmt.Sprintf("<p><em>Refleksi:</em> %s</p>", reflection))
			}
		}
	}
	return html.String()
}

func safeString(s string) string {
	if s == "" {
		return "Tanpa Judul"
	}
	return s
}
