package repository

import (
	"context"
	"time"

	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetCategory(ids []string) ([]domain.Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var categories []domain.Categories
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, Collection, filters, options.Find())
	if err != nil {
		return categories, err
	}

	err = cursor.All(ctx, &categories)
	return categories, err
}
