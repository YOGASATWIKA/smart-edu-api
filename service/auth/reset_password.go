package auth

import (
	"context"
	"fmt"
	"log"
	"os"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/resend/resend-go/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleForgotPassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if data["email"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email tidak boleh kosong"})
	}

	client := config.GetMongoClient()
	userCollection := client.Database("smart_edu").Collection("users")
	resetCollection := client.Database("smart_edu").Collection("password_resets")

	var user entity.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": data["email"]}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Email tidak terdaftar",
		})
	}

	token, err := helper.GenerateRandomToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat token",
		})
	}

	resetToken := entity.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	_, err = resetCollection.InsertOne(context.TODO(), resetToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
		})
	}

	err = SendResetPasswordEmail(user.Email, token)
	if err != nil {
		log.Println("Gagal mengirim email:", err)
	}

	return c.JSON(fiber.Map{
		"message": "Link reset password telah dikirim ke email Anda" + user.Email,
	})
}
func SendResetPasswordEmail(to string, token string) error {

	fePath := os.Getenv("PATHFE")
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", fePath, token)

	subject := "Reset Password Akun Anda"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6;">

			<h2 style="color: #333;">Reset Password Anda</h2>

			<p>Halo,</p>
			<p>Kami menerima permintaan untuk reset password akun Anda.</p>

			<p>Silakan klik tombol di bawah ini untuk melanjutkan proses reset password:</p>

			<a href="%s" 
				style="
					display: inline-block;
					padding: 12px 20px;
					background-color: #4CAF50;
					color: white;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
				">
				Reset Password
			</a>

			<p style="margin-top: 20px;">Atau jika tombol tidak berfungsi, Anda bisa klik link di bawah ini:</p>

			<p><a href="%s">%s</a></p>

			<p>Link ini berlaku selama <strong>15 menit</strong>.</p>

			<p>Jika Anda tidak meminta reset password, abaikan email ini.</p>

			<br>
			<p>Salam, <br> SmartEdu Support Team</p>

		</body>
		</html>
	`, resetLink, resetLink, resetLink)

	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			body,
	)

	sendmail := os.Getenv("SENDMAILKEY")
	client := resend.NewClient(sendmail)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: subject,
		Html:    string(message),
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("gagal mengirim email reset password: %w", err)
	}

	return nil
}

func HandleResetPassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if data["token"] == "" || data["new_password"] == "" || data["confirm_password"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Semua field harus diisi"})
	}

	if data["new_password"] != data["confirm_password"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Password baru dan konfirmasi tidak sama"})
	}

	client := config.GetMongoClient()
	userCollection := client.Database("smart_edu").Collection("users")
	resetCollection := client.Database("smart_edu").Collection("password_resets")

	var resetToken entity.PasswordReset
	err := resetCollection.FindOne(context.TODO(), bson.M{"token": data["token"]}).Decode(&resetToken)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Token tidak valid atau sudah kadaluarsa"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Token sudah kadaluarsa"})
	}

	hashedPassword, err := HashPassword(data["new_password"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal memproses password"})
	}

	filter := bson.M{"_id": resetToken.UserID}
	update := bson.M{"$set": bson.M{"password": hashedPassword}}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal mengupdate password"})
	}

	resetCollection.DeleteOne(context.TODO(), bson.M{"_id": resetToken.ID})

	return c.JSON(fiber.Map{
		"message": "Password berhasil diubah, silakan login kembali",
	})
}
