package common

// InternalServerError representa a estrutura padrão de resposta de erro de servidor interno
// @Description Estrutura padrão de resposta de erro de servidor interno
type InternalServerError struct {
	// Indica se a operação foi bem sucedida
	Success bool `json:"success" example:"false"`
	// Mensagem descritiva
	Message string `json:"message" example:"Internal Server Error"`
} // @name InternalServerError
