package query

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"smart-edu-api/helper"
	"smart-edu-api/repository"
	"strings"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	"github.com/PuerkitoBio/goquery"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gofiber/fiber/v2"
)

func DownloadEbookById(c *fiber.Ctx) error {
	// Ambil ID modul dari parameter URL
	id := c.Params("id")
	ctx := helper.GetContext()

	// Ambil data eBook berdasarkan modul ID
	existing, err := repository.GetEbookByModulId(ctx, id)
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Outline tidak ditemukan",
		})
	}

	// Validasi jika HTML content kosong
	if strings.TrimSpace(existing.HtmlContent) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Konten eBook kosong",
		})
	}

	// Inisialisasi PDF generator (wkhtmltopdf)
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal inisialisasi PDF generator: %v", err),
		})
	}
	// Buat halaman dari HTML content
	page := wkhtmltopdf.NewPageReader(strings.NewReader(existing.HtmlContent))
	page.EnableLocalFileAccess.Set(true)
	page.FooterRight.Set("[page]")
	page.Zoom.Set(1.0)
	pdfg.AddPage(page)
	pdfg.MarginLeft.Set(20)
	pdfg.MarginRight.Set(20)
	pdfg.MarginTop.Set(25)
	pdfg.MarginBottom.Set(15)
	pdfg.Dpi.Set(300)

	// Generate PDF
	if err := pdfg.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal membuat PDF: %v", err),
		})
	}

	// Bersihkan nama file agar aman
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	safeName := re.ReplaceAllString(existing.Title, "_")

	// Pastikan direktori penyimpanan ada
	saveDir := "./storage/pdf"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal membuat folder penyimpanan: %v", err),
		})
	}

	// Tentukan path file PDF yang akan disimpan
	filePath := filepath.Join(saveDir, fmt.Sprintf("%s.pdf", safeName))

	// Simpan PDF ke file
	if err := pdfg.WriteFile(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menulis file PDF: %v", err),
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.pdf\"", safeName))
	c.Set("X-Filename", fmt.Sprintf("%s.pdf", safeName))
	return c.Download(filePath, safeName+".pdf")
}

func DownloadEbookWordById(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := helper.GetContext()

	existing, err := repository.GetEbookByModulId(ctx, id)
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Outline tidak ditemukan",
		})
	}

	if strings.TrimSpace(existing.HtmlContent) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Konten eBook kosong",
		})
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	safeName := re.ReplaceAllString(existing.Title, "_")

	saveDir := "./storage/word"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal membuat folder penyimpanan: %v", err),
		})
	}

	filePath := filepath.Join(saveDir, fmt.Sprintf("%s.docx", safeName))

	// 🟢 Buat dokumen Word baru
	doc := document.New()

	// Tambahkan judul besar
	titlePara := doc.AddParagraph()
	titleRun := titlePara.AddRun()
	titleRun.Properties().SetBold(true)
	titleRun.Properties().SetSize(32)
	titleRun.AddText(existing.Title)
	doc.AddParagraph().AddRun().AddBreak()

	addFormattedHTML(doc, existing.HtmlContent)

	// Simpan file Word
	if err := doc.SaveToFile(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menyimpan file Word: %v", err),
		})
	}

	return c.Download(filePath, safeName+".docx")
}

func addFormattedHTML(doc *document.Document, html string) {
	reader := strings.NewReader(html)
	dom, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	dom.Find("h1, h2, h3, p, b, strong, i, em, u, ul, ol, li, br").Each(func(i int, s *goquery.Selection) {
		tag := goquery.NodeName(s)
		text := strings.TrimSpace(s.Text())
		if text == "" {
			return
		}

		switch tag {
		case "h1", "h2", "h3":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			switch tag {
			case "h1":
				run.Properties().SetSize(28)
			case "h2":
				run.Properties().SetSize(24)
			case "h3":
				run.Properties().SetSize(20)
			}
			run.AddText(text)

		case "b", "strong":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			run.AddText(text)

		case "i", "em":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetItalic(true)
			run.AddText(text)

		case "u":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Auto)

			run.AddText(text)

		case "ul", "ol", "li":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.AddText("• " + text)

		case "p":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.AddText(text)

		case "br":
			doc.AddParagraph().AddRun().AddBreak()
		}
	})
}
