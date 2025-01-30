package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/presentation"

	_ "github.com/vkunssec/contabius/docs"
)

// Login é uma função que loga um usuário
// @Summary Rota para logar um usuário
// @Description Rota para logar um usuário
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.UserRequest true "Dados do usuário"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /user/login [post]
func Login(ctx *fiber.Ctx) error {
	user := new(domain.UserLogin)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	requestErrors := presentation.RequestValidation(user)
	if len(requestErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestErrors)
	}

	saved, err := repository.LoginUser(*user)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Usuário logado com sucesso",
		Success: true,
		Data:    saved,
	})
}
