package router

import "github.com/gofiber/fiber/v2"

func Routes(app fiber.Router) {
	DocsRoutes(app)
	AccountRoutes(app)
}
