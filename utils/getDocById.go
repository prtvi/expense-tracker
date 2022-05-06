package utils

import (
	"context"

	config "webdev/config"
	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// returns a model.Transaction obj by id

func GetDocumentById(id string) (model.Transaction, error) {
	// convert string id to object.id
	docID, _ := primitive.ObjectIDFromHex(id)

	cursor := config.Transactions.FindOne(context.TODO(), bson.D{{Key: "_id", Value: docID}})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		return model.Transaction{}, decodeErr
	}

	transaction := BsonDocToTransaction(fetchedDoc)
	return transaction, nil
}
