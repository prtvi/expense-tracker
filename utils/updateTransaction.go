package utils

import (
	"context"

	config "webdev/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// update a transaction, given the string id and toUpdateTransaction

func UpdateTransaction(id string, toUpdateTransaction bson.D) error {
	// convert string id to object.id
	docID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: docID}}
	update := bson.M{"$set": toUpdateTransaction}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Transactions.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		return cursor.Err()
	}

	return nil
}
