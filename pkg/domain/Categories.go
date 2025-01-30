package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Categories é uma estrutura que representa uma categoria
type Categories struct {
	Id        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID da categoria
	Category  string              `json:"category" bson:"category" example:"Alimentação"`                                  // Nome da categoria
	Parent    *primitive.ObjectID `json:"parent" bson:"parent,omitempty" example:"678079f6f5080a39a8eedc1e"`               // ID do pai da categoria
	CreatedAt time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação
	UpdatedAt time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização
}

// CategoryRequest é uma estrutura que representa uma requisição para criar uma categoria
type CategoryRequest struct {
	Category string              `json:"category" bson:"category" example:"Alimentação" validate:"required"` // Nome da categoria
	Parent   *primitive.ObjectID `json:"parent" bson:"parent,omitempty" example:"678079f6f5080a39a8eedc1e"`  // ID do pai da categoria
}

type PartialCategoryRequest struct {
	Id       primitive.ObjectID `json:"id" bson:"_id" example:"678079f6f5080a39a8eedc1e" validate:"required"` // ID da categoria
	Category string             `json:"category" bson:"category" example:"Alimentação" validate:"required"`   // Nome da categoria
}
