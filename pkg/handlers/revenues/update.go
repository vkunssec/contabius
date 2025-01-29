package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar uma receita
type updateParams struct {
	Id string `params:"id"` // ID da receita
}

// UpdateRevenues é uma função que atualiza uma receita
// @Summary Rota para atualizar uma receita
// @Description Rota para atualizar uma receita
// @Tags Revenues
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Revenues true "Dados da receita"
// @Param id path string true "ID da receita"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /revenues/update/{id} [put]
func UpdateRevenues(ctx *fiber.Ctx) error {
	revenue := new(domain.Revenues)
	if err := ctx.BodyParser(revenue); err != nil {
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

	updated, err := repository.UpdateRevenues(params.Id, revenue)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Receita atualizada com sucesso",
		Success: true,
		Data:    updated,
	})
}
