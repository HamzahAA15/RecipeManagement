package authentication

import (
	"encoding/base64"
	"recipeApps/internals/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware(userRepo repository.IUserRepository) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var username, password string
		authHeader := c.Get("Authorization")

		if strings.HasPrefix(authHeader, "Basic ") {
			credentials, err := base64.StdEncoding.DecodeString(authHeader[6:])
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
			}

			creds := strings.Split(string(credentials), ":")
			username = creds[0]
			password = creds[1]
		}

		user, err := userRepo.GetUser(username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid username with error: " + err.Error()})
		}

		c.Locals("username", user.Name)
		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password with error: " + err.Error()})
		}

		return c.Next()
	}
}

func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
