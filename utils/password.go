package utils

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 8
)

// HashPassword hash da senha usando bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compara uma senha com uma senha hashed usando bcrypt
func ComparePassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// CheckPasswordStrength verifica a força de uma senha
func CheckPasswordStrength(password string) (bool, error) {
	if len(password) < MinPasswordLength {
		return false, errors.New("senha deve ter pelo menos 8 caracteres")
	}

	hasUppercase, err := regexp.MatchString(`[A-Z]`, password)
	if err != nil {
		return false, err
	}
	if !hasUppercase {
		return false, errors.New("senha deve ter pelo menos uma letra maiúscula")
	}

	hasLowercase, err := regexp.MatchString(`[a-z]`, password)
	if err != nil {
		return false, err
	}
	if !hasLowercase {
		return false, errors.New("senha deve ter pelo menos uma letra minúscula")
	}

	hasDigit, err := regexp.MatchString(`[0-9]`, password)
	if err != nil {
		return false, err
	}
	if !hasDigit {
		return false, errors.New("senha deve ter pelo menos um número")
	}

	hasSpecialChar, err := regexp.MatchString(`[!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]`, password)
	if err != nil {
		return false, err
	}

	return hasSpecialChar, nil
}
