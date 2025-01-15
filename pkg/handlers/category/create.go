package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// CreateCategory é uma função que cria uma categoria
// @Summary Rota para criar uma categoria
// @Description Rota para criar uma categoria
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request body domain.Categories true "Dados da categoria"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.BadRequest
// @Failure 500 {object} common.InternalServerError
// @Router /category [post]
func CreateCategory(ctx *fiber.Ctx) error {
	category := new(domain.Categories)
	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	saved, err := repository.CreateCategory(category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.InternalServerError{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.Response{
		Message: "Categoria criada com sucesso",
		Success: true,
		Data:    saved,
	})
}
