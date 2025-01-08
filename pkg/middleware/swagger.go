package middleware

import (
	"log"
	"strings"

	"github.com/vkunssec/contabius/configs"
	"github.com/vkunssec/contabius/docs"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
)

// Swagger é um middleware que serve a interface Swagger UI para a documentação da API.
func Swagger(app *fiber.App) {
	if configs.Env("STAGE") != "development" {
		return
	}

	log.Println("Swagger is enabled")

	swaggerJSON := docs.SwaggerInfo.ReadDoc()

	log.Println("Swagger JSON: ", swaggerJSON)

	app.Get("/docs/swagger.json", func(c *fiber.Ctx) error {
		modifiedJSON := strings.Replace(swaggerJSON, "${HOST}", configs.Host, -1)
		return c.SendString(modifiedJSON)
	})

	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		host := ctx.BaseURL()

		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: host + "/docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "CofCards API Documentation",
			},
			DarkMode: true,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		ctx.Set("Content-Type", "text/html; charset=utf-8")
		return ctx.SendString(html)
	})
}
