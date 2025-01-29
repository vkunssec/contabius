package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/revenues"
)

// RevenuesRoutes é uma função que define as rotas para a entidade Revenues
// @Summary Rota para a entidade Revenues
// @Description Rota para a entidade Revenues
func RevenuesRoutes(app fiber.Router) {
	app.Get(
		"/revenues",
		handlers.GetRevenues)
	app.Post(
		"/revenues",
		handlers.CreateRevenues)
	app.Put(
		"/revenues/update/:id",
		handlers.UpdateRevenues)
	app.Delete(
		"/revenues/remove/:id",
		handlers.DeleteRevenues)
}
