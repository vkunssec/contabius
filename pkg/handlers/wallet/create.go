package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/pkg/domain/common"
)

func CreateWallet(ctx *fiber.Ctx) error {
	wallet := new(domain.Wallet)
	if err := ctx.BodyParser(wallet); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BadRequest{
			Message: err.Error(),
			Success: false,
		})
	}

	installment := 1
	now := time.Now()

	wallet.Month = time.Now().Month()
	wallet.Revenues = []domain.Revenues{{
		Revenue:   "Vylex",
		Amount:    domain.Money{Quantity: 500000, Currency: domain.CurrencyBRL},
		Method:    &domain.Methods{Method: "pix"},
		CreatedAt: now,
		UpdatedAt: now,
	}}
	wallet.Costs = []domain.Costs{{
		Cost:         "uber",
		Amount:       domain.Money{Quantity: 3600, Currency: domain.CurrencyBRL},
		Installments: &installment,
		Payment:      domain.Methods{Method: "credit"},
		Category:     domain.Categories{Category: "Transporte", CreatedAt: now, UpdatedAt: now},
		CreatedAt:    now,
		UpdatedAt:    now,
	}, {
		Cost:         "padaria",
		Amount:       domain.Money{Quantity: 1572, Currency: domain.CurrencyBRL},
		Installments: &installment,
		Payment:      domain.Methods{Method: "debit"},
		Category:     domain.Categories{Category: "Alimentação", CreatedAt: now, UpdatedAt: now},
		CreatedAt:    now,
		UpdatedAt:    now,
	}}

	wallet.CreatedAt = time.Now()
	wallet.UpdatedAt = time.Now()

	return ctx.Status(fiber.StatusOK).JSON(wallet)
}
