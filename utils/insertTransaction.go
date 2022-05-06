package utils

import (
	"context"

	config "webdev/config"

	"go.mongodb.org/mongo-driver/bson"
)

// insert transaction document to database

func InsertTransaction(transaction bson.D) error {
	_, err := config.Transactions.InsertOne(context.TODO(), transaction)
	if err != nil {
		return err
	}
	return nil
}
