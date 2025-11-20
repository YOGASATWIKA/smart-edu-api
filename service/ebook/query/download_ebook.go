package query

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"smart-edu-api/repository"
	"strings"

	"baliance.com/gooxml/document"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

func DownloadEbookById(c *fiber.Ctx) error {
	id := c.Params("id")

	existing, err := repository.GetEbookByModulId(id)
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ebook Not Found",
		})
	}

	if strings.TrimSpace(existing.HtmlContent) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ebook Content Is Empty",
		})
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(existing.HtmlContent))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error parsing HTML: %v", err),
		})
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 20)
	pdf.MultiCell(0, 10, existing.Title, "", "C", false)
	pdf.Ln(8)

	// BODY
	renderPDF(doc, pdf)

	// File saving
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	safeName := re.ReplaceAllString(existing.Title, "_")

	saveDir := "./storage/pdf"
	os.MkdirAll(saveDir, os.ModePerm)
	filePath := filepath.Join(saveDir, safeName+".pdf")

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error Writing PDF: %v", err),
		})
	}

	return c.Download(filePath, safeName+".pdf")
}

func renderPDF(doc *goquery.Document, pdf *gofpdf.Fpdf) {
	doc.Find("body").Contents().Each(func(i int, s *goquery.Selection) {

		switch goquery.NodeName(s) {

		case "h1":
			pdf.SetFont("Arial", "B", 18)
			pdf.MultiCell(0, 10, s.Text(), "", "", false)
			pdf.Ln(4)

		case "h2":
			pdf.SetFont("Arial", "B", 16)
			pdf.MultiCell(0, 9, s.Text(), "", "", false)
			pdf.Ln(4)

		case "h3":
			pdf.SetFont("Arial", "B", 14)
			pdf.MultiCell(0, 8, s.Text(), "", "", false)
			pdf.Ln(4)

		case "p":
			pdf.SetFont("Arial", "", 12)
			pdf.MultiCell(0, 7, s.Text(), "", "", false)
			pdf.Ln(3)

		case "ul":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				pdf.SetFont("Arial", "", 12)
				pdf.MultiCell(0, 7, "• "+li.Text(), "", "", false)
			})
			pdf.Ln(2)

		case "strong", "b":
			pdf.SetFont("Arial", "B", 12)
			pdf.MultiCell(0, 7, s.Text(), "", "", false)
			pdf.Ln(3)

		case "i", "em":
			pdf.SetFont("Arial", "I", 12)
			pdf.MultiCell(0, 7, s.Text(), "", "", false)
			pdf.Ln(3)
		}
	})
}

func DownloadEbookWordById(c *fiber.Ctx) error {
	id := c.Params("id")
	existing, err := repository.GetEbookByModulId(id)
	if err != nil || existing == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ebook Not Found",
		})
	}

	if strings.TrimSpace(existing.HtmlContent) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ebook Content Is Empty",
		})
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	safeName := re.ReplaceAllString(existing.Title, "_")

	saveDir := "./storage/word"
	os.MkdirAll(saveDir, os.ModePerm)
	filePath := filepath.Join(saveDir, safeName+".docx")

	doc := document.New()

	// Title
	title := doc.AddParagraph()
	run := title.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetSize(32)
	run.AddText(existing.Title)
	doc.AddParagraph()

	// BODY
	renderWord(doc, existing.HtmlContent)

	if err := doc.SaveToFile(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error Saving File Word: %v", err),
		})
	}

	return c.Download(filePath, safeName+".docx")
}

func renderWord(doc *document.Document, html string) {
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	dom.Find("body").Children().Each(func(i int, s *goquery.Selection) {
		tag := goquery.NodeName(s)
		text := strings.TrimSpace(s.Text())
		if text == "" {
			return
		}

		switch tag {

		case "h1":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			run.Properties().SetSize(28)
			run.AddText(text)

		case "h2":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			run.Properties().SetSize(24)
			run.AddText(text)

		case "h3":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			run.Properties().SetSize(20)
			run.AddText(text)

		case "p":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetSize(12)
			run.AddText(text)

		case "ul":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				para := doc.AddParagraph()
				run := para.AddRun()
				run.Properties().SetSize(12)
				run.AddText("• " + li.Text())
			})

		case "strong", "b":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetBold(true)
			run.Properties().SetSize(12)
			run.AddText(text)

		case "i", "em":
			para := doc.AddParagraph()
			run := para.AddRun()
			run.Properties().SetItalic(true)
			run.Properties().SetSize(12)
			run.AddText(text)
		}
	})
}
