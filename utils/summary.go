package utils

import (
	"context"
	"encoding/json"
	"fmt"

	config "prtvi/expense-tracker/config"
	model "prtvi/expense-tracker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BsonDocToSummary(doc bson.M) model.Summary {
	var summary model.Summary
	docByte, _ := json.Marshal(doc)
	json.Unmarshal(docByte, &summary)

	return summary
}

// Create a model.Summary object for given transactions
func GetSummary(ts []model.Transaction) model.Summary {
	var summary model.Summary

	for _, transaction := range ts {
		if transaction.Type == config.TypeIncomeID {
			summary.TotalBalance += transaction.Amount
			summary.TotalIncome += transaction.Amount
		} else {
			summary.TotalBalance -= transaction.Amount
			summary.TotalExpense += transaction.Amount
		}
	}

	return summary
}

// loops over all transactions and returns a model.Summary object with transaction summary

func UpdateMainSummary(allTransactions []model.Transaction) model.Summary {
	summary := GetSummary(allTransactions)

	// create filter, update and options for querying
	filter := bson.M{}
	update := bson.M{
		"$set": bson.D{
			{Key: "total_income", Value: summary.TotalIncome},
			{Key: "total_expense", Value: summary.TotalExpense},
			{Key: "total_balance", Value: summary.TotalBalance},
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

func FetchMainSummary() model.Summary {
	cursor := config.Summary.FindOne(context.TODO(), bson.M{})

	fetchedDoc := bson.M{}
	decodeErr := cursor.Decode(&fetchedDoc)
	if decodeErr != nil {
		fmt.Println("error")
	}

	return BsonDocToSummary(fetchedDoc)
}
