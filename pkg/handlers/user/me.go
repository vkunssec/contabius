package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/tools"
)

// Me é uma função que retorna o usuário logado
// @Summary Rota para retornar o usuário logado
// @Description Rota para retornar o usuário logado
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /user/me [get]
func Me(ctx *fiber.Ctx) error {
	var user domain.User

	userId := tools.GetUserId(ctx)

	user, err := repository.GetUserById(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Usuário logado com sucesso",
		Success: true,
		Data:    user,
	})
}
