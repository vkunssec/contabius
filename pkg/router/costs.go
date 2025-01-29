package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/costs"
)

// CostsRoutes é uma função que define as rotas para a entidade Revenues
// @Summary Rota para a entidade Revenues
// @Description Rota para a entidade Revenues
func CostsRoutes(app fiber.Router) {
	app.Get(
		"/costs",
		handlers.GetCosts)
	app.Post(
		"/costs",
		handlers.CreateCosts)
	app.Put(
		"/costs/update/:id",
		handlers.UpdateCosts)
	app.Delete(
		"/costs/remove/:id",
		handlers.DeleteCosts)
}
