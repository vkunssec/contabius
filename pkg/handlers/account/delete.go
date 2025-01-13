package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

// deleteParams é a estrutura que contém os parâmetros da requisição
type deleteParams struct {
	Id string `params:"id" example:"678079f6f5080a39a8eedc1e"` // ID da conta bancária a ser deletada
}

// DeleteAccount é uma função que deleta uma conta bancária
// @Summary Rota para deletar uma conta bancária
// @Description Rota para deletar uma conta bancária
// @Tags Account
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param id path string true "ID da conta bancária"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /account/delete/{id} [delete]
func DeleteAccount(ctx *fiber.Ctx) error {
	params := new(deleteParams)
	if err := ctx.ParamsParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	ok, err := repository.DeleteBankAccount(params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Conta bancária deletada com sucesso",
		Success: true,
		Data:    ok,
	})
}
