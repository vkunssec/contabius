package repository

import (
	"contabius/pkg/structs"
	"contabius/tools"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Collection = "bank_account"
)

func CreateBankAccount(account *structs.Accounts) (structs.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	result, err := tools.InsertOne(ctx, Collection, account)

	account.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return *account, err
}

func GetBankAccount(ids []string) ([]structs.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var accounts []structs.Accounts
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

func UpdateBankAccount(id string, newAccount *structs.Accounts) (structs.Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var account structs.Accounts

	filter := bson.M{"_id": tools.StringToObjectId(id)}

	result := tools.FindOne(ctx, Collection, filter, options.FindOne())
	result.Decode(&account)

	if newAccount.Account == "" {
		newAccount.Account = account.Account
	}
	if newAccount.Color == "" {
		newAccount.Color = account.Color
	}

	newAccount.CreatedAt = account.CreatedAt
	newAccount.UpdatedAt = time.Now()

	_, err := tools.UpdateOne(ctx, Collection, filter, newAccount)

	return *newAccount, err
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
