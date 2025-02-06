package common

// NotFound representa a estrutura padrão de resposta de erro não encontrado
// @Description Estrutura padrão de resposta de erro não encontrado
type NotFound struct {
	// Mensagem de erro
	Message string `json:"message" example:"Resource not found"`
	Success bool   `json:"success" example:"false"`
} // @name NotFound
