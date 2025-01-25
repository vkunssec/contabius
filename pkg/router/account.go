package router

import (
	handlers "github.com/vkunssec/contabius/pkg/handlers/account"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// AccountRoutes é uma função que define as rotas para a entidade Account
// @Summary Rota para a entidade Account
// @Description Rota para a entidade Account
func AccountRoutes(app *fiber.App) {
	app.Get(
		"/account",
		handlers.GetBankAccount)
	app.Post(
		"/account",
		handlers.CreateAccount)
	app.Put(
		"/account/update/:id",
		handlers.UpdateAccount)
	app.Delete(
		"/account/remove/:id",
		handlers.DeleteAccount)
}
