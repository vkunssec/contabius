package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// GetByIdQueryParams é a estrutura que contém os parâmetros da requisição
type GetByIdQueryParams struct {
	Ids []string `query:"ids" example:"678079f6f5080a39a8eedc1e"` // Ids dos métodos a serem retornados
}

// GetMethodsResponse é a estrutura que contém a resposta da requisição
type GetMethodsResponse struct {
	Success bool             `json:"success" example:"true"`                           // Sucesso da operação
	Message string           `json:"message" example:"Métodos retornados com sucesso"` // Mensagem de sucesso ou erro
	Methods []domain.Methods `json:"methods"`                                          // Dados dos métodos
}

// GetMethods é uma função que retorna todos os métodos
// @Summary Rota para retornar todos os métodos
// @Description Rota para retornar todos os métodos
// @Tags Methods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar os métodos"
// @Success 200 {object} GetMethodsResponse
// @Failure 500 {object} common.InternalServerError
// @Router /methods [get]
func GetMethods(ctx *fiber.Ctx) error {
	queries := new(GetByIdQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	methods, err := repository.GetMethod(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.InternalServerError{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetMethodsResponse{
		Message: "Métodos retornados com sucesso",
		Success: true,
		Methods: methods,
	})
}
