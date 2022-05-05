package utils

import (
	"context"

	"fmt"

	config "webdev/config"
	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// loops over all transactions and returns a model.CurrentStatus object with transaction summary

func UpdateCurrentStatus(allTransactions []model.Transaction) model.CurrentStatus {
	var currentStatus model.CurrentStatus

	for _, transaction := range allTransactions {
		if transaction.Type == "Income" {
			currentStatus.CurrentBalance += transaction.Amount
			currentStatus.TotalIncome += transaction.Amount
		} else {
			currentStatus.CurrentBalance -= transaction.Amount
			currentStatus.TotalExpense += transaction.Amount
		}
	}

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "total_income", Value: currentStatus.TotalIncome},
			{Key: "total_expense", Value: currentStatus.TotalExpense},
			{Key: "current_balance", Value: currentStatus.CurrentBalance},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.CurrentStatus.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		fmt.Println("err")
	}

	return currentStatus
}
