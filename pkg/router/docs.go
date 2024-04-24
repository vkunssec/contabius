package router

import "github.com/gofiber/fiber/v2"

func DocsRoutes(app fiber.Router) {
	app.Static("/docs", "docs")
}
