package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateInvestment é uma função que cria um investimento
// @Summary Rota para criar um investimento
// @Description Rota para criar um investimento
// @Tags Investments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Investments true "Dados do investimento"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Router /investments [post]
func CreateInvestment(ctx *fiber.Ctx) error {
	investment := new(domain.Investments)
	if err := ctx.BodyParser(investment); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	saved, err := repository.CreateInvestment(investment)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Investimento criado com sucesso",
		Success: true,
		Data:    saved,
	})
}
