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

func CreateRevenues(revenue *domain.Revenues) (domain.Revenues, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	revenue.CreatedAt = time.Now()
	revenue.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, constant.CollectionRevenues, revenue)

	revenue.Id = result.InsertedID.(primitive.ObjectID)

	return *revenue, err
}

func GetRevenues(ids []string) ([]domain.Revenues, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var revenues []domain.Revenues
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, constant.CollectionRevenues, filters, options.Find())
	if err != nil {
		return revenues, err
	}

	err = cursor.All(ctx, &revenues)
	return revenues, err
}

func UpdateRevenues(id string, newRevenue *domain.Revenues) (domain.Revenues, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}
	updateFields := bson.M{}

	if newRevenue.Revenue != "" {
		updateFields["revenue"] = newRevenue.Revenue
	}
	if !newRevenue.Amount.IsZero() {
		updateFields["amount"] = newRevenue.Amount
	}

	if newRevenue.Method != nil {
		updateFields["method"] = newRevenue.Method
	}

	updateFields["updated_at"] = time.Now()

	update := bson.M{"$set": updateFields}

	res, err := tools.UpdateOne(ctx, constant.CollectionRevenues, filter, update)
	if err != nil {
		return domain.Revenues{}, err
	}

	if res.ModifiedCount > 0 {
		var updatedRevenue domain.Revenues
		err = tools.FindOne(ctx, constant.CollectionRevenues, filter, nil).Decode(&updatedRevenue)
		if err != nil {
			return domain.Revenues{}, err
		}
		return updatedRevenue, nil
	}

	return domain.Revenues{}, errors.New("receita não encontrada")
}

func DeleteRevenues(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, constant.CollectionRevenues, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
