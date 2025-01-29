package presentation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// IError representa a estrutura de um erro de validação
type IError struct {
	Field string `json:"field,omitempty"` // Nome do campo com erro
	Error string `json:"error,omitempty"` // Mensagem de erro
}

// Validator é uma instância global do validador
var Validator = validator.New()

// RequestValidation realiza a validação de uma estrutura de dados
func RequestValidation(request interface{}) []*IError {
	var errs []*IError
	err := Validator.Struct(request)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, err := range err.(validator.ValidationErrors) {
				var el IError
				errorSplit := strings.Split(err.Error(), "Error:")
				el.Field = err.Field()
				el.Error = errorSplit[1]
				errs = append(errs, &el)
			}
		}
	}
	return errs
}
