package handlers

import (
	"contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type QueryParams struct {
	Ids []string `query:"ids"`
}

func GetBankAccount(ctx *fiber.Ctx) error {
	queries := new(QueryParams)
	ctx.QueryParser(queries)

	accounts, err := repository.GetBankAccount(queries.Ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(accounts)
}
