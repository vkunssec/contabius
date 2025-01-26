package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/wallet"
)

func Routes(app *fiber.App) {
	AccountRoutes(app)
	CategoryRoutes(app)
	MethodsRoutes(app)
	InvestmentsRoutes(app)

	// teste
	app.Post("/wallet", handlers.CreateWallet)
}
