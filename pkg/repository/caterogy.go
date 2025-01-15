package repository

import (
	"context"
	"time"

	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(category *domain.Categories) (domain.Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, Collection, category)

	category.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return *category, err
}
