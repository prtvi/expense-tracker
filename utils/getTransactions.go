package utils

import (
	"context"
	"encoding/json"
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
		docByte, _ := json.Marshal(resultItem)
		var transaction model.Transaction
		json.Unmarshal(docByte, &transaction)

		allTransactions[i] = transaction
	}

	return allTransactions
}
