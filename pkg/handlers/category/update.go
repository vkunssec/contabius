package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar uma categoria
type updateParams struct {
	Id string `params:"id"` // ID da categoria
} // @name updateParams

// UpdateCategory é uma função que atualiza uma categoria
// @Summary Rota para atualizar uma categoria
// @Description Rota para atualizar uma categoria
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Categories true "Dados da categoria"
// @Param id path string true "ID da categoria"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /category/update/{id} [put]
func UpdateCategory(ctx *fiber.Ctx) error {
	category := new(domain.Categories)
	if err := ctx.BodyParser(category); err != nil {
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

	updated, err := repository.UpdateCategory(params.Id, category)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Categoria atualizada com sucesso",
		Success: true,
		Data:    updated,
	})
}
