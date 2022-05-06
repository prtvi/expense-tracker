package utils

import (
	"context"
	"fmt"

	config "webdev/config"
	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
)

// returns an array of model.Transaction from database

func GetTransactions() []model.Transaction {
	cursor, err := config.Transactions.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	allTransactions := make([]model.Transaction, len(results))

	for i, resultItem := range results {
		allTransactions[i] = BsonDocToTransaction(resultItem)
	}

	return allTransactions
}
