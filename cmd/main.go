package main

import (
	"contabius/configs"
	"contabius/database"
	"contabius/pkg/middleware"
	"contabius/pkg/router"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	openConnections(ctx)

	app := startServer()
	app.Listen(":" + configs.Env("PORT"))
}

func openConnections(ctx context.Context) {
	database.MongoDBConnection(ctx)
}

func startServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  configs.ServerName,
	})

	middleware.Common(app)
	router.Routes(app)

	return app
}
