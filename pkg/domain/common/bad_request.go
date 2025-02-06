package common

// BadRequest representa a estrutura padrão de resposta de erro de requisição
// @Description Estrutura padrão de resposta de erro de requisição
type BadRequest struct {
	// Indica se a operação foi bem sucedida
	Success bool `json:"success" example:"false"`
	// Mensagem descritiva
	Message string `json:"message" example:"Erro ao processar a requisição"`
} // @name BadRequest
