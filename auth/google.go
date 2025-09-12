package auth

import (
	"context"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fungsi utama untuk mendaftarkan semua rute terkait otentikasi
func RegisterGoogleRoutes(app *fiber.App) {
	// Rute publik untuk login/register
	app.Post("/api/auth/google", HandleGoogleLogin)

	// Grup rute yang memerlukan otentikasi
	profileGroup := app.Group("/api/profile", AuthenticateToken)

	// Rute yang sudah ada untuk mendapatkan semua data profil
	profileGroup.Get("/", HandleGetProfile) // Sebelumnya app.Get("/api/profile", ...)

	// --- RUTE BARU YANG DITAMBAHKAN ---
	// Rute untuk mendapatkan nama user
	profileGroup.Get("/name", HandleGetUserName)
	// Rute untuk mendapatkan URL foto profil
	profileGroup.Get("/picture", HandleGetUserPicture)
}

// HandleGoogleLogin menangani callback dari frontend setelah login Google
func HandleGoogleLogin(c *fiber.Ctx) error {
	// 1. Parse request body dari frontend
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
	// 2. Cari user di database berdasarkan GoogleID

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"googleId": req.GoogleID}).Decode(&user)

	if err != nil {
		// Jika user tidak ditemukan, buat user baru
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
		} else { // Error database lainnya
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}
	}

	// 3. Buat JWT Token
	claims := jwt.MapClaims{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token berlaku 3 hari
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	// 4. Kirim token sebagai response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}

// auth/google_auth.go (lanjutan...)

// AuthenticateToken adalah middleware untuk memvalidasi JWT
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

	// Simpan claims ke context Fiber agar bisa diakses oleh handler selanjutnya
	c.Locals("userClaims", token.Claims.(jwt.MapClaims))
	return c.Next()
}

// HandleGetProfile adalah contoh handler untuk rute yang dilindungi
func HandleGetProfile(c *fiber.Ctx) error {
	// Ambil data user dari context yang sudah divalidasi oleh middleware
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

// --- HANDLER BARU UNTUK NAMA USER ---
func HandleGetUserName(c *fiber.Ctx) error {
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userIDHex := claims["id"].(string)
	userID, _ := primitive.ObjectIDFromHex(userIDHex)

	var userProfile entity.User
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&userProfile)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Hanya kembalikan nama dalam format JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"name": userProfile.Name})
}

// --- HANDLER BARU UNTUK FOTO PROFIL ---
func HandleGetUserPicture(c *fiber.Ctx) error {
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userIDHex := claims["id"].(string)
	userID, _ := primitive.ObjectIDFromHex(userIDHex)

	var userProfile entity.User
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&userProfile)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Hanya kembalikan URL foto profil dalam format JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"picture": userProfile.Picture})
}
