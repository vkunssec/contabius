package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar uma conta bancária
type updateParams struct {
	Id string `params:"id"` // ID da conta bancária
} // @name updateParams

// UpdateAccount é uma função que atualiza uma conta bancária
// @Summary Rota para atualizar uma conta bancária
// @Description Rota para atualizar uma conta bancária
// @Tags Account
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Accounts true "Dados da conta bancária"
// @Param id path string true "ID da conta bancária"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /account/update/{id} [put]
func UpdateAccount(ctx *fiber.Ctx) error {
	account := new(domain.Accounts)
	if err := ctx.BodyParser(account); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}
	params := new(updateParams)
	if err := ctx.ParamsParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	updated, err := repository.UpdateBankAccount(params.Id, account)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Conta bancária atualizada com sucesso",
		Success: true,
		Data:    updated,
	})
}
