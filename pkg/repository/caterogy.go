package repository

import (
	"context"
	"errors"
	"time"

	"github.com/vkunssec/contabius/pkg/constant"
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

	result, err := tools.InsertOne(ctx, constant.CollectionCategory, category)

	category.Id = result.InsertedID.(primitive.ObjectID)

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

	cursor, err := tools.Find(ctx, constant.CollectionCategory, filters, options.Find())
	if err != nil {
		return categories, err
	}

	err = cursor.All(ctx, &categories)
	return categories, err
}

func UpdateCategory(id string, newCategory *domain.Categories) (domain.Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	newCategory.UpdatedAt = time.Now()

	res, err := tools.UpdateOne(ctx, constant.CollectionCategory, filter, newCategory)
	if err != nil {
		return domain.Categories{}, err
	}

	if res.ModifiedCount > 0 {
		return *newCategory, nil
	}

	return domain.Categories{}, errors.New("categoria não encontrada")
}

func DeleteCategory(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, constant.CollectionCategory, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
