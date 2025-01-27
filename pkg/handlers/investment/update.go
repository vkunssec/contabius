package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar um investimento
type updateParams struct {
	Id string `params:"id"` // ID do investimento
}

// UpdateInvestment é uma função que atualiza um investimento
// @Summary Rota para atualizar um investimento
// @Description Rota para atualizar um investimento
// @Tags Investments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Investments true "Dados do investimento"
// @Param id path string true "ID do investimento"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /investment/update/{id} [put]
func UpdateInvestment(ctx *fiber.Ctx) error {
	investment := new(domain.Investments)
	if err := ctx.BodyParser(investment); err != nil {
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

	updated, err := repository.UpdateInvestment(params.Id, investment)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Investimento atualizado com sucesso",
		Success: true,
		Data:    updated,
	})
}
