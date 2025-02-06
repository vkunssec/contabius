package domain

import (
	"errors"
)

// Method é uma enumeração que representa o método de pagamento
type Method string // @name Method

// MethodId é uma enumeração que representa o ID do método de pagamento
type MethodId int // @name MethodId

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
} // @name Methods

// Validate valida o método de pagamento
func (m *Methods) Validate() error {
	if m.Method == "" {
		return errors.New("método não informado")
	}
	return nil
}

var (
	// methodMap armazena o mapeamento de IDs para Methods
	methodMap = map[MethodId]Method{
		MethodCreditId: MethodCredit,
		MethodDebitId:  MethodDebit,
		MethodPixId:    MethodPix,
		MethodBoletoId: MethodBoleto,
		MethodCashId:   MethodCash,
	}

	// allMethods armazena a lista completa de métodos de pagamento
	allMethods = func() []Methods {
		methods := make([]Methods, 0, len(methodMap))
		for id, method := range methodMap {
			methods = append(methods, Methods{MethodId: id, Method: method})
		}
		return methods
	}()
)

// GetMethod retorna um método de pagamento
func GetMethod(id MethodId) (Methods, error) {
	method, exists := methodMap[id]
	if !exists {
		return Methods{}, errors.New("método de pagamento inválido")
	}

	methodObj := Methods{MethodId: id, Method: method}
	return methodObj, methodObj.Validate()
}

// GetMethods retorna todos os métodos de pagamento solicitados
func GetMethods(ids []MethodId) ([]Methods, error) {
	methods := make([]Methods, 0, len(ids))
	for _, id := range ids {
		method, err := GetMethod(id)
		if err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}
	return methods, nil
}

// AllMethods retorna todos os métodos de pagamento
func AllMethods() []Methods {
	return allMethods
}
