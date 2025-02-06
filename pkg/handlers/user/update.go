package handlers

import (
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// updateParams é a estrutura que contém os parâmetros da requisição
// @Description Parâmetros da requisição para atualizar um usuário
type updateParams struct {
	Id string `params:"id"` // ID do usuário
} // @name updateParams

// UpdateUser é uma função que atualiza um usuário
// @Summary Rota para atualizar um usuário
// @Description Rota para atualizar um usuário
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.UserUpdateRequest true "Dados do usuário"
// @Param id path string true "ID do usuário"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /user/update/{id} [put]
func UpdateUser(ctx *fiber.Ctx) error {
	user := new(domain.UserUpdateRequest)
	if err := ctx.BodyParser(user); err != nil {
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

	updated, err := repository.UpdateUser(params.Id, *user)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Usuário atualizado com sucesso",
		Success: true,
		Data:    updated,
	})
}
