package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// deleteParams é a estrutura que contém os parâmetros da requisição
type deleteParams struct {
	Id string `params:"id" example:"678079f6f5080a39a8eedc1e"` // ID da receita a ser deletada
} // @name deleteParams

// DeleteRevenues é uma função que deleta uma receita
// @Summary Rota para deletar uma receita
// @Description Rota para deletar uma receita
// @Tags Revenues
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param id path string true "ID da receita"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Failure 404 {object} common.BadRequest
// @Router /revenues/remove/{id} [delete]
func DeleteRevenues(ctx *fiber.Ctx) error {
	params := new(deleteParams)
	if err := ctx.ParamsParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	ok, err := repository.DeleteRevenues(params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(common.BadRequest{
			Message: "Receita não encontrada",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Receita deletada com sucesso",
		Success: true,
		Data:    ok,
	})
}
