package auth

import (
	"context"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleGoogleLogin(c *fiber.Ctx) error {
	var req struct {
		GoogleID string `json:"googleId"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Picture  string `json:"picture"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user entity.User

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"googleId": req.GoogleID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			newUser := entity.User{
				ID:       primitive.NewObjectID(),
				GoogleID: req.GoogleID,
				Email:    req.Email,
				Name:     req.Name,
				Picture:  req.Picture,
			}
			_, insertErr := collection.InsertOne(context.TODO(), newUser)
			if insertErr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
			}
			user = newUser
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}
	}

	claims := jwt.MapClaims{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
	})
}
