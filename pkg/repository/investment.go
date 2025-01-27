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

// GetInvestmentTypes retorna um investimento
func GetInvestmentTypes(ids []domain.InvestmentId) ([]domain.InvestmentType, error) {
	if len(ids) == 0 {
		return domain.AllInvestments(), nil
	}

	return domain.GetInvestments(ids)
}

func CreateInvestment(investment *domain.Investments) (domain.Investments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	investment.CreatedAt = time.Now()
	investment.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, constant.CollectionInvestment, investment)

	investment.Id = result.InsertedID.(primitive.ObjectID)

	return *investment, err
}

func GetInvestments(ids []string) ([]domain.Investments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var investments []domain.Investments
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, constant.CollectionInvestment, filters, options.Find())
	if err != nil {
		return investments, err
	}

	err = cursor.All(ctx, &investments)
	return investments, err
}

func UpdateInvestment(id string, newInvestment *domain.Investments) (domain.Investments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}
	updateFields := bson.M{}

	if newInvestment.Description != nil {
		updateFields["description"] = newInvestment.Description
	}
	if newInvestment.Recurrence != "" {
		updateFields["recurrence"] = newInvestment.Recurrence
	}
	if newInvestment.RecurrenceDay != nil {
		updateFields["recurrence_day"] = newInvestment.RecurrenceDay
	}

	if newInvestment.Investment.Id != 0 {
		updateFields["investment_type"] = newInvestment.Investment
	}
	if !newInvestment.Account.Id.IsZero() {
		updateFields["account"] = newInvestment.Account
	}
	if !newInvestment.Amount.IsZero() {
		updateFields["amount"] = newInvestment.Amount
	}

	updateFields["updated_at"] = time.Now()

	update := bson.M{"$set": updateFields}

	res, err := tools.UpdateOne(ctx, constant.CollectionInvestment, filter, update)
	if err != nil {
		return domain.Investments{}, err
	}

	if res.ModifiedCount > 0 {
		var updatedInvestment domain.Investments
		err = tools.FindOne(ctx, constant.CollectionInvestment, filter, nil).Decode(&updatedInvestment)
		if err != nil {
			return domain.Investments{}, err
		}
		return updatedInvestment, nil
	}

	return domain.Investments{}, errors.New("investimento não encontrado")
}

func DeleteInvestment(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, constant.CollectionInvestment, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
