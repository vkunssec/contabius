package main

import (
	"context"
	"time"

	"github.com/vkunssec/contabius/database"
	"github.com/vkunssec/contabius/pkg/constant"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Accounts = []domain.Accounts{
	{
		Account: "Santander",
		Color:   "#cc0000",
	},
	{
		Account: "Will Bank",
		Color:   "#FFD900",
	},
	{
		Account: "Neon",
		Color:   "#0f92ff",
	},
	{
		Account: "XP Investimento",
		Color:   "#121212",
	},
	{
		Account: "Inter",
		Color:   "#ff7a00",
	},
	{
		Account: "Nubank",
		Color:   "#820ad1",
	},
}

func seedAccounts() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingAccounts := make(map[string]bool)
	cursor, err := database.MongoDB.Collection(constant.CollectionBank).Find(ctx, bson.M{})
	if err == nil {
		var accounts []domain.Accounts
		if err := cursor.All(ctx, &accounts); err == nil {
			for _, acc := range accounts {
				existingAccounts[acc.Account] = true
			}
		}
	}

	var newAccounts []interface{}
	for _, account := range Accounts {
		if !existingAccounts[account.Account] {
			account.CreatedAt = time.Now()
			account.UpdatedAt = time.Now()
			newAccounts = append(newAccounts, account)
		}
	}

	if len(newAccounts) > 0 {
		_, err := database.MongoDB.Collection(constant.CollectionBank).InsertMany(ctx, newAccounts)
		if err != nil {
			logger.Logger.Error().Err(err).Send()
		}
	}
}

var Categories = []domain.Categories{
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1624"),
		Category: "Alimentação",
		Parent:   nil,
	},
	{
		Category: "Restaurante",
		Parent:   newObjectId("67946a6512dc7ec3301f1624"),
	},
	{
		Category: "Supermercado",
		Parent:   newObjectId("67946a6512dc7ec3301f1624"),
	},
	{
		Category: "Delivery",
		Parent:   newObjectId("67946a6512dc7ec3301f1624"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1625"),
		Category: "Transporte",
		Parent:   nil,
	},
	{
		Category: "Combustível",
		Parent:   newObjectId("67946a6512dc7ec3301f1625"),
	},
	{
		Category: "Uber/99",
		Parent:   newObjectId("67946a6512dc7ec3301f1625"),
	},
	{
		Category: "Manutenção",
		Parent:   newObjectId("67946a6512dc7ec3301f1625"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1626"),
		Category: "Lazer",
		Parent:   nil,
	},
	{
		Category: "Cinema",
		Parent:   newObjectId("67946a6512dc7ec3301f1626"),
	},
	{
		Category: "Shows",
		Parent:   newObjectId("67946a6512dc7ec3301f1626"),
	},
	{
		Category: "Viagens",
		Parent:   newObjectId("67946a6512dc7ec3301f1626"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1627"),
		Category: "Moradia",
		Parent:   nil,
	},
	{
		Category: "Aluguel",
		Parent:   newObjectId("67946a6512dc7ec3301f1627"),
	},
	{
		Category: "Contas",
		Parent:   newObjectId("67946a6512dc7ec3301f1627"),
	},
	{
		Category: "Manutenção",
		Parent:   newObjectId("67946a6512dc7ec3301f1627"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1628"),
		Category: "Saúde",
		Parent:   nil,
	},
	{
		Category: "Médico",
		Parent:   newObjectId("67946a6512dc7ec3301f1628"),
	},
	{
		Category: "Farmácia",
		Parent:   newObjectId("67946a6512dc7ec3301f1628"),
	},
	{
		Category: "Plano de Saúde",
		Parent:   newObjectId("67946a6512dc7ec3301f1628"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1629"),
		Category: "Educação",
		Parent:   nil,
	},
	{
		Category: "Cursos",
		Parent:   newObjectId("67946a6512dc7ec3301f1629"),
	},
	{
		Category: "Material Escolar",
		Parent:   newObjectId("67946a6512dc7ec3301f1629"),
	},
	{
		Category: "Mensalidade",
		Parent:   newObjectId("67946a6512dc7ec3301f1629"),
	},
	{
		Id:       *newObjectId("67946a6512dc7ec3301f1630"),
		Category: "Outros",
		Parent:   nil,
	},
}

func seedCategories() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingCategories := make(map[string]bool)
	cursor, err := database.MongoDB.Collection(constant.CollectionCategory).Find(ctx, bson.M{})
	if err == nil {
		var categories []domain.Categories
		if err := cursor.All(ctx, &categories); err == nil {
			for _, cat := range categories {
				existingCategories[cat.Category] = true
			}
		}
	}

	var newCategories []interface{}
	for _, category := range Categories {
		if !existingCategories[category.Category] {
			category.CreatedAt = time.Now()
			category.UpdatedAt = time.Now()
			newCategories = append(newCategories, category)
		}
	}

	if len(newCategories) > 0 {
		_, err := database.MongoDB.Collection(constant.CollectionCategory).InsertMany(ctx, newCategories)
		if err != nil {
			logger.Logger.Error().Err(err).Send()
		}
	}
}

func openConnections(ctx context.Context) {
	err := database.MongoDBConnection(ctx)
	if err != nil {
		logger.Logger.Error().Err(err).Send()
	}
}

func newObjectId(id string) *primitive.ObjectID {
	s, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		logger.Logger.Error().Err(e).Send()
	}
	return &s
}

func main() {
	ctx := context.Background()
	openConnections(ctx)

	seedAccounts()

	seedCategories()
}
