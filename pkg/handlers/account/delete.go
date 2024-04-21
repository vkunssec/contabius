package handlers

import (
	"contabius/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type deleteParams struct {
	Id string `params:"id"`
}

func DeleteAccount(ctx *fiber.Ctx) error {
	params := new(deleteParams)
	ctx.ParamsParser(params)

	ok, err := repository.DeleteBankAccount(params.Id)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}
