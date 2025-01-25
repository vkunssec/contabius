package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/methods"

	_ "github.com/vkunssec/contabius/docs"
)

// MethodsRoutes é uma função que define as rotas para a entidade Methods
// @Summary Rota para a entidade Methods
// @Description Rota para a entidade Methods
func MethodsRoutes(app fiber.Router) {
	app.Get(
		"/methods",
		handlers.GetMethods)
}
