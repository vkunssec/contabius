package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/user"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user/signup", handlers.SignUp)
	app.Post("/user/login", handlers.Login)
	app.Get("/user/me", handlers.Me)
	app.Put("/user/update/:id", handlers.UpdateUser)
	app.Put("/user/update/password", handlers.UpdateUserPassword)
}
