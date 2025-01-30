package middleware

import (
	"github.com/vkunssec/contabius/configs"
	"github.com/vkunssec/contabius/pkg/domain"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(allowedProfiles []domain.UserType) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(configs.Env("JWT_SECRET"))},
		ErrorHandler:   jwtErrorHandler,
		SuccessHandler: jwtSuccessHandler(allowedProfiles),
		AuthScheme:     "Bearer",
		ContextKey:     "user",
	})
}

func jwtSuccessHandler(allowedProfiles []domain.UserType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		profileIdUuid := claims["profileId"].(string)
		profileId := domain.UserType(profileIdUuid)
		if !containsProfileId(allowedProfiles, profileId) {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"status": "error", "message": "No permission"})
		}
		return c.Next()
	}
}

func containsProfileId(users []domain.UserType, user domain.UserType) bool {
	for _, profileId := range users {
		if profileId == user {
			return true
		}
	}
	return false
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
}
