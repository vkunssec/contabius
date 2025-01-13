package common

// Conflict representa a estrutura padrão de resposta de erro de conflito
// @Description Estrutura padrão de resposta de erro de conflito
type Conflict struct {
	// Indica se a operação foi bem sucedida
	Success bool `json:"success" example:"false"`
	// Mensagem descritiva
	Message string `json:"message" example:"Conflict"`
}
