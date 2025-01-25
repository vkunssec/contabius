package router

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	AccountRoutes(app)
	CategoryRoutes(app)
	MethodsRoutes(app)
}
