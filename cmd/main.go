package main

import (
	"context"
	"log"

	"github.com/vkunssec/contabius/configs"
	"github.com/vkunssec/contabius/database"
	"github.com/vkunssec/contabius/pkg/middleware"
	"github.com/vkunssec/contabius/pkg/router"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

// @title Contabius API
// @version 1.0
// @description API para gerenciamento de contas bancárias
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://vkunssec.vercel.app

// @license.name MIT
// @license.url https://github.com/vkunssec/contabius/blob/main/LICENSE

// @host ${HOST}
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer token authentication
func main() {
	ctx := context.Background()
	openConnections(ctx)

	app := startServer()
	app.Listen(":" + configs.Env("PORT"))
}

func openConnections(ctx context.Context) {
	err := database.MongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  configs.ServerName,
	})

	middleware.Common(app)
	middleware.Swagger(app)

	router.Routes(app)

	return app
}
