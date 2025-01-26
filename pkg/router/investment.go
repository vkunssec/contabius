package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/investment"
)

// InvestmentsRoutes é uma função que define as rotas para a entidade Investments
// @Summary Rota para a entidade Investments
// @Description Rota para a entidade Investments
func InvestmentsRoutes(app fiber.Router) {
	app.Get("/investments/get_types", handlers.GetInvestments)
}
