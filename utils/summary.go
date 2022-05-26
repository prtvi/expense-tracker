package utils

import (
	"context"

	"fmt"

	config "webdev/config"
	model "webdev/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// loops over all transactions and returns a model.Summary object with transaction summary

func UpdateSummary(allTransactions []model.Transaction) model.Summary {
	var summary model.Summary

	for _, transaction := range allTransactions {
		if transaction.Type == config.TypeIncomeID {
			summary.CurrentBalance += transaction.Amount
			summary.TotalIncome += transaction.Amount
		} else {
			summary.CurrentBalance -= transaction.Amount
			summary.TotalExpense += transaction.Amount
		}
	}

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "total_income", Value: summary.TotalIncome},
			{Key: "total_expense", Value: summary.TotalExpense},
			{Key: "current_balance", Value: summary.CurrentBalance},
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Summary.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		fmt.Println("err")
	}

	return summary
}
