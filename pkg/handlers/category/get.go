package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
	"github.com/vkunssec/contabius/pkg/repository"

	_ "github.com/vkunssec/contabius/docs"
)

// GetByIdQueryParams é a estrutura que contém os parâmetros da requisição
type GetByIdQueryParams struct {
	Ids []string `query:"ids" example:"678079f6f5080a39a8eedc1e"` // Ids das categorias a serem retornadas
} // @name GetByIdQueryParams

// GetAccountsResponse é a estrutura que contém a resposta da requisição
type GetCategoriesResponse struct {
	Success    bool                `json:"success" example:"true"`                              // Sucesso da operação
	Message    string              `json:"message" example:"Categorias retornadas com sucesso"` // Mensagem de sucesso ou erro
	Categories []domain.Categories `json:"categories"`                                          // Dados das categorias
} // @name GetCategoriesResponse

// GetCategory é uma função que retorna todas as categorias
// @Summary Rota para retornar todas as categorias
// @Description Rota para retornar todas as categorias
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Token Bearer"
// @Param request query GetByIdQueryParams false "Query params para retornar as categorias"
// @Success 200 {object} GetCategoriesResponse
// @Failure 500 {object} common.InternalServerError
// @Router /category [get]
func GetCategory(ctx *fiber.Ctx) error {
	queries := new(GetByIdQueryParams)
	if err := ctx.QueryParser(queries); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	categories, err := repository.GetCategory(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(common.InternalServerError{
			Message: err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(GetCategoriesResponse{
		Message:    "Categorias retornadas com sucesso",
		Success:    true,
		Categories: categories,
	})
}
