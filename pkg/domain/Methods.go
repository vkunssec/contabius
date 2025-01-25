package domain

import (
	"errors"
	"time"
)

// Method é uma enumeração que representa o método de pagamento
type Method string

const (
	MethodCredit Method = "credit"
	MethodDebit  Method = "debit"
	MethodPix    Method = "pix"
	MethodBoleto Method = "boleto"
	MethodCash   Method = "cash"
)

// Methods é uma estrutura que representa um método de pagamento
type Methods struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty" example:"678079f6f5080a39a8eedc1e"`            // ID do método de pagamento
	Method    Method    `json:"method" bson:"method" example:"credit"`                                           // Método de pagamento
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de criação
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" example:"2025-01-01T00:00:00Z"` // Data de atualização
}

func (m *Methods) Validate() error {
	if m.Method == "" {
		return errors.New("método não informado")
	}
	return nil
}
