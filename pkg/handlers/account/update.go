package handlers

import (
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/pkg/structs"

	"github.com/gofiber/fiber/v2"
)

type updateParams struct {
	Id string `params:"id"`
}

func UpdateAccount(ctx *fiber.Ctx) error {
	account := new(structs.Accounts)
	ctx.BodyParser(account)
	params := new(updateParams)
	ctx.ParamsParser(params)

	updated, err := repository.UpdateBankAccount(params.Id, account)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(updated)
}
