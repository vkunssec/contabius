package repository

import (
	"errors"

	"github.com/vkunssec/contabius/pkg/constant"
	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/tools"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMethod(method *domain.Methods) (domain.Methods, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	method.CreatedAt = time.Now()
	method.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, constant.CollectionMethod, method)

	method.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return *method, err
}

func GetMethod(ids []string) ([]domain.Methods, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var methods []domain.Methods
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, constant.CollectionMethod, filters, options.Find())
	if err != nil {
		return methods, err
	}

	err = cursor.All(ctx, &methods)
	return methods, err
}

func UpdateMethod(id string, newMethod *domain.Methods) (domain.Methods, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	newMethod.UpdatedAt = time.Now()

	res, err := tools.UpdateOne(ctx, constant.CollectionMethod, filter, newMethod)
	if err != nil {
		return domain.Methods{}, err
	}

	if res.ModifiedCount > 0 {
		return *newMethod, nil
	}

	return domain.Methods{}, errors.New("método não encontrado")
}

func DeleteMethod(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, constant.CollectionMethod, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
