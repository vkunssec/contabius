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
	Ids []string `query:"ids" example:"678079f6f5080a39a8eedc1e"` // Ids das contas bancárias a serem retornadas
} // @name GetByIdQueryParams

// GetAccountsResponse é a estrutura que contém a resposta da requisição
type GetAccountsResponse struct {
	Success  bool              `json:"success" example:"true"`                                 // Sucesso da operação
	Message  string            `json:"message" example:"Conta bancária retornada com sucesso"` // Mensagem de sucesso ou erro
	Accounts []domain.Accounts `json:"accounts"`                                               // Dados da conta bancária
} // @name GetAccountsResponse

// GetBankAccount é uma função que retorna as contas bancárias
// @Summary Rota para retornar as contas bancárias
// @Description Rota para retornar as contas bancárias
// @Tags Account
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar as contas bancárias"
// @Success 200 {object} []domain.Accounts
// @Failure 400 {object} common.BadRequest
// @Router /account [get]
func GetBankAccount(ctx *fiber.Ctx) error {
	queries := new(GetByIdQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	accounts, err := repository.GetBankAccount(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(GetAccountsResponse{
			Success:  false,
			Message:  "Erro ao retornar as contas bancárias",
			Accounts: []domain.Accounts{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetAccountsResponse{
		Success:  true,
		Message:  "Conta bancária retornada com sucesso",
		Accounts: accounts,
	})
}
