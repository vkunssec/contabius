package domain

import "time"

type Method string

const (
	MethodCredit Method = "credit"
	MethodDebit  Method = "debit"
	MethodPix    Method = "pix"
	MethodBoleto Method = "boleto"
	MethodCash   Method = "cash"
)

type Methods struct {
	Method    Method    `json:"method" bson:"method"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
