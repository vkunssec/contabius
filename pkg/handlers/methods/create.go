package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateMethod é uma função que cria um método
// @Summary Rota para criar um método
// @Description Rota para criar um método
// @Tags Methods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Methods true "Dados do método"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Failure 500 {object} common.InternalServerError
// @Router /methods [post]
func CreateMethod(ctx *fiber.Ctx) error {
	method := new(domain.Methods)
	if err := ctx.BodyParser(method); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	saved, err := repository.CreateMethod(method)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.InternalServerError{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Método criado com sucesso",
		Success: true,
		Data:    saved,
	})
}
