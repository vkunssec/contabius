package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Common(app *fiber.App) {
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(healthcheck.New())

	app.Use(logger.New(Logger()))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}
