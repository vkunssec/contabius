package common

// Response representa a estrutura padrão de resposta da API
// @Description Estrutura padrão de resposta
type Response struct {
	// Indica se a operação foi bem sucedida
	Success bool `json:"success" example:"true"`
	// Mensagem descritiva
	Message string `json:"message" example:"Operação realizada com sucesso"`
	// Dados da resposta
	Data interface{} `json:"data,omitempty"`
} // @name Response
