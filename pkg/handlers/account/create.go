package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/presentation"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateAccount é uma função que cria uma conta bancária
// @Summary Rota para criar uma conta bancária
// @Description Rota para criar uma conta bancária
// @Tags Account
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.AccountRequest true "Dados da conta bancária"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /account [post]
func CreateAccount(ctx *fiber.Ctx) error {
	account := new(domain.AccountRequest)
	if err := ctx.BodyParser(account); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	requestErrors := presentation.RequestValidation(account)
	if len(requestErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestErrors)
	}

	saved, err := repository.CreateBankAccount(account)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Conta bancária criada com sucesso",
		Success: true,
		Data:    saved,
	})
}
