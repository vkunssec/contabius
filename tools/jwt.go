package tools

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vkunssec/contabius/configs"
	"github.com/vkunssec/contabius/pkg/domain"
)

// GenerateJWT gera um token JWT para um usuário
func GenerateJWT(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
		"type":  user.Type,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   configs.ServerName,
		"nbf":   time.Now().Unix(),
		"iat":   time.Now().Unix(),
		"sub":   "user",
		"aud":   "users",
		"jti":   uuid.New(),
	})

	tokenString, err := token.SignedString([]byte(configs.Env("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT valida um token JWT
func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Env("JWT_SECRET")), nil
	})
}

// GetUserId retorna o ID do usuário do token JWT
func GetUserId(ctx *fiber.Ctx) string {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	return userId
}

// GetType retorna o tipo do usuário do token JWT
func GetType(ctx *fiber.Ctx) string {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userType := claims["type"].(string)
	return userType
}

// GetEmail retorna o email do usuário do token JWT
func GetEmail(ctx *fiber.Ctx) string {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return email
}
