package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// deleteParams é a estrutura que contém os parâmetros da requisição
type deleteParams struct {
	Id string `params:"id" example:"678079f6f5080a39a8eedc1e"` // ID da conta bancária a ser deletada
} // @name deleteParams

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
// @Failure 404 {object} common.BadRequest
// @Router /account/remove/{id} [delete]
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

	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(common.BadRequest{
			Message: "Conta bancária não encontrada",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Conta bancária deletada com sucesso",
		Success: true,
		Data:    ok,
	})
}
