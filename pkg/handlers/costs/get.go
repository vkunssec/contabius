package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// GetByIdQueryParams é a estrutura que contém os parâmetros da requisição
type GetByIdQueryParams struct {
	Ids []string `query:"ids" example:"678079f6f5080a39a8eedc1e"` // Ids dos custos a serem retornados
} // @name GetByIdQueryParams

// GetAccountsResponse é a estrutura que contém a resposta da requisição
type GetAccountsResponse struct {
	Success  bool           `json:"success" example:"true"`                        // Sucesso da operação
	Message  string         `json:"message" example:"Custo retornado com sucesso"` // Mensagem de sucesso ou erro
	Accounts []domain.Costs `json:"accounts"`                                      // Dados do custo
} // @name GetAccountsResponse

// GetCosts é uma função que retorna os custos
// @Summary Rota para retornar os custos
// @Description Rota para retornar os custos
// @Tags Costs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar os custos"
// @Success 200 {object} []domain.Costs
// @Failure 400 {object} common.BadRequest
// @Router /costs [get]
func GetCosts(ctx *fiber.Ctx) error {
	queries := new(GetByIdQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	costs, err := repository.GetCosts(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(GetAccountsResponse{
			Success:  false,
			Message:  "Erro ao retornar os custos",
			Accounts: []domain.Costs{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetAccountsResponse{
		Success:  true,
		Message:  "Custo retornado com sucesso",
		Accounts: costs,
	})
}
