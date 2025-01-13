package repository

import (
	"errors"

	"github.com/vkunssec/contabius/pkg/domain"
	"github.com/vkunssec/contabius/tools"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Collection = "bank_account"
)

func CreateBankAccount(account *domain.Accounts) (domain.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, Collection, account)

	account.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return *account, err
}

func GetBankAccount(ids []string) ([]domain.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var accounts []domain.Accounts
	filters := bson.M{}

	if len(ids) > 0 {
		filters["_id"] = bson.M{
			"$in": tools.ArrayStringToObjectId(ids),
		}
	}

	cursor, err := tools.Find(ctx, Collection, filters, options.Find())
	if err != nil {
		return accounts, err
	}

	err = cursor.All(ctx, &accounts)
	return accounts, err
}

func UpdateBankAccount(id string, newAccount *domain.Accounts) (domain.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	newAccount.UpdatedAt = time.Now()

	res, err := tools.UpdateOne(ctx, Collection, filter, newAccount)
	if err != nil {
		return domain.Accounts{}, err
	}

	if res.ModifiedCount > 0 {
		return *newAccount, nil
	}

	return domain.Accounts{}, errors.New("conta bancária não encontrada")
}

func DeleteBankAccount(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	res, err := tools.DeleteOne(ctx, Collection, filter)

	if res.DeletedCount > 0 {
		return true, err
	}
	return false, err
}
