package router

import (
	accountHandlers "contabius/pkg/handlers/account"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Get(
		"/account",
		accountHandlers.GetBankAccount)
	app.Post(
		"/account/create",
		accountHandlers.CreateAccount)
	app.Put(
		"/account/update/:id",
		accountHandlers.UpdateAccount)
}
