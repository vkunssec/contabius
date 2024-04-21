package router

import (
	handlers "contabius/pkg/handlers/account"

	"github.com/gofiber/fiber/v2"
)

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
