package service

import (
	request "smart-edu-api/data/baseMateri/request"
	"smart-edu-api/helper"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func UpdateMateriPokok(c *fiber.Ctx) error {
	id := c.Params("id") // ambil ID dari path parameter

	// Cek apakah data dengan ID tersebut ada
	existing, err := repository.GetMateriPokokByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Base Materi tidak ditemukan",
		})
	}

	// Parse dan validasi body request
	request := new(request.UpdateMateriPokokRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Invalid request body",
		})
	}

	isValid, err := govalidator.ValidateStruct(request)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	existing.Namajabatan = request.Namajabatan
	existing.Tugasjabatan = request.Tugasjabatan
	existing.Keterampilan = request.Keterampilan
	existing.Klasifikasi = request.Klasifikasi
	existing.UpdatedAt = helper.GetCurrentTime()

	// Simpan perubahan
	updated, err := repository.UpdateMateriPokok(existing)
	if err != nil {
		logrus.Printf("Terjadi error saat update: %s\n", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Gagal mengupdate data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Base materi berhasil diperbarui",
		"data":    updated,
	})
}
