package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/presentation"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateCosts é uma função que cria um custo
// @Summary Rota para criar um custo
// @Description Rota para criar um custo
// @Tags Costs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.CostRequest true "Dados do custo"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /costs [post]
func CreateCosts(ctx *fiber.Ctx) error {
	cost := new(domain.CostRequest)
	if err := ctx.BodyParser(cost); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	requestErrors := presentation.RequestValidation(cost)
	if len(requestErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestErrors)
	}

	saved, err := repository.CreateCosts(cost)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Custo criado com sucesso",
		Success: true,
		Data:    saved,
	})
}
