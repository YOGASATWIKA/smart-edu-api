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
	html.WriteString(fmt.Sprintf(`<h1 style="text-align:center;">%s</h1>`, safeString(ebook.Title)))

	if len(ebook.Parts) == 0 {
		html.WriteString("<p>Konten tidak tersedia.</p>")
		return html.String()
	}

	for i, part := range ebook.Parts {
		if part == nil {
			continue
		}

		html.WriteString(fmt.Sprintf(`<h2 style="margin-top:30px;">BAB %02d - %s</h2>`, i+1, safeString(part.Subject)))

		if len(part.Introductions) > 0 {
			html.WriteString(`<h4>Latar Belakang</h4>`)
			for _, intro := range part.Introductions {
				html.WriteString(fmt.Sprintf(`<p>%s</p>`, intro))
			}
		}

		if len(part.Urgencies) > 0 {
			html.WriteString(`<h4>Urgensi Mempelajari Materi Ini</h4>`)
			for _, urg := range part.Urgencies {
				html.WriteString(fmt.Sprintf(`<p>%s</p>`, urg))
			}
		}

		for j, chapter := range part.Chapters {
			if chapter == nil {
				continue
			}

			subLetter := string(rune('A' + j))
			html.WriteString(fmt.Sprintf(`<h3>%s. %s</h3>`, subLetter, safeString(chapter.Title)))

			if len(chapter.BaseCompetitions) > 0 {
				html.WriteString(`<h4>Kompetensi Dasar</h4><ol>`)
				for _, base := range chapter.BaseCompetitions {
					html.WriteString(fmt.Sprintf(`<li>%s</li>`, base))
				}
				html.WriteString(`</ol>`)
			}

			if len(chapter.TriggerQuestions) > 0 {
				html.WriteString(`<h4>Pertanyaan Pematik</h4><ol>`)
				for _, tq := range chapter.TriggerQuestions {
					html.WriteString(fmt.Sprintf(`<li>%s</li>`, tq))
				}
				html.WriteString(`</ol>`)
			}

			for _, material := range chapter.Materials {
				if material == nil {
					continue
				}
				html.WriteString(fmt.Sprintf(`<h4>%s</h4>`, safeString(material.Title)))

				if material.Short != "" {
					html.WriteString(fmt.Sprintf(`<p><em>%s</em></p>`, material.Short))
				}

				for _, detail := range material.Details {
					if detail == nil {
						continue
					}
					if detail.Content != "" {
						html.WriteString(fmt.Sprintf(`<p>%s</p>`, detail.Content))
					}
					if detail.Expanded != "" {
						html.WriteString(fmt.Sprintf(`<p>%s</p>`, detail.Expanded))
					}
					for _, chunk := range detail.ExpandChunks {
						html.WriteString(fmt.Sprintf(`<p>%s</p>`, chunk))
					}
				}
			}

			if chapter.Conclusion != "" {
				html.WriteString(`<h4>Kesimpulan</h4>`)
				html.WriteString(fmt.Sprintf(`<p>%s</p>`, chapter.Conclusion))
			}

			if len(chapter.Reflections) > 0 {
				html.WriteString(`<h4>Refleksi Pembelajaran</h4><ol>`)
				for _, reflection := range chapter.Reflections {
					html.WriteString(fmt.Sprintf(`<li>%s</li>`, reflection))
				}
				html.WriteString(`</ol>`)
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
