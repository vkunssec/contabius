package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// deleteParams é a estrutura que contém os parâmetros da requisição
type deleteParams struct {
	Id string `params:"id" example:"678079f6f5080a39a8eedc1e"` // ID do custo a ser deletado
} // @name deleteParams

// DeleteCosts é uma função que deleta um custo
// @Summary Rota para deletar um custo
// @Description Rota para deletar um custo
// @Tags Costs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param id path string true "ID do custo"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Failure 404 {object} common.BadRequest
// @Router /costs/remove/{id} [delete]
func DeleteCosts(ctx *fiber.Ctx) error {
	params := new(deleteParams)
	if err := ctx.ParamsParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	ok, err := repository.DeleteCosts(params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(common.BadRequest{
			Message: "Custo não encontrado",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Custo deletado com sucesso",
		Success: true,
		Data:    ok,
	})
}
