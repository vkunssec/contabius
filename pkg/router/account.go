package router

import (
	handlers "github.com/vkunssec/contabius/pkg/handlers/account"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// AccountRoutes é uma função que define as rotas para a entidade Account
// @description Rota para a entidade Account
// @tags Account
// @accept json
// @produce json
// @security ApiKeyAuth
// @in header
// @name Authorization
func AccountRoutes(app fiber.Router) {
	app.Get(
		"/account",
		handlers.GetBankAccount)
	app.Post(
		"/account/create",
		handlers.CreateAccount)
	app.Put(
		"/account/update/:id",
		handlers.UpdateAccount)
	app.Delete(
		"/account/remove/:id",
		handlers.DeleteAccount)
}
