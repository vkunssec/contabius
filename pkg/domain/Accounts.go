package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Accounts é uma estrutura que representa uma conta
type Accounts struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID da conta
	Account   string             `json:"account,omitempty" bson:"account,omitempty" example:"Conta Corrente"`             // Nome da conta
	Color     string             `json:"color,omitempty" bson:"color,omitempty" example:"#000000"`                        // Cor da conta
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização
}
