package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Costs é uma estrutura que representa um custo
type Costs struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID do custo
	Cost         string             `json:"cost" bson:"cost" example:"aluguel"`                                    // Descrição do custo
	Amount       Money              `json:"amount" bson:"amount"`                                                  // Valor do custo
	Installments int                `json:"installments" bson:"installments,omitempty" example:"12"`               // Número de parcelas
	Methods      *Methods           `json:"methods" bson:"methods"`                                                // Método de pagamento
	Category     Categories         `json:"category" bson:"category"`                                              // Categoria do custo
	CreatedAt    time.Time          `json:"created_at" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação do custo
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização do custo
} // @name Costs

// CostRequest é uma estrutura que representa uma requisição para criar um custo
type CostRequest struct {
	Cost         string                 `json:"cost" bson:"cost" example:"aluguel" validate:"required"`  // Descrição do custo
	Amount       Money                  `json:"amount" bson:"amount" validate:"required"`                // Valor do custo
	Installments int                    `json:"installments" bson:"installments,omitempty" example:"12"` // Número de parcelas
	Methods      *Methods               `json:"methods" bson:"methods"`                                  // Método de pagamento
	Category     PartialCategoryRequest `json:"category" bson:"category" validate:"required"`            // ID da categoria
} // @name CostRequest
