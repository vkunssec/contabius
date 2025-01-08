package handlers

import (
	"time"

	"github.com/vkunssec/contabius/pkg/structs"

	"github.com/gofiber/fiber/v2"
)

func CreateWallet(ctx *fiber.Ctx) error {
	wallet := new(structs.Wallet)
	ctx.BodyParser(wallet)

	installment := 1
	now := time.Now()

	wallet.Month = time.Now().Month()
	wallet.Revenues = []structs.Revenues{{
		Revenue:   "Vylex",
		Amount:    structs.Money{Quantity: 506890, Currency: "BRL"},
		CreatedAt: now,
		UpdatedAt: now,
	}}
	wallet.Costs = []structs.Costs{{
		Cost:         "uber",
		Amount:       structs.Money{Quantity: 3600, Currency: "BRL"},
		Installments: &installment,
		Payment:      structs.Methods{Method: "credit", CreatedAt: now, UpdatedAt: now},
		Category:     structs.Categories{Grade: "Transporte", CreatedAt: now, UpdatedAt: now},
		CreatedAt:    now,
		UpdatedAt:    now,
	}, {
		Cost:         "padaria",
		Amount:       structs.Money{Quantity: 1572, Currency: "BRL"},
		Installments: &installment,
		Payment:      structs.Methods{Method: "debit", CreatedAt: now, UpdatedAt: now},
		Category:     structs.Categories{Grade: "Alimentação", CreatedAt: now, UpdatedAt: now},
		CreatedAt:    now,
		UpdatedAt:    now,
	}}

	wallet.CreatedAt = time.Now()
	wallet.UpdatedAt = time.Now()

	return ctx.Status(fiber.StatusOK).JSON(wallet)
}
