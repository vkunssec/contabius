package domain

import (
	"errors"
)

// Method é uma enumeração que representa o método de pagamento
type Method string

type MethodId int

const (
	MethodCreditId MethodId = iota
	MethodDebitId
	MethodPixId
	MethodBoletoId
	MethodCashId

	MethodCredit Method = "credit"
	MethodDebit  Method = "debit"
	MethodPix    Method = "pix"
	MethodBoleto Method = "boleto"
	MethodCash   Method = "cash"
)

// Methods é uma estrutura que representa um método de pagamento
type Methods struct {
	MethodId MethodId `json:"id" example:"1"`          // ID do método de pagamento
	Method   Method   `json:"method" example:"credit"` // Método de pagamento
}

// Validate valida o método de pagamento
func (m *Methods) Validate() error {
	if m.Method == "" {
		return errors.New("método não informado")
	}
	return nil
}

// AllMethods retorna todos os métodos de pagamento
func AllMethods() []Methods {
	return []Methods{
		{Method: MethodCredit, MethodId: MethodCreditId},
		{Method: MethodDebit, MethodId: MethodDebitId},
		{Method: MethodPix, MethodId: MethodPixId},
		{Method: MethodBoleto, MethodId: MethodBoletoId},
		{Method: MethodCash, MethodId: MethodCashId},
	}
}

// GetMethod retorna um método de pagamento
func GetMethod(id MethodId) (Methods, error) {
	var method Method
	switch id {
	case MethodCreditId:
		method = MethodCredit
	case MethodDebitId:
		method = MethodDebit
	case MethodPixId:
		method = MethodPix
	case MethodBoletoId:
		method = MethodBoleto
	case MethodCashId:
		method = MethodCash
	default:
		return Methods{}, errors.New("método de pagamento inválido")
	}

	methodObj := Methods{
		MethodId: id,
		Method:   method,
	}

	if err := methodObj.Validate(); err != nil {
		return Methods{}, err
	}

	return methodObj, nil
}

// GetMethods retorna todos os métodos de pagamento
func GetMethods(ids []MethodId) ([]Methods, error) {
	methods := []Methods{}
	for _, id := range ids {
		method, err := GetMethod(id)
		if err != nil {
			return []Methods{}, err
		}
		methods = append(methods, method)
	}
	return methods, nil
}
