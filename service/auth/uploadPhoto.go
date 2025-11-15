package auth

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

type CloudinaryHandler struct {
	Cloud *cloudinary.Cloudinary
}

func NewCloudinaryHandler() (*CloudinaryHandler, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)
	if err != nil {
		return nil, err
	}

	return &CloudinaryHandler{
		Cloud: cld,
	}, nil
}

func (h *CloudinaryHandler) uploadToCloudinary(file multipart.File, header *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	result, err := h.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   "profile-images",
		PublicID: header.Filename,
	})
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

func (h *CloudinaryHandler) UploadImage(c *fiber.Ctx) error {
	formFile, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "image not found",
		})
	}

	file, err := formFile.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot open file",
		})
	}
	defer file.Close()

	url, err := h.uploadToCloudinary(file, formFile)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "upload failed: " + err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"url": url,
	})
}
