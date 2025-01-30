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

func CreateCosts(request *domain.CostRequest) (domain.Costs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cost domain.Costs
	var category domain.Categories

	cost.Cost = request.Cost
	cost.Amount = request.Amount
	if request.Installments > 0 {
		cost.Installments = request.Installments
	}
	if request.Methods != nil {
		cost.Methods = request.Methods
	}

	category.Id = request.Category.Id
	category.Category = request.Category.Category
	cost.Category = category

	cost.CreatedAt = time.Now()
	cost.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, constant.CollectionCosts, cost)

	cost.Id = result.InsertedID.(primitive.ObjectID)

	return cost, err
}

func GetCosts(ids []string) ([]domain.Costs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var costs []domain.Costs
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, constant.CollectionCosts, filters, options.Find())
	if err != nil {
		return costs, err
	}

	err = cursor.All(ctx, &costs)
	return costs, err
}

func UpdateCosts(id string, newCost *domain.Costs) (domain.Costs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}
	updateFields := bson.M{}

	if newCost.Cost != "" {
		updateFields["cost"] = newCost.Cost
	}
	if !newCost.Amount.IsZero() {
		updateFields["amount"] = newCost.Amount
	}

	if newCost.Methods != nil {
		updateFields["methods"] = newCost.Methods
	}

	if newCost.Category != (domain.Categories{}) {
		updateFields["category"] = newCost.Category
	}

	updateFields["updated_at"] = time.Now()

	update := bson.M{"$set": updateFields}

	res, err := tools.UpdateOne(ctx, constant.CollectionCosts, filter, update)
	if err != nil {
		return domain.Costs{}, err
	}

	if res.ModifiedCount > 0 {
		var updatedCost domain.Costs
		err = tools.FindOne(ctx, constant.CollectionCosts, filter, nil).Decode(&updatedCost)
		if err != nil {
			return domain.Costs{}, err
		}
		return updatedCost, nil
	}

	return domain.Costs{}, errors.New("custo não encontrado")
}

func DeleteCosts(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, constant.CollectionCosts, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
