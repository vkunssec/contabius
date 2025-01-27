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
	Ids []domain.InvestmentId `query:"ids" example:"0,1,2,3"` // Ids dos investimentos a serem retornados
}

// GetMethodsResponse é a estrutura que contém a resposta da requisição
type GetInvestmentsResponse struct {
	Success bool                    `json:"success" example:"true"`                                 // Sucesso da operação
	Message string                  `json:"message" example:"Investimentos retornados com sucesso"` // Mensagem de sucesso ou erro
	Types   []domain.InvestmentType `json:"types"`                                                  // Dados dos investimentos
}

// GetTypesInvestments é uma função que retorna todos os tipos de investimentos
// @Summary Rota para retornar todos os tipos de investimentos
// @Description Rota para retornar todos os tipos de investimentos
// @Tags Investments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar os investimentos"
// @Success 200 {object} GetInvestmentsResponse
// @Failure 500 {object} common.InternalServerError
// @Router /investments/get_types [get]
func GetTypesInvestments(ctx *fiber.Ctx) error {
	queries := new(GetByIdQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	investments, err := repository.GetInvestmentTypes(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.InternalServerError{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetInvestmentsResponse{
		Message: "Investimentos retornados com sucesso",
		Success: true,
		Types:   investments,
	})
}
