package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/tools"
)

// UpdateUserPassword é uma função que atualiza a senha de um usuário
// @Summary Rota para atualizar a senha de um usuário
// @Description Rota para atualizar a senha de um usuário
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.UserUpdatePassword true "Dados da senha"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /user/update/password/{id} [put]
func UpdateUserPassword(ctx *fiber.Ctx) error {
	user := new(domain.UserUpdatePassword)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	userId := tools.GetUserId(ctx)
	if userId == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.BadRequest{
			Message: "Usuário não autenticado",
			Success: false,
		})
	}

	updated, err := repository.UpdateUserPassword(userId, *user)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Senha atualizada com sucesso",
		Success: true,
		Data:    updated,
	})
}
