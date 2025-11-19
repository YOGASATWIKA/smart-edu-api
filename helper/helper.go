package helper

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fiber.NewError(fiber.StatusBadRequest, "Password minimal 8 karakter")
	}

	var (
		hasUpper = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasDigit = regexp.MustCompile(`[0-9]`).MatchString(password)
	)

	if !hasUpper || !hasLower || !hasDigit {
		return fiber.NewError(fiber.StatusBadRequest,
			"Password harus mengandung huruf besar, huruf kecil, dan angka")
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
