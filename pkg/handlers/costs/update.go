package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar um custo
type updateParams struct {
	Id string `params:"id"` // ID do custo
}

// UpdateCosts é uma função que atualiza um custo
// @Summary Rota para atualizar um custo
// @Description Rota para atualizar um custo
// @Tags Costs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Costs true "Dados do custo"
// @Param id path string true "ID do custo"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /costs/update/{id} [put]
func UpdateCosts(ctx *fiber.Ctx) error {
	cost := new(domain.Costs)
	if err := ctx.BodyParser(cost); err != nil {
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

	updated, err := repository.UpdateCosts(params.Id, cost)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Custo atualizado com sucesso",
		Success: true,
		Data:    updated,
	})
}
