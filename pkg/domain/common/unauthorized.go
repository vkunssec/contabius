package common

// Unauthorized representa a estrutura padrão de resposta de erro de requisição
// @Description Estrutura padrão de resposta de erro de requisição
type Unauthorized struct {
	// Indica se a operação foi bem sucedida
	Success bool `json:"success" example:"false"`
	// Mensagem descritiva
	Message string `json:"message" example:"Invalid or expired JWT"`
} // @name Unauthorized
