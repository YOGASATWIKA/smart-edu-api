package query

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"smart-edu-api/helper"
	"smart-edu-api/repository"
	"strings"

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
	page.HeaderCenter.Set(existing.Title)
	page.FooterRight.Set("[page]")
	page.Zoom.Set(1.0)
	pdfg.AddPage(page)

	// Set margin dan kualitas
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
