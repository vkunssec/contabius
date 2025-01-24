package domain

import "time"

// Categories é uma estrutura que representa uma categoria
type Categories struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID da categoria
	Category  string    `json:"category" bson:"category" example:"Alimentação"`                                  // Nome da categoria
	Parent    *string   `json:"parent" bson:"parent" example:"678079f6f5080a39a8eedc1e"`                         // ID do pai da categoria
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização
}
