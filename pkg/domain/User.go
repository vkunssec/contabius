package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserType é um tipo de usuário
type UserType string

const (
	UserTypeAdmin UserType = "admin"
	UserTypeUser  UserType = "user"
)

// User é uma estrutura que representa um usuário
type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID do usuário
	Name      string             `json:"name" bson:"name" example:"John Doe"`                                             // Nome do usuário
	Email     string             `json:"email" bson:"email" example:"user@example.com"`                                   // Email do usuário
	Password  string             `json:"password" bson:"password" example:"password123"`                                  // Senha do usuário
	Type      UserType           `json:"type" bson:"type" example:"admin"`                                                // Tipo de usuário
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2021-01-01T00:00:00Z"` // Data de criação do usuário
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2021-01-01T00:00:00Z"` // Data de atualização do usuário
}

// UserRequest é uma estrutura que representa a requisição de criação de um usuário
type UserRequest struct {
	Name     string   `json:"name" bson:"name" example:"John Doe" validate:"required"`                 // Nome do usuário
	Email    string   `json:"email" bson:"email" example:"user@example.com" validate:"required,email"` // Email do usuário
	Password string   `json:"password" bson:"password" example:"password123" validate:"required"`      // Senha do usuário
	Type     UserType `json:"type" bson:"type" example:"admin" validate:"required"`                    // Tipo de usuário
}

// UserUpdateRequest é uma estrutura que representa a requisição de atualização de um usuário
type UserUpdateRequest struct {
	Name  string    `json:"name" bson:"name" example:"John Doe" validate:"required"`                 // Nome do usuário
	Email string    `json:"email" bson:"email" example:"user@example.com" validate:"required,email"` // Email do usuário
	Type  *UserType `json:"type,omitempty" bson:"type,omitempty" example:"admin"`                    // Tipo de usuário
}

// UserUpdatePassword é uma estrutura que representa a requisição de atualização de senha de um usuário
type UserUpdatePassword struct {
	Password        string `json:"password" bson:"password" example:"password123" validate:"required"`                 // Senha do usuário
	PasswordConfirm string `json:"password_confirm" bson:"password_confirm" example:"password123" validate:"required"` // Senha do usuário
}

// UserLogin é uma estrutura que representa a requisição de login de um usuário
type UserLogin struct {
	Email    string `json:"email" bson:"email" example:"user@example.com" validate:"required,email"` // Email do usuário
	Password string `json:"password" bson:"password" example:"password123" validate:"required"`      // Senha do usuário
}
