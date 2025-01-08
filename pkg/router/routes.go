package router

import "github.com/gofiber/fiber/v2"

func Routes(app fiber.Router) {
	AccountRoutes(app)
}
