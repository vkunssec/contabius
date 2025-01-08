package handlers

import (
	"github.com/vkunssec/contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"

	_ "github.com/vkunssec/contabius/docs"
)

type QueryParams struct {
	Ids []string `query:"ids"`
}

// GetBankAccount é uma função que retorna as contas bancárias
// @description Rota para retornar as contas bancárias
// @tags Account
// @accept json
// @produce json
// @security ApiKeyAuth
// @in header
// @name Authorization
func GetBankAccount(ctx *fiber.Ctx) error {
	queries := new(QueryParams)
	ctx.QueryParser(queries)

	accounts, err := repository.GetBankAccount(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(accounts)
}
