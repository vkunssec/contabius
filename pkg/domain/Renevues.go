package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Revenues são as receitas do usuário
type Revenues struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID da receita
	Revenue   string             `json:"revenue" bson:"revenue" example:"salário"`                                        // Receita
	Amount    Money              `json:"amount" bson:"amount"`                                                            // Valor da receita
	Method    *Methods           `json:"method" bson:"method"`                                                            // Método de pagamento
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação da receita
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização da receita
}
