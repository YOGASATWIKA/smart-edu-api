package auth

import (
	"context"
	"os"
	"regexp"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	ID    primitive.ObjectID `json:"_id,omitempty"`
	Email string             `json:"email,omitempty"`
	Name  string             `json:"name,omitempty"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func HandleRegister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	requiredFields := map[string]string{
		"name":     "Name tidak boleh kosong",
		"email":    "Email tidak boleh kosong",
		"password": "Password tidak boleh kosong",
	}

	for field, message := range requiredFields {
		if val, ok := data[field]; !ok || val == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": message,
			})
		}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(data["email"]) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format email tidak valid",
		})
	}

	if err := helper.ValidatePassword(data["password"]); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := HashPassword(data["password"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not hash password",
		})
	}

	user := entity.User{
		ID:       primitive.NewObjectID(),
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")

	err = collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&entity.User{})
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email Sudah Terdaftar",
		})
	}
	if err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	_, insertErr := collection.InsertOne(context.TODO(), user)
	if insertErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to register user"})
	}
	claims := jwt.MapClaims{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	res := response{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": tokenString,
		"user":  res,
	})
}
