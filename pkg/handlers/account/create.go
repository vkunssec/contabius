package handlers

import (
	"github.com/vkunssec/contabius/pkg/repository"
	"github.com/vkunssec/contabius/pkg/structs"

	"github.com/gofiber/fiber/v2"
)

func CreateAccount(ctx *fiber.Ctx) error {
	account := new(structs.Accounts)
	ctx.BodyParser(account)

	saved, err := repository.CreateBankAccount(account)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(saved)
}
