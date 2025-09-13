package auth

import (
	"context"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthenticateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header is required"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid or expired token"})
	}
	c.Locals("userClaims", token.Claims.(jwt.MapClaims))
	return c.Next()
}

func HandleGetProfile(c *fiber.Ctx) error {
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userIDHex := claims["id"].(string)

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	var userProfile entity.User
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&userProfile)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(userProfile)
}
