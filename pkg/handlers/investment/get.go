package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// GetByIdQueryParams é a estrutura que contém os parâmetros da requisição
type GetInvestmentsQueryParams struct {
	Ids []string `query:"ids" example:"678079f6f5080a39a8eedc1e"` // Ids dos investimentos a serem retornados
}

// GetInvestmentsResponse é a estrutura que contém a resposta da requisição
type GetInvestmentResponse struct {
	Success     bool                 `json:"success" example:"true"`                               // Sucesso da operação
	Message     string               `json:"message" example:"Investimento retornado com sucesso"` // Mensagem de sucesso ou erro
	Investments []domain.Investments `json:"investments"`                                          // Dados do investimento
}

// GetInvestments é uma função que retorna os investimentos
// @Summary Rota para retornar os investimentos
// @Description Rota para retornar os investimentos
// @Tags Investments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar os investimentos"
// @Success 200 {object} []domain.Investments
// @Failure 400 {object} common.BadRequest
// @Router /investments [get]
func GetInvestments(ctx *fiber.Ctx) error {
	queries := new(GetInvestmentsQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	investments, err := repository.GetInvestments(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(GetInvestmentResponse{
			Success:     false,
			Message:     "Erro ao retornar os investimentos",
			Investments: []domain.Investments{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetInvestmentResponse{
		Success:     true,
		Message:     "Investimento retornado com sucesso",
		Investments: investments,
	})
}
