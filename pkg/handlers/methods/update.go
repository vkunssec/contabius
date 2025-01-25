package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar um método
type updateParams struct {
	Id string `params:"id"` // ID do método
}

// UpdateMethod é uma função que atualiza um método
// @Summary Rota para atualizar um método
// @Description Rota para atualizar um método
// @Tags Methods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Methods true "Dados do método"
// @Param id path string true "ID do método"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /methods/update/{id} [put]
func UpdateMethod(ctx *fiber.Ctx) error {
	method := new(domain.Methods)
	if err := ctx.BodyParser(method); err != nil {
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

	updated, err := repository.UpdateMethod(params.Id, method)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Método atualizado com sucesso",
		Success: true,
		Data:    updated,
	})
}
