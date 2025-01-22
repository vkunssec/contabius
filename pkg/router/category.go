package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/vkunssec/contabius/pkg/handlers/category"

	_ "github.com/vkunssec/contabius/docs"
)

// CategoryRoutes é uma função que define as rotas para a entidade Category
// @Summary Rota para a entidade Category
// @Description Rota para a entidade Category
func CategoryRoutes(app fiber.Router) {
	app.Get(
		"/category",
		handlers.GetCategory)
	app.Post(
		"/category",
		handlers.CreateCategory)
	// app.Put(
	// 	"/category/update/:id",
	// 	handlers.UpdateCategory)
	// app.Delete(
	// 	"/category/remove/:id",
	// 	handlers.DeleteCategory)
}
