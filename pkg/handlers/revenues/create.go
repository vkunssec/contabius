package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/presentation"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateRevenues é uma função que cria uma receita
// @Summary Rota para criar uma receita
// @Description Rota para criar uma receita
// @Tags Revenues
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.RevenueRequest true "Dados da receita"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /revenues [post]
func CreateRevenues(ctx *fiber.Ctx) error {
	revenue := new(domain.RevenueRequest)
	if err := ctx.BodyParser(revenue); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	requestErrors := presentation.RequestValidation(revenue)
	if len(requestErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestErrors)
	}

	saved, err := repository.CreateRevenues(revenue)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Receita criada com sucesso",
		Success: true,
		Data:    saved,
	})
}
