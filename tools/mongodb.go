package tools

import (
	"context"

	"github.com/vkunssec/contabius/database"
	"github.com/vkunssec/contabius/utils/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StringToObjectId(id string) primitive.ObjectID {
	s, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		logger.Logger.Error().Err(e).Send()
	}
	return s
}

func ArrayStringToObjectId(arrString []string) []primitive.ObjectID {
	arrObjectId := make([]primitive.ObjectID, len(arrString))
	for i := range arrString {
		arrObjectId[i] = StringToObjectId(arrString[i])
	}
	return arrObjectId
}

func InsertOne(ctx context.Context, collection string, values interface{}) (*mongo.InsertOneResult, error) {
	return database.MongoDB.
		Collection(collection).
		InsertOne(ctx, values)
}

func UpdateOne(ctx context.Context, collection string, filter primitive.M, update primitive.M) (*mongo.UpdateResult, error) {
	return database.MongoDB.
		Collection(collection).
		UpdateOne(ctx, filter, update)
}

func Find(ctx context.Context, collection string, filters primitive.M, options *options.FindOptions) (*mongo.Cursor, error) {
	return database.MongoDB.
		Collection(collection).
		Find(ctx, filters, options)
}

func FindOne(ctx context.Context, collection string, filters primitive.M, options *options.FindOneOptions) *mongo.SingleResult {
	return database.MongoDB.
		Collection(collection).
		FindOne(ctx, filters, options)
}

func DeleteOne(ctx context.Context, collection string, filters primitive.M) (*mongo.DeleteResult, error) {
	return database.MongoDB.
		Collection(collection).
		DeleteOne(ctx, filters)
}
