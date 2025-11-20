package auth

import (
	"context"
	"smart-edu-api/config"
	request "smart-edu-api/data/user/request"
	"smart-edu-api/entity"
	"smart-edu-api/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleUpdateProfile(c *fiber.Ctx) error {
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userIDHex := claims["id"].(string)

	existing, err := repository.GetUserById(userIDHex)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	request := new(request.UpdateUserRequest)
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

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")

	err = collection.FindOne(context.TODO(), bson.M{"email": request.Email}).Decode(&entity.User{})
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email Sudah Terdaftar",
		})
	}
	if err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	existing.Name = request.Name
	existing.Email = request.Email
	existing.Picture = request.Picture

	updateProfile, err := repository.UpdateUser(existing)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Model Successfully",
		"data":    updateProfile,
	})
}
