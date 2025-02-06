package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	handlers "github.com/vkunssec/contabius/pkg/handlers/user"
	"github.com/vkunssec/contabius/pkg/middleware"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user/signup", handlers.SignUp)
	app.Post("/user/login", handlers.Login)
	app.Get("/user/me",
		middleware.Protected([]domain.UserType{
			domain.UserTypeAdmin,
			// domain.UserTypeUser,
		}),
		handlers.Me)
	app.Put("/user/update/:id", handlers.UpdateUser)
	app.Put("/user/update/password", handlers.UpdateUserPassword)
}
